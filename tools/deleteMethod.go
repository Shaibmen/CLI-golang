package tools

import (
	"database/sql"
	"fmt"
)

func Delete(db *sql.DB, idtask int) {
	deleteRequest, err := db.Exec(`DELETE FROM Task WHERE idtask = $1`, idtask)
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteRequest.RowsAffected())
}
