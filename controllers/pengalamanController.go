package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/portofolio/models"
)

// gin.context adalah cara untuk memanggil fungsi dari gin gonic
func Getpengalaman(c *gin.Context) {
	// memanggi file mode pengalaman
	var pengalaman []models.Pengalaman
	// perintah di bawah berfungsi untuk mengambil semua data pengalaman
	models.DB.Find(&pengalaman)
	// jika semua data sudah di masukan kedalam variable pengalaman maka kirim response json
	// gin.H untuk menambahkan header khusus ke respons HTTP
	c.JSON(http.StatusOK, gin.H{"data": pengalaman})
}

// gin.context adalah cara untuk memanggil fungsi dari gin gonic
func SavePengalaman(c *gin.Context) {
	// memanggi file mode pengalaman
	var pengalaman models.Pengalaman

	// cek apakah data yg dikirim ada atau kosong
	if err := c.ShouldBindJSON(&pengalaman); err != nil {
		// jika data yang dikirm kosong maka kirim response
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// simpan data kedalam database
	models.DB.Create(&pengalaman)
	// jika sudah kirim response jika data berhasil disimpan
	// gin.H untuk menambahkan header khusus ke respons HTTP
	c.JSON(http.StatusOK, gin.H{"data": pengalaman, "message": "Data Berhasil Disimpan"})
}

// gin.context adalah cara untuk memanggil fungsi dari gin gonic
func EditPengalaman(c *gin.Context) {
	// memanggi file mode pengalaman
	var pengalaman models.Pengalaman
	// mengambil nilai dari parameter id
	id := c.Param("id")

	// maksud if dibawah adalah
	// jika data dari model pengalaman dengan id yang kita cari ada maka return data array jika tidak maka data null
	// jika di php maka kondisi seperti if($err != null){}
	if err := models.DB.First(&pengalaman, id).Error; err != nil {
		switch err {
		// jika pendari dari database tidak ditemukan maka
		case gorm.ErrRecordNotFound:
			// fungsi dibawah membatalkan dengan status 404 dan mengirim response dengan header HTTP
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Tidak Ditemukan"})
			// return dibawah untuk menghentikan proses
			return
		default:
			// fungsi dibawah membatalkan dengan status 404 dan mengirim response dengan header HTTP
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			// return dibawah untuk menghentikan proses
			return
		}
	}
	// jika semua data sudah di masukan kedalam variable pengalaman maka kirim response json
	// gin.H untuk menambahkan header khusus ke respons HTTP
	c.JSON(http.StatusOK, gin.H{"data": pengalaman})
}

// gin.context adalah cara untuk memanggil fungsi dari gin gonic
func UpdatePengalaman(c *gin.Context) {
	// memanggi file mode pengalaman
	var pengalaman models.Pengalaman
	// mengambil nilai dari parameter id
	id := c.Param("id")

	// cek apakah perubahan data ada atau tidak
	if err := c.ShouldBindJSON(&pengalaman); err != nil {
		// fungsi dibawah membatalkan dengan status 404 dan mengirim response dengan header HTTP
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Tidak Ditemukan"})
		// return dibawah untuk menghentikan proses
		return
	}

	// cek apakah perubahan data pada database
	if models.DB.Model(&pengalaman).Where("id = ?", id).Update(&pengalaman).RowsAffected == 0 {
		// fungsi dibawah membatalkan dengan status 404 dan mengirim response dengan header HTTP
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Di Update"})
		// return dibawah untuk menghentikan proses
		return
	}

	// gin.H untuk menambahkan header khusus ke respons HTTP
	c.JSON(http.StatusOK, gin.H{"data": pengalaman, "message": "Data Berhasil Di Update"})
}

// gin.context adalah cara untuk memanggil fungsi dari gin gonic
func DeletePengalaman(c *gin.Context) {
	// memanggi file mode pengalaman
	var pengalaman models.Pengalaman
	// mengambil nilai dari parameter id
	id := c.Param("id")

	// cek apakah perubahan data pada database
	if models.DB.Delete(&pengalaman, id).RowsAffected == 0 {
		// fungsi dibawah membatalkan dengan status 404 dan mengirim response dengan header HTTP
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data Gagal Di Hapus"})
		// return dibawah untuk menghentikan proses
		return
	}

	// gin.H untuk menambahkan header khusus ke respons HTTP
	c.JSON(http.StatusOK, gin.H{"data": pengalaman, "message": "Data Berhasil Di Hapus"})
}

func UploadFile(c *gin.Context) string {
	// cek file foto dari inputan
	file, header, err := c.Request.FormFile("foto")

	// cek jika file ada atau tidak
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return ""
	}

	// deklarasi nama file
	filename := header.Filename

	// lokasi tempat file disimpan
	out, err := os.Create("public/" + filename)

	// cek apaka berhasil atau gagal
	if err != nil {
		log.Fatal(err)
	}

	// defer adalah sebuah kata kunci yang digunakan untuk menjadwalkan eksekusi suatu fungsi hingga fungsi yang berisi
	defer out.Close()
	// pindahkan file foto kedalam file tmp
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return filename

}
