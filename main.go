package main

import (
	"fmt"
	"github.com/camerontrewbase/Co-Pilot-Test/internal"
)

func main() {
	app, err := internal.NewApp()
	if err != nil {
		panic(err)
	}
	fmt.Print("Starting server on port 8080")
	app.Run()
}
