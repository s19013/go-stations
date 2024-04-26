package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
		// confirmで指定しているカラムとQueryRow,QueryRowContextの引数をあわせないとエラーになるので注意
	)

	// prepareを使ってinjectionとかを回避する
	preparedInsert, err := s.db.PrepareContext(ctx, insert)
	if err != nil {
		return nil, err
	}

	// PrepareContext,Prepareを使う
	defer preparedInsert.Close()

	// データベースに登録
	result, err := preparedInsert.ExecContext(ctx, subject, description)
	if err != nil {
		return nil, err
	}

	// 今保存したデータのIDを取り出す
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	preparedConfirm, err := s.db.PrepareContext(ctx, confirm)
	if err != nil {
		return nil, err
	}

	defer preparedConfirm.Close()

	// IDを元にtodoのデータを取り出す
	row := preparedConfirm.QueryRowContext(ctx, lastID)

	var todo model.TODO
	todo.ID = int(lastID)

	err = row.Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &todo, err
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	if prevID == 0 {
		TODOs, err := withOutPrevID(s, ctx, read, size)
		if err != nil {
			log.Print("err", err)
			return nil, err
		}
		return TODOs, nil
	}

	TODOs, err := withPrevID(s, ctx, read, prevID, size)
	if err != nil {
		log.Print("err", err)
		return nil, err
	}
	return TODOs, nil

}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	// prepareを使ってinjectionとかを回避する
	preparedUpdate, err := s.db.PrepareContext(ctx, update)
	if err != nil {
		log.Panicln("err", err)
		return nil, err
	}

	// PrepareContext,Prepareを使う
	defer preparedUpdate.Close()

	// 更新
	result, err := preparedUpdate.ExecContext(ctx, subject, description, id)
	if err != nil {
		return nil, err
	}
	// resultいらないんだけどなんか使わないとエラーになるからここで適当につかう
	result.LastInsertId() //更新なので0

	preparedConfirm, err := s.db.PrepareContext(ctx, confirm)
	if err != nil {
		log.Panicln("err", err)
		return nil, err
	}

	defer preparedConfirm.Close()

	// IDを元にtodoのデータを取り出す
	row := preparedConfirm.QueryRowContext(ctx, id)

	var todo model.TODO
	todo.ID = int(id)

	err = row.Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		log.Panicln("err", err)
		return nil, err
	}

	return &todo, err
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}

// 1つだけほしいならQueryRowContext
// 複数返してほしいならQueryContext
func withOutPrevID(s *TODOService, ctx context.Context, sql string, size int64) ([]*model.TODO, error) {
	preparedQuery, err := s.db.PrepareContext(ctx, sql)
	if err != nil {
		log.Print("err", err)
		return nil, err
	}

	defer preparedQuery.Close()

	rows, err := preparedQuery.QueryContext(ctx, size)
	if err != nil {
		return nil, err
	}

	var TODOs []*model.TODO

	for rows.Next() {
		todo := &model.TODO{}

		err := rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// 配列に追加
		TODOs = append(TODOs, todo)
	}

	return TODOs, nil
}

func withPrevID(s *TODOService, ctx context.Context, sql string, PrevID, size int64) ([]*model.TODO, error) {
	preparedQuery, err := s.db.PrepareContext(ctx, sql)
	if err != nil {
		log.Print("err", err)
	}

	defer preparedQuery.Close()

	rows, err := preparedQuery.QueryContext(ctx, PrevID, size)
	if err != nil {
		return nil, err
	}

	var TODOs []*model.TODO

	for rows.Next() {
		var todo *model.TODO

		err := rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// 配列に追加
		TODOs = append(TODOs, todo)
	}

	return TODOs, err
}
