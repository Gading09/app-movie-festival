package main

import (
	"fmt"
	"log"
	"movie-festival/delivery/container"
	"movie-festival/delivery/http"
)

func main() {
	container := container.SetupContainer()
	handler := http.SetupHandler(container)
	http := http.ServeHttp(handler, container.Database)
	err := http.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.Port))

	if err != nil {
		log.Fatal("Service http server error: ", err.Error())
	}
}
