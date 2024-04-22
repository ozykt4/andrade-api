package service

import (
	"github.com/ozykt4/andrade-api/internal/model"
	"github.com/ozykt4/andrade-api/internal/repository"
)

type IUserService interface {
	CreateUser(userReq *model.UserReq) *model.Response
	FindAllUsers() *model.Response
	FindUserByID(id uint) *model.Response
	UpdateUser(id uint, userReq *model.UserReq) *model.Response
	DeleteUser(id uint) *model.Response
}

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) CreateUser(userReq *model.UserReq) *model.Response {
	user := userReq.ToUser()

	createUser, err := us.repo.Create(user)
	if err != nil {
		return model.NewErrorResponse(err, 500)
	}

	return model.NewSuccessResponse(createUser.ToUserRes())
}

func (us *UserService) FindAllUsers() *model.Response {
	users, err := us.repo.FindAll()
	if err != nil {
		return model.NewErrorResponse(err, 500)
	}

	userRes := []*model.UserRes{}
	for _, u := range users {
		userRes = append(userRes, u.ToUserRes())
	}

	return model.NewSuccessResponse(userRes)
}

func (us *UserService) FindUserByID(id uint) *model.Response {
	user, err := us.repo.FindByID(id)
	if err != nil {
		return model.NewErrorResponse(err, 404)
	}

	return model.NewSuccessResponse(user.ToUserRes())
}

func (us *UserService) UpdateUser(id uint, userReq *model.UserReq) *model.Response {
	user, err := us.repo.FindByID(id)
	if err != nil {
		return model.NewErrorResponse(err, 404)
	}

	user.Name = userReq.Name
	user.Email = userReq.Email

	updateUSer, err := us.repo.Update(user)
	if err != nil {
		return model.NewErrorResponse(err, 500)
	}

	return model.NewSuccessResponse(updateUSer.ToUserRes())
}

func (us *UserService) DeleteUser(id uint) *model.Response {
	_, err := us.repo.Delete(id)
	if err != nil {
		return model.NewErrorResponse(err, 404)
	}

	return model.NewSuccessResponse(nil)
}
