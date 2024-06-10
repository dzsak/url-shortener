package store

import (
	"testing"

	"github.com/dzsak/url-shortener/pkg/model"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestPodCRUD(t *testing.T) {
	s, err := NewTest()
	assert.Nil(t, err)
	defer s.Close()

	url := model.Url{
		Original: "https://github.com",
		ShortKey: "abc",
	}

	err = s.InsertUrl(url)
	assert.Nil(t, err)

	urlFromStore, err := s.GetUrlByOriginal(url.Original)
	assert.Nil(t, err)

	assert.Equal(t, url.Original, urlFromStore.Original)
	assert.Equal(t, url.ShortKey, urlFromStore.ShortKey)
}
