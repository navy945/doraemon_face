package domain

// DBカラムに対応した構造体
type Task struct {
	ID   int
	Name string
}

// JSON形式対応（API上やりとりする）構造体
type Jtask struct {
	ID   int
	Name string
}

// NewBook create a new book
func NewTask(id int, name string) (*Task, error) {
	b := &Task{
		ID:   id,
		Name: name,
	}
	return b, nil
}
