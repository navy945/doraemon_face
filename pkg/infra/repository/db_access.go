package repository

import (
	"database/sql"
	"go/pkg/domain"
)

//DBアクセスのCRUD処理をここで実装
//DBへのクエリ操作は外部に依存するため、infra層で定義しなくてはならない

type DBConnector struct {
	dbPool *sql.DB
}

func NewDbAccess(dbPool *sql.DB) *DBConnector {
	return &DBConnector{
		dbPool: dbPool,
	}
}

// Get a Task
func (r *DBConnector) Get(id int) (*domain.Task, error) {
	stmt, err := r.dbPool.Prepare(`select id, name from t_task where id = $1;`)
	if err != nil {
		return nil, err
	}
	var task domain.Task
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		_ = rows.Scan(&task.ID, &task.Name)
	}
	return &task, nil
}

// List Tasks
func (r *DBConnector) List() ([]*domain.Task, error) {
	stmt, err := r.dbPool.Prepare(`select id, name from t_task;`)
	if err != nil {
		return nil, err
	}
	var tasks []*domain.Task
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task domain.Task
		err = rows.Scan(&task.ID, &task.Name)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (r *DBConnector) Create(t *domain.Task) (int, error) {
	stmt, err := r.dbPool.Prepare("insert into t_task (id, name) values ($1, $2);")
	if err != nil {
		return t.ID, err
	}
	_, err = stmt.Exec(t.ID, t.Name)
	if err != nil {
		return t.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return t.ID, err
	}
	return t.ID, nil
}

func (r *DBConnector) Update(t *domain.Task) error {
	//t.UpdatedAt = time.Now()
	_, err := r.dbPool.Exec("update t_task set name = $1 where id = $2", t.Name, t.ID)
	if err != nil {
		return err
	}
	return nil
}
