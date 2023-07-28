package api

import (
	"context"
	"fmt"
	connection "golang1/conection"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Testimonial struct {
	ID      int
	Nama    string
	Content string
	Rating  int
	Images  string
}

func TestimonialApi(c echo.Context) error {
	connection.DatabaseConnecttion()
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
	// fmt.Println("ini  testimonial", hasil)

	return c.JSON(http.StatusOK, hasil)

}
