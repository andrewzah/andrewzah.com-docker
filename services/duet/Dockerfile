FROM php:7.2.5-fpm-alpine3.6
MAINTAINER Andrew Zah <zah@andrewzah.com>
LABEL caddy_version="0.10.12" architecture="amd64"

ARG PLUGINS=git
RUN apk add --no-cache openssh-client tar curl

RUN curl --silent --show-error --fail --location \
      --header "Accept: application/tar+gzip, application/x-gzip, application/octet-stream" -o - \
      "https://caddyserver.com/download/linux/amd64?plugins=${plugins}&license=personal&telemetry=off" \
    | tar --no-same-owner -C /usr/bin/ -xz caddy \
 && chmod 0755 /usr/bin/caddy \
 && /usr/bin/caddy -version

# why
RUN docker-php-ext-install mysqli pdo pdo_mysql

COPY ./src /srv/duet

ENTRYPOINT ["/usr/bin/caddy"]
CMD ["--conf", "/etc/Caddyfile", "--log", "stdout"]
