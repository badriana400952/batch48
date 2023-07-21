package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	databaseUrl := "postgres://postgres:123456789@localhost:5432/dataBlog"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintln(os.Stderr, "database connection %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Succesfully connected to database.")
}
