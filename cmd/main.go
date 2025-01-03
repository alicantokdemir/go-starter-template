package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	currentPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
		return
	}

	app := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir: fmt.Sprintf("%s/pb_data", currentPath),
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
