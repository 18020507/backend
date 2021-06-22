package connectDB

import (
	"database/sql"
	"fmt"
	"hieu/pkgs/dotenv"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	dotenv.GetMySQL()

	fmt.Print(dotenv.GetMySQL())

	db, err := sql.Open("mysql", dotenv.GetMySQL())

	if err != nil {
		panic(err.Error())
	}

	// defer db.Close()
	return db
}
