package route

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
// var dataBlogs = []Blog{}

func Home(c echo.Context) error {
	connection.DatabaseConnecttion()
	tmpl, err := template.ParseFiles("views/index.html")

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, nama, stardate, enddate, durasi, deskripsi, react, nodejs, nextjs, typescript FROM coba5")

	var result []Blog
	// baca data secara berurutan dari database
	for data.Next() {
		var each = Blog{}
		// mengambil nilai dari database
		err := data.Scan(&each.Id, &each.Nama, &each.PostDate, &each.EndDate, &each.Durasi, &each.Deskripsi, &each.NodeJs, &each.ReactJS, &each.NextJs, &each.TypeScript)
		// lalu tampilkan kesalalahan jika ada
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

func AddFormBlog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/addFormBlog.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func Bloger(c echo.Context) error {

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, nama, stardate, enddate, durasi, deskripsi, react, nodejs, nextjs, typescript FROM coba5")

	var result []Blog
	for data.Next() {
		var each = Blog{}
		data.Scan(&each.Id, &each.Nama, &each.PostDate, &each.EndDate, &each.Durasi, &each.Deskripsi, &each.NodeJs, &each.ReactJS, &each.NextJs, &each.TypeScript)

		result = append(result, each)
	}

	tmpl, err := template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	datas := map[string]interface{}{
		"Blogs": result,
	}

	return tmpl.Execute(c.Response(), datas)
}

func Testimonials(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func Kontak(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/kontak.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func Detail(c echo.Context) error {

	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	stc, _ := strconv.Atoi(id)

	var details = Blog{}

	errs := connection.Conn.QueryRow(context.Background(), "SELECT id, nama, stardate, enddate, durasi, deskripsi, react, nodejs, nextjs, typescript FROM coba5 WHERE id=$1", stc).Scan(&details.Id, &details.Nama, &details.PostDate, &details.EndDate, &details.Durasi, &details.Deskripsi, &details.NodeJs, &details.ReactJS, &details.NextJs, &details.TypeScript)

	if errs != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}

	data := map[string]interface{}{
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
	fmt.Println(daysInt)
	durasii := fmt.Sprintf("Durasi: %d tahun, %d bulan", yearsInt, monthsInt)

	return durasii
}

func Addblog(c echo.Context) error {
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

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO coba5( nama, stardate, enddate, durasi, deskripsi, react, nodejs, nextjs, typescript) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", nama, start, endDate, Durasi, deskripsi, nodeJss, react, next, typeScript)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func EditBlog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/editFormBlog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	// ubah string menjadi tipe data integer
	id, _ := strconv.Atoi(c.Param("id"))
	var edit = Blog{}

	errs := connection.Conn.QueryRow(context.Background(), "SELECT id, nama, stardate, enddate, durasi, deskripsi, react, nodejs, nextjs, typescript FROM coba5 WHERE id=$1", id).Scan(&edit.Id, &edit.Nama, &edit.PostDate, &edit.EndDate, &edit.Durasi, &edit.Deskripsi, &edit.NodeJs, &edit.ReactJS, &edit.NextJs, &edit.TypeScript)

	if errs != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}

	data := map[string]interface{}{
		"Id":     id,
		"Blogsh": edit,
	}

	return tmpl.Execute(c.Response(), data)
}

func FormEditditBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	nama := c.FormValue("nama")
	start := c.FormValue("start")
	endDate := c.FormValue("endDate")
	Durasi := dursasiTanggal(start, endDate)
	deskripsi := c.FormValue("deskripsi")

	react := c.FormValue("react")
	nodeJss := c.FormValue("nodeJss")
	next := c.FormValue("next")
	typeScript := c.FormValue("typeScript")
	fmt.Println(id, "ini iddddd")

	fmt.Println(nama)
	fmt.Println(start)
	fmt.Println(endDate)
	fmt.Println(Durasi)
	fmt.Println(deskripsi)
	fmt.Println(react)
	fmt.Println(nodeJss)
	fmt.Println(next)
	fmt.Println(typeScript)

	// img := c.FormValue("a.jpg")

	_, errs := connection.Conn.Exec(context.Background(), "UPDATE coba5 SET nama=$1, stardate=$2, enddate=$3, durasi=$4, deskripsi=$5, react=$6, nodejs=$7, nextjs=$8, typescript=$9 WHERE id=$10", nama, start, endDate, Durasi, deskripsi, react != "", nodeJss != "", next != "", typeScript != "", id)

	if errs != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"erooooorrrrrrrrr": errs.Error()})
	}
	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func DeleteBlog(c echo.Context) error {
	i := c.Param("id")

	id, _ := strconv.Atoi(i)
	connection.Conn.Exec(context.Background(), "DELETE FROM coba5 WHERE id=$1 ", id)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
