package task

import "go/pkg/domain"

//taskに関わるインターフェースを定義

// Reader interface
type Reader interface {
	Get(id int) (*domain.Task, error)
	//Search(query string) ([]*domain.Task, error)
	List() ([]*domain.Task, error)
}

// Writer book writer
type Writer interface {
	Create(e *domain.Task) (int, error)
	Update(e *domain.Task) error
	//Delete(id int) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
// 実装はインフラ層
type UseCase interface {
	GetTask(id int) (*domain.Task, error)
	//SearchBooks(query string) ([]*domain.Task, error)
	GetAllTask() ([]*domain.Task, error)
	CreateTask(id int, name string) (int, error)
	UpdateTask(*domain.Task) error
	//DeleteBook(id int) error
}
