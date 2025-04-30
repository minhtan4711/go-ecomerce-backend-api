package main

import (
	"github.com/go-ecomerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run()
}
