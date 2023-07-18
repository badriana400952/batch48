package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/views", "views")

	e.GET("/", home)
	e.GET("/blog", blog)
	e.GET("/testimonial", testimonial)
	e.GET("/kontak", kontak)
	e.GET("/detail/:id", detail)
	e.POST("/addblog", addBblog)
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
	tmpl, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
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
