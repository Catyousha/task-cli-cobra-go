package test

import (
	"tenessine/task-cli/db"
	"testing"
)

func TestGetAllTasks(t *testing.T) {
	setup()
	defer teardown()
	ts, err := db.GetAllTasks()
	if err != nil {
		t.Errorf("Failed to get all tasks: %v", err)
	}
	if len(ts) != taskCount {
		t.Errorf("Expected %d tasks, but got %d", taskCount, len(ts))
	}
}

func TestGetTasksByStatus(t *testing.T) {
	setup()
	defer teardown()
	ts, err := db.GetTasksByStatus("done")
	if err != nil {
		t.Errorf("Failed to get tasks by status: %v", err)
	}
	if len(ts) != doneTaskCount {
		t.Errorf("Expected %d done tasks, but got %d", doneTaskCount, len(ts))
	}

	ts, err = db.GetTasksByStatus("in-progress")
	if err != nil {
		t.Errorf("Failed to get tasks by status: %v", err)
	}
	if len(ts) != inProgressTaskCount {
		t.Errorf("Expected %d in-progress tasks, but got %d", inProgressTaskCount, len(ts))
	}

	ts, err = db.GetTasksByStatus("todo")
	if err != nil {
		t.Errorf("Failed to get tasks by status: %v", err)
	}
	if len(ts) != todoTaskCount {
		t.Errorf("Expected %d todo tasks, but got %d", todoTaskCount, len(ts))
	}
}

func TestModifyTaskStatus(t *testing.T) {
	setup()
	defer teardown()
	db.UpdateTaskStatus(doneTaskCount+1, "done")
	ts, err := db.GetTasksByStatus("done")
	if err != nil {
		t.Errorf("Failed to get tasks by status: %v", err)
	}
	if len(ts) != doneTaskCount + 1 {
		t.Errorf("Expected %d done tasks, but got %d", doneTaskCount + 1, len(ts))
	}
}
