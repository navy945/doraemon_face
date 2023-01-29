package task

import "go/pkg/domain"

type Service struct {
	//repositoryで定義
	//reader,writerを保持
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetAllTask() ([]*domain.Task, error) {
	//service→repositoryのList()→infraのList()を呼び出す
	//usecase層からinfra層を呼び出すため
	tasks, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *Service) GetTask(id int) (*domain.Task, error) {
	task, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Service) CreateTask(id int, name string) (int, error) {
	//Task構造体を設定値で定義
	task, err := domain.NewTask(id, name)
	if err != nil {
		return id, err
	}
	return s.repo.Create(task)
}

func (s *Service) UpdateTask(t *domain.Task) error {
	return s.repo.Update(t)
}
