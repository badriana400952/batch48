package route

import (
	"context"
	"database/sql"
	"fmt"
	connection "golang1/conection"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Blog struct {
	Id         int
	Authdor    string
	Nama       string
	PostDate   string
	EndDate    string
	Durasi     string
	Deskripsi  string
	NodeJs     bool
	ReactJS    bool
	NextJs     bool
	TypeScript bool
	Img        string
}
type Testimonial struct {
	ID      int
	Nama    string
	Content string
	Rating  int
	Images  string
}
type Users struct {
	ID       int
	Name     string
	Email    string
	Password string
}
type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

// slice
// var dataBlogs = []Blog{}

// relasi
// SELECT tb_user.id, tb_user.name, tb_user.email, coba5.id, coba5.nama, coba5.stardate, coba5.enddate, coba5.durasi, coba5.deskripsi, coba5.react, coba5.nodejs, coba5.nextjs, coba5.typescript, coba5.img FROM coba5 LEFT JOIN tb_user ON coba5.authdor_id = tb_user.id

func Home(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	datas := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
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

	data, _ := connection.Conn.Query(context.Background(), "SELECT coba5.id, tb_user.name, coba5.nama, coba5.stardate, coba5.enddate, coba5.durasi, coba5.deskripsi, coba5.nodejs, coba5.react, coba5.nextjs, coba5.typescript, coba5.img FROM coba5 LEFT JOIN tb_user ON coba5.authdor_id = tb_user.id")

	var result []Blog
	for data.Next() {
		var each = Blog{}
		var authdorSementara sql.NullString

		errs := data.Scan(&each.Id, &authdorSementara, &each.Nama, &each.PostDate, &each.EndDate, &each.Durasi, &each.Deskripsi, &each.NodeJs, &each.ReactJS, &each.NextJs, &each.TypeScript, &each.Img)

		if errs != nil {
			return c.JSON(http.StatusInternalServerError, errs.Error())
		}

		each.Authdor = authdorSementara.String
		result = append(result, each)
	}

	tmpl, err := template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	datas := map[string]interface{}{
		"Blogs":       result,
		"DataSession": userData,
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
	var thempAuthor sql.NullString
	errs := connection.Conn.QueryRow(context.Background(), "SELECT coba5.id, tb_user.name, coba5.nama, coba5.stardate, coba5.enddate, coba5.durasi, coba5.deskripsi, coba5.nodejs, coba5.react, coba5.nextjs, coba5.typescript, img FROM coba5 LEFT JOIN tb_user ON coba5.authdor_id = tb_user.id WHERE coba5.id=$1", stc).Scan(&details.Id, &thempAuthor, &details.Nama, &details.PostDate, &details.EndDate, &details.Durasi, &details.Deskripsi, &details.NodeJs, &details.ReactJS, &details.NextJs, &details.TypeScript, &details.Img)

	details.Authdor = thempAuthor.String
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
	img := c.Get("dataFile").(string) //string
	sess, _ := session.Get("session", c)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO coba5( nama, stardate, enddate, durasi, deskripsi, react, nodejs, nextjs, typescript, img, authdor_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10,$11)", nama, start, endDate, Durasi, deskripsi, nodeJss, react, next, typeScript, img, sess.Values["id"].(int))

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

	errs := connection.Conn.QueryRow(context.Background(), "SELECT id, nama, stardate, enddate, durasi, deskripsi, react, nodejs, nextjs, typescript, img FROM coba5 WHERE id=$1", id).Scan(&edit.Id, &edit.Nama, &edit.PostDate, &edit.EndDate, &edit.Durasi, &edit.Deskripsi, &edit.NodeJs, &edit.ReactJS, &edit.NextJs, &edit.TypeScript, &edit.Img)

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
	img := c.Get("dataFile").(string) //string

	fmt.Println(id, "ini iddddd")

	// fmt.Println(nama)
	// fmt.Println(start)
	// fmt.Println(endDate)
	// fmt.Println(Durasi)
	// fmt.Println(deskripsi)
	// fmt.Println(react)
	// fmt.Println(nodeJss)
	// fmt.Println(next)
	// fmt.Println(typeScript)

	_, errs := connection.Conn.Exec(context.Background(), "UPDATE coba5 SET nama=$1, stardate=$2, enddate=$3, durasi=$4, deskripsi=$5, react=$6, nodejs=$7, nextjs=$8, typescript=$9, img=$10 WHERE id=$11", nama, start, endDate, Durasi, deskripsi, react != "", nodeJss != "", next != "", typeScript != "", img, id)

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

func Regis(c echo.Context) error {
	var template, err = template.ParseFiles("views/regis.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return template.Execute(c.Response(), nil)
}
func Logins(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())
	var template, err = template.ParseFiles("views/login.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	return template.Execute(c.Response(), flash)
}
func RegisterMethod(c echo.Context) error {
	errs := c.Request().ParseForm()
	if errs != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": "ini error c.Request"})
	}

	name := c.FormValue("inputName")
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/register")
	}

	return redirectWithMessage(c, "Register success!", true, "/login")
}
func LoginMethod(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	user := Users{}
	ers := connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, password FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if ers != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}
	// komperesi password yang dari database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect!", false, "/login")
	}
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}

// SELECT tb_user.id, tb_user.name, tb_user.email, coba5.id, coba5.nama, coba5.stardate, coba5.enddate, coba5.durasi, coba5.deskripsi, coba5.react, coba5.nodejs, coba5.nextjs, coba5.typescript, coba5.img LEFT JOIN tb_user ON coba5.authdor_id = tb_user.id
