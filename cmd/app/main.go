package main

import "github.com/stickpro/go-shop/internal/app"

const configDir = "configs"

func main() {
	app.Run(configDir)
}
