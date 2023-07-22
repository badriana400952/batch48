package main

import (
	"context"
	"fmt"
	connection "golang1/conection"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	Id         int
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
type Testimonial struct {
	ID      int
	Nama    string
	Content string
	Rating  int
	Images  string
}

// slice
var dataBlog = []Blog{}

func main() {
	connection.DatabaseConnect()
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
	e.POST("/formEditBlog", formEditditBlog)

	e.POST("/delete-blog/:id", deleteBlog)

	// daftarkan HandlerMiddleware ke global
	e.Use(HandlerMiddleware)
	e.GET("/testimonialApi", testimonialApi)
	e.Start(":3000")
}
func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html")

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, nama, stardate, enddate, durasi, deskripsi, react, nodejs, nextjs, typescript FROM coba5")

	var result []Blog
	for data.Next() {
		var each = Blog{}
		err := data.Scan(&each.Id, &each.Nama, &each.PostDate, &each.EndDate, &each.Durasi, &each.Deskripsi, &each.NodeJs, &each.ReactJS, &each.NextJs, &each.TypeScript)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		result = append(result, each)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	datas := map[string]interface{}{
		"Blogs": result,
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
func testimonialApi(c echo.Context) error {
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
	fmt.Println("ini  testimonial", hasil)

	return c.JSON(http.StatusOK, hasil)
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
func formEditditBlog(c echo.Context) error {

	id, _ := strconv.Atoi(c.FormValue("id"))
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
	fmt.Println("ga dapet", id)
	dataBlog[id].Nama = nama
	dataBlog[id].PostDate = start
	dataBlog[id].EndDate = endDate
	dataBlog[id].Durasi = Durasi
	dataBlog[id].Deskripsi = deskripsi
	dataBlog[id].NodeJs = nodeJss
	dataBlog[id].ReactJS = react
	dataBlog[id].NextJs = next
	dataBlog[id].TypeScript = typeScript

	return c.Redirect(http.StatusMovedPermanently, "/blog")

}
func deleteBlog(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)
	// skip index yg di pilih lalu gabungkan index yg lain lalu tampilkan
	dataBlog = append(dataBlog[:idToInt], dataBlog[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
