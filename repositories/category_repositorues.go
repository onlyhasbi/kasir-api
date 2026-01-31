package repositories

import (
	"cashier/models"
	"database/sql"
	"errors"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAll() ([]models.Category, error) {
	query := "SELECT id, name, description FROM category"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	categorys := make([]models.Category, 0)

	for rows.Next() {
		var p models.Category
		if err := rows.Scan(&p.ID, &p.Name, &p.Description); err != nil {
			return nil, err
		}
		categorys = append(categorys, p)
	}

	return categorys, nil
}

func (repo *CategoryRepository) Create(category *models.Category) error {
	query := "INSERT INTO categorys (name, description) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	return err
}

func (repo *CategoryRepository) GetByID(id int) (*models.Category, error) {
	query := "SELECT id, name, description FROM categorys WHERE id = $1"

	var p models.Category
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description)

	if err == sql.ErrNoRows {
		return nil, errors.New("Category not found")
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *CategoryRepository) Update(category *models.Category) error {
	query := "UPDATE categorys SET name = $1, price = $2, stock = $3 WHERE id = $4"
	result, err := repo.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Category not found")
	}

	return nil
}

func (repo *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM category WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if rows == 0 {
		return errors.New("Category not found")
	}

	return nil
}
