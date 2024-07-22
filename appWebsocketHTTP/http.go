package appWebsocketHTTP

import (
	"Systemge/Config"
	"Systemge/Http"
	"net/http"
)

func (app *AppWebsocketHTTP) GetHTTPComponentConfig() *Config.Http {
	return &Config.Http{
		Server: &Config.TcpServer{
			Port:        8080,
			TlsCertPath: "MyCertificate.crt",
			TlsKeyPath:  "MyKey.key",
		},
		Handlers: map[string]http.HandlerFunc{
			"/": Http.SendDirectory("../frontend"),
		},
	}
}
