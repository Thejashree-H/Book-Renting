package repositories
import (
	"database/sql"
	"bookrental/models"
	"net/http"
)

type UsersRepository struct {
	dbHandler *sql.DB
}

func NewUsersRepository(dbHandler *sql.DB) *UsersRepository {
	return &UsersRepository{dbHandler: dbHandler}
}

func (ur UsersRepository) CreateUser(user *models.User) (*models.User, *models.ResponseError) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := ur.dbHandler.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	return user, nil
}

func (ur *UsersRepository) UpdateUser(user *models.User) (*models.User, *models.ResponseError) {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := ur.dbHandler.Exec(query, user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	return user, nil
}

func (ur *UsersRepository) DeleteUser(id string) *models.ResponseError {
	query := "DELETE FROM users WHERE id = ?"
	_, err := ur.dbHandler.Exec(query, id)
	if err != nil {
		return &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	return nil
}

func (ur *UsersRepository) GetUser(id string) (*models.User, *models.ResponseError) {
	query := "SELECT id, name, email, password FROM users WHERE id = ?"
	var user models.User
	err := ur.dbHandler.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	return &user, nil
}

func (ur *UsersRepository) GetUsersBatch() ([]models.User, *models.ResponseError) {
	query := "SELECT id, name, email, password FROM users"
	rows, err := ur.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, &models.ResponseError{Message: err.Error(), Status: http.StatusInternalServerError}
		}
		users = append(users, user)
	}

	return users, nil
}
