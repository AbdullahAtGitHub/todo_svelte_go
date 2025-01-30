package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

type task struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Completed bool      `db:"completed"`
	CreatedAt time.Time `db:"created_at"`
}

func main() {

	//connecting to postgres db container
	var err error
	db, err = sqlx.Connect("postgres", "user=postgres dbname=todo sslmode=disable password=todo_123 host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected with todo db.")
	}

	// Setup Gin router
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/tasks", getTasks)
	r.POST("/tasks", createTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	// Run on port 8080
	r.Run(":8000")
}

func getTasks(c *gin.Context) {

	var tasks []task

	err := db.Select(&tasks, "Select id,title,completed,created_at from tasks;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var t task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err := db.Exec("INSERT INTO TASKS (title,completed) VALUES ($1,$2)", t.Title, t.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func updateTask(c *gin.Context) {
	var t task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err := db.Exec("UPDATE TASKS SET title=$1,completed=$2 WHERE id=$3",
		t.Title, t.Completed, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func deleteTask(c *gin.Context) {
	_, err := db.Exec("DELETE FROM TASKS WHERE id=$1", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
