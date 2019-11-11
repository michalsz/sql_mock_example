package repository_test

import (
	"testing"

	"github.com/michalsz/sql_mock_example/models"
	"github.com/michalsz/sql_mock_example/repository"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT INTO users"
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs("michal").
		WillReturnResult(sqlmock.NewResult(1, 1))

	userRepo := repository.UserRepository{DB: db}
	user := models.User{Name: "michal"}

	result, err := userRepo.Save(&user)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, result)
}

func TestSaveNoName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT INTO users"
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs("").
		WillReturnResult(sqlmock.NewResult(1, 1))

	userRepo := repository.UserRepository{DB: db}
	user := models.User{}

	result, err := userRepo.Save(&user)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, result)
}
