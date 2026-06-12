package main

import (
	"log"

	"github.com/jeanscalabrin/bolao-copa/internal/router"
)

func main() {
	r := router.Setup()

	log.Fatal(r.Run(":3000"))
}
