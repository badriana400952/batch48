package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	Nama       string
	PostDate   string
	EndDate    string
	Durasi     string
	Deskripsi  string
	NodeJs     bool
	ReactJS    bool
	NextJs     bool
	TypeScript bool
}

// slice
var dataBlog = []Blog{}

func main() {
	e := echo.New()
	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/addFormBlog", addFormBlog)
	e.GET("/blog", blog)
	e.GET("/testimonial", testimonial)
	e.GET("/kontak", kontak)
	e.GET("/detail/:id", detail)
	e.POST("/addblog", addblog)
	e.GET("/editBlog/:id", editBlog)

	e.POST("/delete-blog/:id", deleteBlog)

	e.Start(":1323")
}
func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	datas := map[string]interface{}{
		"Blogs": dataBlog,
	}
	return tmpl.Execute(c.Response(), datas)
}
func addFormBlog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/addFormBlog.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	datas := map[string]interface{}{
		"Blogs": dataBlog,
	}
	fmt.Println("ini data blog", datas)

	return tmpl.Execute(c.Response(), datas)
}

func testimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
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
	stc, _ := strconv.Atoi(id)
	details := Blog{}
	for index, data := range dataBlog {
		if index == stc {
			details = Blog{
				Nama:       data.Nama,
				PostDate:   data.PostDate,
				EndDate:    data.EndDate,
				Deskripsi:  data.Deskripsi,
				NodeJs:     data.NodeJs,
				ReactJS:    data.ReactJS,
				NextJs:     data.NextJs,
				TypeScript: data.TypeScript,
			}
		}
	}

	data := map[string]interface{}{
		"Id":     id,
		"Blogsh": details,
	}
	return tmpl.Execute(c.Response(), data)
}
func dursasiTanggal(start string, endDate string) string {
	awalMulai, _ := time.Parse("2006-01-02", start)

	akhirMulai, _ := time.Parse("2006-01-02", endDate)
	// untuk mengurangi dua waktu dan menghasilkan selisih waktu di antara keduanya.
	durasi := akhirMulai.Sub(awalMulai)

	years := durasi.Hours() / 24 / 365
	yearsInt := int(years)

	months := (durasi.Hours() / 24) / 30
	monthsInt := int(months)

	days := durasi.Hours() / 24
	daysInt := int(days)

	durasii := fmt.Sprintf("Durasi: %d tahun, %d bulan, %d hari", yearsInt, monthsInt, daysInt)

	return durasii

}

func addblog(c echo.Context) error {
	nama := c.FormValue("nama")
	deskripsi := c.FormValue("deskripsi")
	start := c.FormValue("start")
	endDate := c.FormValue("endDate")

	Durasi := dursasiTanggal(start, endDate)

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

	newBlog := Blog{
		Nama:       nama,
		PostDate:   start,
		EndDate:    endDate,
		Durasi:     Durasi,
		Deskripsi:  deskripsi,
		NodeJs:     nodeJss,
		ReactJS:    react,
		NextJs:     next,
		TypeScript: typeScript,
	}

	// lalu buat variabel untuk di timpa
	// appen membutuhkan 2 parameter 1 variabel baru 2 variabel lama untuk di timpa
	dataBlog = append(dataBlog, newBlog)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
func editBlog(c echo.Context) error {

	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/editFormBlog.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	stc, _ := strconv.Atoi(id)
	details := Blog{}
	for index, data := range dataBlog {
		if index == stc {
			details = Blog{
				Nama:       data.Nama,
				PostDate:   data.PostDate,
				EndDate:    data.EndDate,
				Deskripsi:  data.Deskripsi,
				NodeJs:     data.NodeJs,
				ReactJS:    data.ReactJS,
				NextJs:     data.NextJs,
				TypeScript: data.TypeScript,
			}
		}
	}

	data := map[string]interface{}{
		"Id":     id,
		"Blogsh": details,
	}
	return tmpl.Execute(c.Response(), data)

}

func deleteBlog(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	dataBlog = append(dataBlog[:idToInt], dataBlog[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
