package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	Id          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (*Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)", id, name, description)
	if err != nil {
		return nil, err
	}

	return &Category{Id: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{Id: id, Name: name, Description: description})
	}
	return categories, nil
}

func (c *Category) FindById(categoryId string) (*Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories WHERE id=?", categoryId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// TODO: o que acontece se nao vier resultados?
	rows.Next()
	var id, name, description string
	if err := rows.Scan(&id, &name, &description); err != nil {
		return nil, err
	}
	category := Category{Id: id, Name: name, Description: description}
	return &category, nil
}
