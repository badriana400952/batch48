package main

import (
	"context"
	"fmt"
	"golang1/views/connection"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Testimonial struct {
	ID      int
	Nama    string
	Content string
	Rating  int
	Images  string
}

func main() {
	connection.DatabaseConnect()
	e := echo.New()
	e.Static("/views", "views")

	e.GET("/", home)
	e.GET("/blog", blog)
	e.GET("/testimonial", testimonial)
	e.GET("/kontak", kontak)
	e.GET("/detail/:id", detail)
	e.POST("/addblog", addBblog)

	e.Use(HandlerMiddleware)
	e.Start(":1323")
}
func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func blog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/blog.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func testimonial(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, nama, content, rating, images FROM testimonoial")

	var hasil []Testimonial
	for data.Next() {
		var each = Testimonial{}

		err := data.Scan(&each.ID, &each.Nama, &each.Content, &each.Rating, &each.Images)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
		}
		hasil = append(hasil, each)

	}
	// testimonial := map[string]interface{}{
	// 	"Testimonial": hasil,
	// }

	// tmpl, err := template.ParseFiles("views/testimonial.html")

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{
	// 		"massage": err.Error()})
	// }

	return c.JSON(http.StatusOK, hasil)
}
func kontak(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/kontak.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func detail(c echo.Context) error {

	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	data := map[string]interface{}{
		"id":      id,
		"title":   "hallo",
		"content": "lorems",
	}
	return tmpl.Execute(c.Response(), data)
}

func addBblog(c echo.Context) error {
	nama := c.FormValue("nama")
	deskripsi := c.FormValue("deskripsi")
	start := c.FormValue("start")
	endDate := c.FormValue("endDate")

	var nodeJss bool
	if c.FormValue("nodeJss") == "yes" {
		nodeJss = true
	}
	var react bool
	if c.FormValue("react") == "yes" {
		react = true
	}
	var next bool
	if c.FormValue("next") == "yes" {
		next = true
	}
	var typeScript bool
	if c.FormValue("typeScript") == "yes" {
		typeScript = true
	}

	fmt.Println(nama, "nama")
	fmt.Println(deskripsi, "desk")
	fmt.Println(start, "start")
	fmt.Println(endDate, "endDate")

	fmt.Println("nodeJss", nodeJss)
	fmt.Println("react", react)
	fmt.Println("next", next)
	fmt.Println("typeScript", typeScript)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

// mid
func HandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		return next(c)
	}
}
