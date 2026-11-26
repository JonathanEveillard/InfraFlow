package main

import (
    "log"
    "os"

    "ganigeorgiev.com/pocketbase/pocketbase"
    "ganigeorgiev.com/pocketbase/pocketbase/apis"
    "ganigeorgiev.com/pocketbase/pocketbase/core"
)

// Initializes Pocket Base
func initPocketBase(){
	app := pocketbase.New()

    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
        // serves static files from the provided public dir (if exists)
        se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

        return se.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}

func main() {
    initPocketBase()
}