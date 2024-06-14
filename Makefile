LDFLAGS = '-s -w -extldflags "-static" -X github.com/dzsak/url-shortener/pkg/version.Version='${VERSION}

build-web:
	(cd web; npm install; npm run build)
	rm -rf cmd/url-shortener/web/build
	mkdir -p cmd/url-shortener/web/build
	@cp -r web/dist/* cmd/url-shortener/web/build

build:
	CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o build/url-shortener github.com/dzsak/url-shortener/cmd/url-shortener

dist:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/linux/amd64/url-shortener github.com/dzsak/url-shortener/cmd/url-shortener
	GOOS=linux GOARCH=arm64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/linux/arm64/url-shortener github.com/dzsak/url-shortener/cmd/url-shortener

.PHONY: build-ui build dist
