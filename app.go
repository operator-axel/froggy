package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Completed   bool   `json:"completed"`
}

var (
	tasks  []Task
	mutex  sync.Mutex
	nextID = 1
)

func (a *App) AddTask(title, description, tag string, completed bool) Task {
	mutex.Lock()
	defer mutex.Unlock()

	task := Task{ID: nextID, Title: title, Description: description, Tag: tag, Completed: completed}
	tasks = append(tasks, task)
	nextID++

	err := a.SaveTasks()
	if err != nil {
		fmt.Println("Error saving tasks after adding a new task:", err)
	}
	return task
}

func (a *App) SaveTasks() error {
	fmt.Println("Saving tasks...", tasks)
	file, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error marshalling tasks:", err)
		return err
	}
	err = os.WriteFile("tasks.json", file, 0644)
	if err != nil {
		fmt.Println("Error  writing file:", err)
	}
	return err
}

func (a *App) LoadTasks() error {
	var loadedTasks []Task
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		// If the file does not exist, it's not necessarily an error that should stop the application.
		// It could simply mean that no tasks have been saved yet.
		// However, log or handle other kinds of errors as needed.
		if !os.IsNotExist(err) {
			fmt.Println("Error reading tasks from file:", err)
			return err
		}
		return nil // No tasks to load, but not an error.
	}
	err = json.Unmarshal(file, &loadedTasks)
	if err != nil {
		fmt.Println("Error unmarshalling tasks from file:", err)
		return err
	}

	// Mutex lock to ensure thread-safe operation
	mutex.Lock()
	defer mutex.Unlock()

	// If loading was successful, update the global tasks slice with the loaded tasks.
	tasks = loadedTasks

	// Determine the next ID for new tasks based on loaded tasks.
	// This ensures that new task IDs will not clash with existing ones.
	if len(tasks) > 0 {
		maxID := tasks[0].ID
		for _, task := range tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
		nextID = maxID + 1
	}

	return nil
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := a.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks on startup:", err)
		// Depending on your requirements, you might want to handle this error more gracefully.
		// For instance, if the file doesn't exist, that's fine for the first run of the app.
		// So, you could check for a specific error type here if needed.
		return
	}
}
