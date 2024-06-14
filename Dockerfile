FROM --platform=$TARGETPLATFORM alpine:3

RUN addgroup -S url-shortener && adduser -S url-shortener -G url-shortener

RUN mkdir /url-shortener
RUN chown -R url-shortener:url-shortener /url-shortener
WORKDIR /url-shortener

ARG TARGETPLATFORM
ARG BUILDPLATFORM

COPY --chown=url-shortener:url-shortener bin/${TARGETPLATFORM}/url-shortener url-shortener
COPY --chown=url-shortener:url-shortener web/dist ./web/build/

USER url-shortener

EXPOSE 8080
CMD ["/url-shortener/url-shortener"]
