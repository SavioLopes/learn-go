package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskService struct {
	DB          *sql.DB
	TaskChannel chan Task
}

func (t *TaskService) AddTask(ts *Task) error {
	query := "INSERT INTO tasks (title, description, status, created_at) VALUES (?, ?, ?, ?)"
	result, err := t.DB.Exec(query, ts.Title, ts.Description, ts.Status, ts.CreatedAt)
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	ts.ID = int(id)
	return err
}

func (t *TaskService) UpdateTaskStatus(ts Task) error {
	query := "UPDATE tasks SET status = ? WHERE id = ?"
	_, err := t.DB.Exec(query, ts.Status, ts.ID)
	return err
}

func (t *TaskService) ListTasks() ([]Task, error) {
	rows, err := t.DB.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *TaskService) ProcessTasks() {
	for task := range t.TaskChannel {
		log.Printf("Processing task: %s", task.Title)
		time.Sleep(time.Second * 5)
		task.Status = "completed"
		t.UpdateTaskStatus(task)
		log.Printf("Task %s processed", task.Title)
	}
}

func (t *TaskService) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Task %s processed", task.Title)
	task.Status = "pending"
	task.CreatedAt = time.Now()
	err = t.AddTask(&task)
	if err != nil {
		http.Error(w, "Error adding task", http.StatusInternalServerError)
		return
	}

	t.TaskChannel <- task
	w.WriteHeader(http.StatusCreated)
}

func (t *TaskService) HandleListSTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.ListTasks()
	if err != nil {
		http.Error(w, "Error listing task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	taskChannel := make(chan Task)
	taskService := TaskService{
		DB:          db,
		TaskChannel: taskChannel,
	}

	go taskService.ProcessTasks()

	http.HandleFunc("POST /tasks", taskService.HandleCreateTask)
	http.HandleFunc("GET /tasks", taskService.HandleListSTasks)

	log.Printf("Server running on port 8081")
	http.ListenAndServe(":8081", nil)
}
