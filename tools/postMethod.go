package tools

import (
	"database/sql"
	"fmt"
)

func Post(db *sql.DB, title string, description string, status bool) {
	postRequest, err := db.Exec(`
	INSERT INTO Task(Title, Description, Status)
	VALUES
	($1, $2, $3)`, title, description, status)
	if err != nil {
		panic(err)
	}
	fmt.Println(postRequest.RowsAffected())
}
