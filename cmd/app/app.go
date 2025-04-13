package main

import "github.com/starpia-forge/bunjang-watch/internal/app"

func main() {
	core, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	core.Run()
}
