package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
)

type TeacherRepository interface {
	FetchAll() ([]model.Teacher, error)
	FetchByID(id int) (*model.Teacher, error)
	Store(g *model.Teacher) error
}

type teacherRepoImpl struct {
	db *sql.DB
}

func NewTeacherRepo(db *sql.DB) *teacherRepoImpl {
	return &teacherRepoImpl{db}
}

func (g *teacherRepoImpl) FetchAll() ([]model.Teacher, error) {

	rows, err := g.db.Query("SELECT * FROM teachers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []model.Teacher
	for rows.Next() {
		var teacher model.Teacher
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Address, &teacher.Subject)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil // TODO: replace this
}

func (g *teacherRepoImpl) FetchByID(id int) (*model.Teacher, error) {
	row := g.db.QueryRow("SELECT id, name, address, subject FROM teachers WHERE id = $1", id)

	var teacher model.Teacher
	err := row.Scan(&teacher.ID, &teacher.Name, &teacher.Address, &teacher.Subject)
	if err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (g *teacherRepoImpl) Store(teacher *model.Teacher) error {
	_, err := g.db.Exec("INSERT INTO teachers(name, address, subject) VALUES ($1, $2, $3)",
		teacher.Name, teacher.Address, teacher.Subject)
	if err != nil {
		return err
	}

	return nil// TODO: replace this
}
