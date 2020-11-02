package usecase

import (
	"database/sql"
	"fmt"
	"net/http"

	. "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/consts"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/models"
	. "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/error_response"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/user"
)

type UserUsecase struct {
	userRep user.UserRep
}

func NewUserUsecase(userRep user.UserRep) *UserUsecase {
	return &UserUsecase{
		userRep: userRep,
	}
}

func (uUc *UserUsecase) CreateUser(user *models.User) *ErrorResponse { //TODO: переименовать в просто Create
	exists, err := uUc.checkEmailExists(user.Email)

	if err != nil {
		return NewErrorResponse(ErrInternal, err)
	}
	if exists {
		return NewErrorResponse(ErrEmailAlreadyExist, nil)
	}

	exists, err = uUc.checkNameExists(user.Name)

	if err != nil {
		return NewErrorResponse(ErrInternal, err)
	}
	if exists {
		return NewErrorResponse(ErrNameAlreadyExist, nil)
	}

	if err := uUc.userRep.Insert(user); err != nil {
		return NewErrorResponse(ErrInternal, err)
	}

	return nil
}

func (uUc *UserUsecase) GetByEmail(email string) (*models.User, *ErrorResponse) {
	user, err := uUc.userRep.SelectByEmail(email)

	if err == sql.ErrNoRows {
		return nil,
			NewCustomErrorResponse(http.StatusNotFound, err, fmt.Sprintf("User with email %s not found", email))
	}

	if err != nil {
		return nil, NewErrorResponse(ErrInternal, err)
	}

	return user, nil
}

func (uUc *UserUsecase) GetByName(name string) (*models.User, *ErrorResponse) {
	user, err := uUc.userRep.SelectByName(name)

	if err == sql.ErrNoRows {
		return nil,
			NewCustomErrorResponse(http.StatusNotFound, err, fmt.Sprintf("User with name %s not found", name))
	}

	if err != nil {
		return nil, NewErrorResponse(ErrInternal, err)
	}

	return user, nil
}

func (uUc *UserUsecase) LoginUser(login string, password string) (*models.User, *ErrorResponse) { //TODO: переименовать в просто Login
	user, err := uUc.userRep.SelectByLoginAndPassword(login, password)

	if err == sql.ErrNoRows {
		return nil, NewErrorResponse(ErrIncorrectLoginOrPassword, err)
	}

	if err != nil {
		return nil, NewErrorResponse(ErrInternal, err)
	}

	return user, nil
}

func (uUc *UserUsecase) GetByID(id uint64) (*models.User, *ErrorResponse) {
	user, err := uUc.userRep.SelectByID(id)

	if err == sql.ErrNoRows {
		return nil, NewErrorResponse(ErrNotAuthorized, err)
	}

	if err != nil {
		return nil, NewErrorResponse(ErrInternal, err)
	}

	return user, nil
}

func (uUc *UserUsecase) UpdateProfile(id uint64) *ErrorResponse {
	user, errResp := uUc.GetByID(id)
	if errResp != nil {
		return errResp
	}

	existingUser, err := uUc.checkNameOrEmailExists(user.Name, user.Email, user.ID)
	if err != nil {
		return NewErrorResponse(ErrInternal, err)
	}
	if existingUser != nil {
		if existingUser.Name == user.Name {
			return NewErrorResponse(ErrNameAlreadyExist, nil)
		} else {
			return NewErrorResponse(ErrEmailAlreadyExist, nil)
		}
	}

	if err := uUc.userRep.Update(user); err != nil {
		return NewErrorResponse(ErrInternal, err)
	}

	return nil
}

func (uUc *UserUsecase) UpdatePassword(id uint64) *ErrorResponse {
	user, errResp := uUc.GetByID(id)
	if errResp != nil {
		return errResp
	}

	if err := uUc.userRep.Update(user); err != nil {
		return NewErrorResponse(ErrInternal, err)
	}

	return nil
}

func (uUc *UserUsecase) checkEmailExists(email string) (bool, error) {
	_, err := uUc.userRep.SelectByEmail(email)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (uUc *UserUsecase) checkNameExists(name string) (bool, error) {
	_, err := uUc.userRep.SelectByName(name)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (uUc *UserUsecase) checkNameOrEmailExists(name string, email string, id uint64) (*models.User, error) {
	user, err := uUc.userRep.SelectByNameOrEmail(name, email, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
