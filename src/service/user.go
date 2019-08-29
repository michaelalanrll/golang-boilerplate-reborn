package services

import (
	"fmt"
	"sync"
	"time"
	"github.com/jinzhu/copier"
	httpEntity "example_app/entity/http"
	dbEntity "example_app/entity/db"
	APIEntity "example_app/entity/api"
	repository "example_app/repository/db"
	repositoryAPI "example_app/repository/api"
)

type UserService struct {
	userRepository repository.UserRepositoryInterface
	userRepositoryAPI repositoryAPI.FriendAPIRepositoryInterface
}

func UserServiceHandler() *UserService {
	return &UserService{
		userRepository: repository.UserRepositoryHandler(),
		userRepositoryAPI: repositoryAPI.ThirdPartyAPIHandler(),
	}
}

type UserServiceInterface interface {
	GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse
	GetAllUser(page int,count int) []httpEntity.UserResponse
	UpdateUserByID(id int, payload httpEntity.UserRequest) bool
	StoreUser(payload httpEntity.UserRequest) bool
	DeleteUser(id int) *httpEntity.UserResponse
}

func (service *UserService) GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse{
	waitGroup.Add(1)
	user := &dbEntity.User{}
	go service.userRepository.GetUserByID(id,user,waitGroup)

	waitGroup.Add(1)
	friend := &APIEntity.FriendResponse{}
	go service.userRepositoryAPI.GetFriendID(id,friend,waitGroup)
	
	waitGroup.Wait()

	result := &httpEntity.UserDetailResponse{}
	if user != nil {
		copier.Copy(result, user)
		if nil != user.UserStatus{
			result.Status = &user.UserStatus.Name
		}
		if friend != nil {
			result.Avatar = &friend.Data.Avatar
		}
	}
	return result
}

func (service *UserService) GetAllUser(page int,count int) []httpEntity.UserResponse {
	users, _ := service.userRepository.GetUsersList(page,count)
	result := []httpEntity.UserResponse{}
	copier.Copy(&result, &users)
	return result
}

func (service *UserService) UpdateUserByID(id int, payload httpEntity.UserRequest) bool {
	now := time.Now()
	user := &dbEntity.User{
		Name: payload.Name,
		IDCardNumber: payload.IDCardNumber,
		Address: payload.Address,
		UserStatusID: payload.UserStatusId,
		UpdatedAt: &now,
	}
	err := service.userRepository.UpdateUserByID(id, user)
	if nil != err {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (service *UserService) StoreUser(payload httpEntity.UserRequest) bool {
	now := time.Now()
	user := &dbEntity.User{
		Name: payload.Name,
		IDCardNumber: payload.IDCardNumber,
		Address: payload.Address,
		UserStatusID: payload.UserStatusId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := service.userRepository.StoreUser(user)
	if nil != err {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (service *UserService) DeleteUser(id int) *httpEntity.UserResponse{
	user := dbEntity.User{}
	result := service.userRepository.DeleteUser(id, &user)

	output := &httpEntity.UserResponse{}
	if result == nil {
		copier.Copy(output, user)
	}
	return output
}