package main

import (
	"database/sql"
	"time"
)

type DBManager struct{
	db *sql.DB
}

func NewDBManager(db *sql.DB) * DBManager{
	return &DBManager{db : db}
}

func (b *DBManager) CreateUser(user *User) (*User, error){
	time := time.Now()
	query := `
		INSERT INTO users2(name, age, created_at)
		VALUES($1,$2,$3)
		RETURNING id, name, age, created_at, updated_at, deleted_at
	`

	row := b.db.QueryRow(query,
		user.Name,
		user.Age,
		time,
	)

	var result User

	err := row.Scan(
		&result.Id,
		&result.Name,
		&result.Age,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	if err != nil{
		return nil, err
	}
	return &result, nil
}

func (b *DBManager) UpdateUser(user *User) (*User, error){
	updatedTime := time.Now()

	query := `
		UPDATE users2 SET
		name=$1,
		age=$2,
		updated_at=$3
		where id=$4
		RETURNING id, name, age, created_at, updated_at, deleted_at
	`
	row := b.db.QueryRow(query,
		user.Name,
		user.Age,
		updatedTime,
		user.Id,
	)

	var result User

	err := row.Scan(
		&result.Id,
		&result.Name,
		&result.Age,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	if err != nil{
		return nil, err
	}

	return &result, nil
}


func (b *DBManager) Get(id int) (*User, error){
	query := `
		select id, name, age, created_at from users2 where id=$1
	`

	row := b.db.QueryRow(
		query,
		id,
	)
	var result User

	err := row.Scan(
		&result.Id,
		&result.Name,
		&result.Age,
		&result.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}