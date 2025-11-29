
package main

import (
    "go-admin/internal/model"
    "go-admin/internal/repository"
    "go-admin/internal/service"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // Initialize database
    db, err := gorm.Open(sqlite.Open("school.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }
    db.AutoMigrate(&model.Student{})

    // Repositories
    studentRepo := repository.NewStudentRepository(db)

    // Services
    studentService := service.NewStudentService(studentRepo)

    // Set up Gin
    r := gin.Default()

    r.POST("/students", func(c *gin.Context) {
        var student model.Student
        if err := c.ShouldBindJSON(&student); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        createdStudent, err := studentService.CreateStudent(student.Name, student.Age, student.Class, student.Gender)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }

        c.JSON(200, createdStudent)
    })

    r.GET("/students", func(c *gin.Context) {
        students, err := studentService.GetAllStudents()
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, students)
    })

    r.Run(":8080")
}
