package main

import (
	_ "github.com/aablinov/caddy-geoip"
	_ "github.com/abiosoft/caddy-git"
	_ "github.com/hacdias/caddy-webdav"

	"github.com/caddyserver/caddy/caddy/caddymain"
)

func main() {
	caddymain.EnableTelemetry = false

	caddymain.Run()
}
