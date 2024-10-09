package db

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func GetAllTasks() ([]Task, error) {
	f, err := os.Open(dbLoc)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer f.Close()

	d, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	if len(d) == 0 {
		return []Task{}, nil
	}

	var ts []Task
	err = json.Unmarshal(d, &ts)
	
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	return ts, nil
}

func GetTasksByStatus(status string) ([]Task, error) {
	ts, err := GetAllTasks()
	if err != nil {
		return nil, err
	}

	var tasks []Task

	for _, t := range ts {
		if t.Status == status {
			tasks = append(tasks, t)
		}
	}

	return tasks, nil
}

func InsertTask(t Task) (Task, error) {
	ts, err := GetAllTasks()
	if err != nil {
		return Task{}, err
	}

	if len(ts) == 0 {
		n := t
		n.ID = 1
		ts = append(ts, n)
	} else {
		n := t
		n.ID = ts[len(ts)-1].ID + 1
		ts = append(ts, n)
	}

	d, err := json.MarshalIndent(ts, "", "  ")
	if err != nil {
		return Task{}, fmt.Errorf("error marshalling JSON: %w", err)
	}

	err = os.WriteFile(dbLoc, d, 0644)
	if err != nil {
		return Task{}, fmt.Errorf("error writing file: %w", err)
	}
	return t, nil
}

func UpdateTaskDescription(id int, description string) (Task, error) {
	ts, err := GetAllTasks()
	if err != nil {
		return Task{}, err
	}

	t := Task{}

	for i, t2 := range ts {
		if t2.ID == id {
			t = ts[i]
			t.Description = description
			ts[i] = t
			break
		}
	}

	d, err := json.MarshalIndent(ts, "", "  ")
	if err != nil {
		return Task{}, fmt.Errorf("error marshalling JSON: %w", err)
	}

	err = os.WriteFile(dbLoc, d, 0644)
	if err != nil {
		return Task{}, fmt.Errorf("error writing file: %w", err)
	}

	return t, nil
}

func UpdateTaskStatus(id int, status string) (Task, error) {
	ts, err := GetAllTasks()
	if err != nil {
		return Task{}, err
	}

	t := Task{}

	for i, t2 := range ts {
		if t2.ID == id {
			t = ts[i]
			t.Status = status
			ts[i] = t
			break
		}
	}

	d, err := json.MarshalIndent(ts, "", "  ")
	if err != nil {
		return Task{}, fmt.Errorf("error marshalling JSON: %w", err)
	}

	err = os.WriteFile(dbLoc, d, 0644)
	if err != nil {
		return Task{}, fmt.Errorf("error writing file: %w", err)
	}

	return t, nil
}

func DeleteTask(id int) (Task, error) {
	ts, err := GetAllTasks()
	if err != nil {
		return Task{}, err
	}

	f := false
	for i, t2 := range ts {
		if t2.ID == id {
			ts = append(ts[:i], ts[i+1:]...)
			f = true
			break
		}
	}

	d, err := json.MarshalIndent(ts, "", "  ")
	if err != nil {
		return Task{}, fmt.Errorf("error marshalling JSON: %w", err)
	}

	err = os.WriteFile(dbLoc, d, 0644)
	if err != nil {
		return Task{}, fmt.Errorf("error writing file: %w", err)
	}

	if !f {
		return Task{}, fmt.Errorf("task not found")
	}

	return Task{}, nil
}
