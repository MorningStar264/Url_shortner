package main

import (
	"fmt"

	"github.com/MorningStar264/Url_shortner/api" 
)

func main() {
	fmt.Println("Starting date 25/12/2025")

	conf := settings.GetConfig()
	app := settings.Application{
		Config: conf,
	}
	app.Start()
}
