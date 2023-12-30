package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/portofolio/controllers"
)

// pada function router jalankan gin engin agar pengalamatan bisa di fungsikan
func Routers() *gin.Engine {
	// config gin gonic default
	r := gin.Default()

	// contoh pengalamatan menggunakan method GET
	r.GET("/api/pengalaman", controllers.Getpengalaman)
	// contoh pengalamatan menggunakan method POST
	r.POST("/api/save-pengalaman", controllers.SavePengalaman)
	// contoh pengalamatan menggunakan method GET dengan PARAMETER
	r.GET("/api/edit-pengalaman/:id", controllers.EditPengalaman)
	// contoh pengalamatan menggunakan method PUT dengan PARAMETER
	r.PUT("/api/update-pengalaman/:id", controllers.UpdatePengalaman)
	// contoh pengalamatan menggunakan method DELETE
	r.DELETE("/api/delete-pengalaman/:id", controllers.DeletePengalaman)

	return r
}
