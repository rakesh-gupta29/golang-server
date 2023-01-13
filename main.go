package main

import "github.com/rakesh-gupta29/golang-server/app"

func main() {
	err := app.MountAndRun()
	if err != nil {
		panic(err)
	}
}
