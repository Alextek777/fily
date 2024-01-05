package app

import (
	. "github.com/Alextek777/fily/src/internal/config"
	. "github.com/Alextek777/fily/src/internal/storage"
)

type Fily struct {
	webApp *webServer
	store  Storage
	// auth service
}

func MustNewFily(cfg *Config) *Fily {
	store, err := NewPostgresStore(cfg)
	if err != nil {
		panic(err)
	}

	err = store.InitStorage()
	if err != nil {
		panic(err)
	}

	webApp := newWebServer(cfg, store)

	return &Fily{webApp: webApp, store: store}
}

func (app *Fily) RunApp() {

	app.webApp.run()
}
