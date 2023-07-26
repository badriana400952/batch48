package main

import (
	"golang1/api"
	connection "golang1/conection"
	route "golang1/models"

	"github.com/labstack/echo/v4"
)

func main() {
	connection.DatabaseConnecttion()
	e := echo.New()
	e.Static("/assets", "assets")

	e.GET("/", route.Home)
	e.GET("/addFormBlog", route.AddFormBlog)
	e.GET("/blog", route.Bloger)
	e.GET("/testimonial", route.Testimonials)
	e.GET("/kontak", route.Kontak)
	e.GET("/detail/:id", route.Detail)
	e.POST("/addblog", route.Addblog)
	e.GET("/editBlog/:id", route.EditBlog)
	e.POST("/formEditBlog/:id", route.FormEditditBlog)

	e.POST("/delete-blog/:id", route.DeleteBlog)

	// daftarkan HandlerMiddleware ke global
	e.Use(HandlerMiddleware)
	e.GET("/testimonialApi", api.TestimonialApi)
	e.Start(":3000")
}

// membuat perantara dalam proses pengolahan permintaan request dan pemrosesan respons dalam  API.
func HandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// c.Response().Header().Set() adalah sebuah cara untuk mengatur nilai pada header HTTP dari respons (response) yang akan dikirimkan oleh server.

		// berarti memperbolehkan akses dari sumber yang berbeda (cross-origin).
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		// menetapkan metode HTTP yang diizinkan oleh sumber daya yang teridentifikasi oleh URL. Dalam hal ini, izin diberikan untuk metode GET, POST, PUT, DELETE, dan OPTIONS.
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// daftar header yang diizinkan oleh sumber daya yang teridentifikasi oleh URL. Dalam hal ini, izin diberikan untuk header "Origin", "Content-Type", dan "Authorization".
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		//kembalikan fungsi c
		return next(c)
	}
}
