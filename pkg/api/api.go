package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/dzsak/url-shortener/pkg/model"
	"github.com/dzsak/url-shortener/pkg/store"
	"github.com/go-chi/chi/v5"
)

type formData struct {
	URL string `json:"url"`
}

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	var formData formData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, "cannot decode body: %s", http.StatusBadRequest)
		return
	}

	originalURL := formData.URL
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	store := ctx.Value("store").(*store.Store)
	urlFromDB, err := store.GetUrlByOriginal(originalURL)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "cannot get url db: %s", 500)
		return
	}

	if urlFromDB.Original != "" {
		resp := map[string]string{
			"shortenedUrl": fmt.Sprintf("%s/short/%s", r.Host, urlFromDB.ShortKey),
		}

		respString, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "cannot serialize url: %s", 500)
			return
		}

		w.WriteHeader(200)
		w.Write(respString)
		return
	}

	shortKey := generateShortKey()
	url := model.Url{
		Original: originalURL,
		ShortKey: shortKey,
	}

	err = store.InsertUrl(url)
	if err != nil {
		http.Error(w, "cannot insert url to db: %s", 500)
		return
	}

	resp := map[string]string{
		"shortenedUrl": fmt.Sprintf("%s/short/%s", r.Host, shortKey),
	}

	respString, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "cannot serialize url: %s", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(respString)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := chi.URLParam(r, "key")
	if shortKey == "" {
		http.Error(w, "shortened key is missing", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	store := ctx.Value("store").(*store.Store)
	urlFromDB, err := store.GetUrlByShortKey(shortKey)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "cannot get url from db", http.StatusNotFound)
		return
	}

	if urlFromDB.Original == "" {
		http.Error(w, "original not found", http.StatusNotFound)
		return
	}

	// Redirect the user to the original URL
	http.Redirect(w, r, urlFromDB.Original, http.StatusMovedPermanently)
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
