package main

import (
	"tenessine/task-cli/cmd"
	"tenessine/task-cli/db"
)

func main() {
	db.Initialize("./db/tasks.json");
	cmd.Execute()
}
