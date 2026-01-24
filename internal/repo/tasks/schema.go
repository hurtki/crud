package tasks_repo

import (
	"database/sql"
	"fmt"
)

func ApplyTasksSchema(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		text TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_tasks_id ON tasks(id);
	`)

	if err != nil {
		return fmt.Errorf("cannot create tasks table: %s", err.Error())
	}
	return nil
}
