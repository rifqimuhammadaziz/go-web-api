package book

import (
	"fmt"
)

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) FindAll() ([]Book, error) {
	var books []Book
	fmt.Println("FindAll function") // example for find all data
	return books, nil
}

func (r *fileRepository) FindById(ID int) (Book, error) {
	var book Book
	fmt.Println("FindByID function") // example for find by id
	return book, nil
}

func (r *fileRepository) Create(book Book) (Book, error) {
	fmt.Println("Create function") // example for create data
	return book, nil
}
