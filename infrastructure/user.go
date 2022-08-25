package infrastructure

import (
	"database/sql"
	"go_api_sample/domain"
)

type UserRepositoryInfrastructure struct {
	*sql.DB
}

func NewUserRepositoryInfrastructure(db *sql.DB) domain.UserRepository {
	return &UserRepositoryInfrastructure{db}
}

func (r *UserRepositoryInfrastructure) Find(id int) (*domain.User, error) {
	row := r.QueryRow(`SELECT * FROM users WHERE id = ?`, id)

	user, err := toStructure(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryInfrastructure) Create(user *domain.User) error {
	_, err := r.Exec(`INSERT INTO users(name, created_at, updated_at) values(?, ?, ?)`, user.Name, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryInfrastructure) Update(user *domain.User) error {
	_, err := r.Exec(`UPDATE users SET name = ?, updated_at = ? WHERE id = ?`, user.Name, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryInfrastructure) Delete(id int) error {
	_, err := r.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}

func toStructure(row *sql.Row) (*domain.User, error) {
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
