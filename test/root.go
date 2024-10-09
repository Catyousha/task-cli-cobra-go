package test

import (
	"fmt"
	"os"
	"tenessine/task-cli/db"
)

const testDb = "./tasks.json"
const taskCount = 100
const doneTaskCount = 10
const inProgressTaskCount = 20
const todoTaskCount = 70

func setup() {
	os.Create(testDb)
	db.Initialize(testDb)
	for i := 1; i <= taskCount; i++ {
		if i <= doneTaskCount {
			db.InsertTask(db.Task{
				Description: "Test Task " + fmt.Sprintf("%d", i),
				Status:      "done",
			})
		} else if i <= doneTaskCount + inProgressTaskCount {
			db.InsertTask(db.Task{
				Description: "Test Task " + fmt.Sprintf("%d", i),
				Status:      "in-progress",
			})
		} else {
			db.InsertTask(db.Task{
				Description: "Test Task " + fmt.Sprintf("%d", i),
				Status:      "todo",
			})
		}
	}
}

func teardown() {
	os.Remove(testDb)
}