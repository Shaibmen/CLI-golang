package tools

import (
	"database/sql"
	"fmt"

	"github.com/pterm/pterm"
)

type dbTasks struct {
	idTask      int
	title       string
	description string
	status      bool
}

func Get(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM Task ORDER BY idtask")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tasks := []dbTasks{}

	for rows.Next() {
		t := dbTasks{}

		err := rows.Scan(&t.idTask, &t.title, &t.description, &t.status)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, t)
	}

	toPrint := pterm.TableData{}
	for _, task := range tasks {
		toPrint = append(toPrint, []string{
			fmt.Sprintf("%d", task.idTask),
			task.title,
			task.description,
			fmt.Sprintf("%v", task.status),
		})
	}
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"ID", "Title", "Description", "Status"},
	}).WithData(toPrint).Render()
}
