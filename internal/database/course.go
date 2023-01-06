package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	Id          string
	Name        string
	Description string
	Category    *Category
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, category *Category) (*Course, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?)",
		id, name, description, category.Id)
	if err != nil {
		return nil, err
	}

	course := Course{Id: id, Name: name, Description: description, Category: category}
	return &course, nil
}

func (c *Course) FindAll() ([]Course, error) {
	//TODO: ao invés de recuperar automaticamente a prop relacionada, é melhor declarar a entidade como um model do graphql
	//TODO: para deixar o graphql resolver a prop relacionada quando ela for solicitada nas consultas.
	rows, err := c.db.Query("SELECT co.id, co.name, co.description, ca.id, ca.name, ca.description FROM courses co JOIN categories ca ON co.category_id = ca.id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var courseId, courseName, courseDesc, categoryId, categoryName, categoryDesc string
		if err := rows.Scan(&courseId, &courseName, &courseDesc, &categoryId, &categoryName, &categoryDesc); err != nil {
			return nil, err
		}
		newCat := Category{
			Id:          categoryId,
			Name:        categoryName,
			Description: categoryDesc,
		}
		newCourse := Course{
			Id:          courseId,
			Name:        courseName,
			Description: courseDesc,
			Category:    &newCat,
		}

		courses = append(courses, newCourse)
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{Id: id, Name: name, Description: description, Category: &Category{Id: categoryID}})
	}
	return courses, nil
}
