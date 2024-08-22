package main

import (
	"BookHub/internal/application"
)

func main() {
	err := application.Run()
	if err != nil {
		return
	}
}
