package main

import (
	"os"

	"github.com/TrAyZeN/gimme-repos/routes"
)

func main() {
	p := os.Getenv("PORT")

	if p != "" {
		routes.Listen(p)
	} else {
		routes.Listen("3000")
	}
}
