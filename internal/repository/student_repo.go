
package repository

import (
    "go-admin/internal/model"
    "gorm.io/gorm"
)

type StudentRepository struct {
    DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
    return &StudentRepository{DB: db}
}

func (r *StudentRepository) CreateStudent(student *model.Student) (*model.Student, error) {
    if err := r.DB.Create(student).Error; err != nil {
        return nil, err
    }
    return student, nil
}

func (r *StudentRepository) GetAllStudents() ([]model.Student, error) {
    var students []model.Student
    if err := r.DB.Find(&students).Error; err != nil {
        return nil, err
    }
    return students, nil
}
