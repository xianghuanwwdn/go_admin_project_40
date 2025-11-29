
package service

import (
    "go-admin/internal/model"
    "go-admin/internal/repository"
)

type StudentService struct {
    repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *StudentService {
    return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(name string, age int, class, gender string) (*model.Student, error) {
    student := &model.Student{Name: name, Age: age, Class: class, Gender: gender}
    return s.repo.CreateStudent(student)
}

func (s *StudentService) GetAllStudents() ([]model.Student, error) {
    return s.repo.GetAllStudents()
}
