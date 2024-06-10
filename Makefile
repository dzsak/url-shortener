build-ui:
	(cd web; npm install; npm run build)
	rm -rf cmd/url-shortener/web/build
	mkdir -p cmd/url-shortener/web/build
	@cp -r web/dist/* cmd/url-shortener/web/build
