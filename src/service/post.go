package services

import (
	"fmt"
	"time"
	"github.com/jinzhu/copier"
	httpEntity "example_app/entity/http"
	dbEntity "example_app/entity/db"
	repository "example_app/repository/db"
)

type PostService struct {
	postRepository repository.PostRepositoryInterface
}

func PostServiceHandler() *PostService {
	return &PostService{
		postRepository: repository.PostRepositoryHandler(),
	}
}

type PostServiceInterface interface {
	GetPostById(id int) *httpEntity.PostDetailResponse
	GetAllPosts(page int,count int) []httpEntity.PostResponse
	UpdatePostById(id int, payload httpEntity.PostRequest) bool
	StorePost(payload httpEntity.PostRequest) bool
	DeletePost(id int) *httpEntity.PostResponse
}

func (service *PostService) GetPostById(id int) *httpEntity.PostDetailResponse{
	post := &dbEntity.Post{}
	service.postRepository.GetPostById(id, post)

	result := &httpEntity.PostDetailResponse{}
	if post != nil {
		copier.Copy(result, post)
	}
	return result
}

func (service *PostService) GetAllPosts(page int, count int) []httpEntity.PostResponse {
	posts, _ := service.postRepository.GetPostsList(page,count)
	result := []httpEntity.PostResponse{}
	copier.Copy(&result, &posts)
	return result
}

func (service *PostService) UpdatePostById(id int, payload httpEntity.PostRequest) bool {
	now := time.Now()
	post := &dbEntity.Post{
		Title: payload.Title,
		Body: payload.Body,
		UserId: payload.UserId,
		UpdatedAt: &now,
	}
	err := service.postRepository.UpdatePostById(id, post)
	if nil != err {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (service *PostService) StorePost(payload httpEntity.PostRequest) bool {
	now := time.Now()
	post := &dbEntity.Post{
		Title: payload.Title,
		Body: payload.Body,
		UserId: payload.UserId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := service.postRepository.StorePost(post)
	if nil != err {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (service *PostService) DeletePost(id int) *httpEntity.PostResponse{
	post := dbEntity.Post{}
	result := service.postRepository.DeletePost(id, &post)

	output := &httpEntity.PostResponse{}
	if result == nil {
		copier.Copy(output, post)
	}
	return output
}