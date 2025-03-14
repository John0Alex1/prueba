package main

import (
	server "v1/src/infraestructure/server"
	"v1/src/modules/products"
)

func main() {
	app := server.ProvidersStore{}
	app.Init()
	app.AddModule(products.ModuleProviders())
	app.Up()
}
