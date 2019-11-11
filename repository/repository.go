package repository

import (
	"database/sql"

	"github.com/michalsz/sql_mock_example/models"
)

// UserRepository struct for store DB
type UserRepository struct {
	DB *sql.DB
}

//Save an user
func (repo *UserRepository) Save(user *models.User) (sql.Result, error) {
	sqlQuery, err := repo.DB.Prepare("INSERT INTO users(user.Name) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}

	queryResult, err := sqlQuery.Exec(user.Name)

	return queryResult, err
}
