package main

import (
	"github.com/portofolio/models"
	"github.com/portofolio/routers"
)

func main() {
	// panggil koneksi ke database
	models.Connect()
	// panggil file route
	r := routers.Routers()
	// jalankan router
	r.Run()
}
