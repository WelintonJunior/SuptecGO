package repository

import (
	"github.com/WelintonJunior/SuptecGO/src/server/adapters/domain"
	db "github.com/WelintonJunior/SuptecGO/src/server/adapters/repository/database"
	"golang.org/x/crypto/bcrypt"
)

type UsuRepository interface {
	NewUsu(u domain.USU_users) error
}

type localUsuRepository struct{}

func NewLocalUsuRepository() *localUsuRepository {
	return &localUsuRepository{}
}

func (r *localUsuRepository) NewUsu(u domain.USU_users) error {
	query := `insert into usu_users (usu_name, usu_email, usu_password, usu_level) values ($1, $2, $3, $4)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Usu_password), 14)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Usu_name, u.Usu_email, hashed, u.Usu_level)

	if err != nil {
		return err
	}

	return nil
}
