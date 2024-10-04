package services
import (
	"bookrental/models"
	"bookrental/repositories"
	"regexp"
	"net/http"
)

type UsersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersService(usersRepository *repositories.UsersRepository) *UserService {
	return &UsersService{usersRepository: usersRepository}
}

func (us *UsersService) CreateUser(user *models.User) (*models.User, *models.responseError) {
	if err := validateUser(user); err != nil {
		return nil, err
	}
	return us.usersRepository.CreateUser(user)
}

func (us *UserService) UpdateUser(user *models.User) (*models.User, *models.responseError) {
	return us.usersRepository.UpdateUser(user)
}

func (us *UserService) DeleteUser(id string) *models.responseError {
	return us.usersRepository.DeleteUser(id)
}

func (us *UsersService) GetUser(id string) (*models.User, *models.ResponseError) {
	return us.usersRepository.GetUser(id)
}

func (us *UsersService) GetUsersBatch() ([]models.User, *models.ResponseError) {
	return us.usersRepository.GetUsersBatch()
}

func validateUser(user *models.User) *models.ResponseError {
	if user.Name == "" {
		return &models.ResponseError{
			Message: "Name cannot be empty",
			Status:  http.StatusBadRequest,
		}
	}

	if user.Email == "" {
		return &models.ResponseError{
			Message: "Email cannot be empty",
			Status:  http.StatusBadRequest,
		}
	}

	if !isValidEmail(user.Email) {
		return &models.ResponseError{
			Message: "Invalid email format",
			Status:  http.StatusBadRequest,
		}
	}

	if user.Password == "" {
		return &models.ResponseError{
			Message: "Password cannot be empty",
			Status:  http.StatusBadRequest,
		}
	}

	if len(user.Password) < 6 {
		return &models.ResponseError{
			Message: "Password must be at least 6 characters long",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}