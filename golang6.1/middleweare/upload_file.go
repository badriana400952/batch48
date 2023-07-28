package middleweare

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc { //membutuhkan next echo.HandlerFunc) echo.HandlerFunc
	// jalankan return fungsi echo context
	return func(c echo.Context) error {
		file, err := c.FormFile("fileImage") // tangkap file dari input
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// fmt.Println("ini file", file)
		// buka file
		src, errs := file.Open()

		if errs != nil {
			return c.JSON(http.StatusBadRequest, errs.Error())
		}
		defer src.Close() // lifo fungsi nya untuk jalankan di akhir tapi bebas di taro mau di atas => manfaat agar tidak kebocoran memory

		tmpFIle, errss := ioutil.TempFile("uploads", "images-*.png")
		if errss != nil {
			return c.JSON(http.StatusBadRequest, errss.Error())
		}
		defer tmpFIle.Close()
		fmt.Println("tmpFile", tmpFIle)

		writenCopy, err := io.Copy(tmpFIle, src)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errs.Error())
		}
		fmt.Println("writenCopy", writenCopy)

		// ambil nama file
		data := tmpFIle.Name()
		fileName := data[8:] //ambil nama file dari index ke delapan ke depan
		fmt.Println("fileName", fileName)

		// set ke databse
		c.Set("dataFile", fileName)

		// jalankan handller selanjut nya yaitu function (route.Addblog)
		return next(c)
		// return c.String(http.StatusOK, "berhasil")

	}
}
