package main

import "log"

func main() {
	err := startApp()
	if err != nil {
		log.Fatal("error on starting the app", err)
	}
}
