package tools

import (
	"database/sql"
	"fmt"
)

func Put(db *sql.DB, idPut int, statusBol bool) {
	deleteRequest, err := db.Exec(`UPDATE Task SET Status =  $1 WHERE idtask = $2`, statusBol, idPut)
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteRequest.RowsAffected())
}
