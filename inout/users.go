package inout

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/indahouse-aleksandr/mp_backend/errs"
	"github.com/indahouse-aleksandr/mp_backend/queries"
)

type InputCreateUser struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"email,required"`
	Password string `json:"password" binding:"required"`
}

type OutUser struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Photo     string    `json:"photo"`
	Verified  bool      `json:"verified"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OutCreateUserErr struct {
	IsError bool               `json:"is_error"`
	Errors  []OutValidatorsErr `json:"errors"`
}

type OutCreateUserOk struct {
	IsError bool    `json:"is_error"`
	Data    OutUser `json:"data"`
}

func ErrMsgUser(err error) OutCreateUserErr {

	var target *errs.ErrorUserEmaiExist
	if errors.As(err, &target) {
		return OutCreateUserErr{
			IsError: true,
			Errors: []OutValidatorsErr{{
				Field:   target.Field(),
				Message: target.Error(),
			}},
		}
	}

	return OutCreateUserErr{}
}

func CreateUserOk(user queries.User) OutCreateUserOk {

	var outUser OutUser
	u, _ := json.Marshal(user)
	json.Unmarshal(u, &outUser)

	return OutCreateUserOk{
		IsError: false,
		Data:    outUser,
	}
}
