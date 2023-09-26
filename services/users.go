package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/indahouse-aleksandr/mp_backend/errs"
	"github.com/indahouse-aleksandr/mp_backend/inout"
	"github.com/indahouse-aleksandr/mp_backend/queries"
	"github.com/indahouse-aleksandr/mp_backend/repositories"
)

const UserDefaultRole = "default_role"
const UserDefaultPhoto = "default_photo"

type UserService struct {
	context    *gin.Context
	repository *repositories.Pgx
}

func NewUserService(c *gin.Context, r *repositories.Pgx) *UserService {
	return &UserService{
		context:    c,
		repository: r,
	}
}

func (t *UserService) CreateUser(in inout.InputCreateUser) (queries.User, error) {
	user, err := t.GetUserByEmail(in)
	if err != nil && err != sql.ErrNoRows {
		return user, err
	}

	if user.Email == in.Email {
		err := errs.NewUserEmaiExist("email", errs.UserEmaiExistMsg)
		return queries.User{}, fmt.Errorf("email check: %w", &err)
	}

	arg := queries.CreateUserParams{
		Name:      in.Name,
		Email:     in.Email,
		Photo:     UserDefaultPhoto,
		Verified:  false,
		Password:  in.Password,
		Role:      UserDefaultRole,
		UpdatedAt: time.Now(),
	}

	return t.repository.Q.CreateUser(t.context, arg)
}

func (t *UserService) GetUserByEmail(in inout.InputCreateUser) (queries.User, error) {
	return t.repository.Q.GetUserByEmail(t.context, in.Email)
}
