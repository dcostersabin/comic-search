package table

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func createQuery() string {
	return "CREATE TABLE comics(id INT PRIMARY KEY , month VARCHAR(255) , link VARCHAR(255) , year VARCHAR(255) , news VARCHAR(255) , safe_title VARCHAR(255) , transcript TEXT , alt VARCHAR(255) , img TEXT , title VARCHAR(255) , day VARCHAR(255) );"
}

func CreateTable(conn *pgx.Conn) {
	_, _ = conn.Exec(context.Background(), createQuery())

}
