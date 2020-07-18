package main

import (
	"flag"
	"os"
	"regexp"

	// Plugins
	_ "github.com/lucaslorentz/caddy-docker-proxy/plugin"

	// Plugins
	_ "github.com/BTBurke/caddy-jwt"
	_ "github.com/caddyserver/dnsproviders/cloudflare"
	_ "github.com/captncraig/caddy-realip"
	_ "github.com/jung-kurt/caddy-cgi"
	_ "github.com/nicolasazrak/caddy-cache"
	_ "github.com/tarent/loginsrv/caddy"

	// Caddy
	"github.com/caddyserver/caddy/caddy/caddymain"
)

var enableTelemetryFlag bool
var isTrue = regexp.MustCompile("(?i)^(true|yes|1)$")

func main() {
	flag.BoolVar(&enableTelemetryFlag, "enable-telemetry", false, "Enable caddy telemetry")

	flag.Parse()

	if enableTelemetryEnv := os.Getenv("CADDY_ENABLE_TELEMETRY"); enableTelemetryEnv != "" {
		caddymain.EnableTelemetry = isTrue.MatchString(enableTelemetryEnv)
	} else {
		caddymain.EnableTelemetry = enableTelemetryFlag
	}

	caddymain.Run()

	// Keep caddy running after main instance is stopped
	select {}
}
