package services

import (
	"fmt"
	"testing"
)
import (
	modelDB "example_app/entity/db"
)
import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// SETUP
type repositoryPostMock struct {
	mock.Mock
}

func (repository *repositoryPostMock) GetPostById(id int, post *modelDB.Post) error {
	repository.Called(id, post)
	post.Id = uint(id)
	post.Title = "Mary story"
	post.Body = "Mary has a litle lamb"
	return nil
}

func (repository *repositoryPostMock) DeletePost(id int, post *modelDB.Post) error {
	repository.Called(id, post)
	post.Id = uint(id)
	post.Title = fmt.Sprintf("Removed - %s", post.Title)
	post.Body = fmt.Sprintf("Removed - %s", post.Body)
	return nil
}

func (repository *repositoryPostMock) StorePost(post *modelDB.Post) error {
	repository.Called(post)
	post.Title = fmt.Sprintf("Stored - %s", post.Title)
	post.Body = fmt.Sprintf("Stored - %s", post.Body)
	return nil
}

func (repository *repositoryPostMock) UpdatePostById(id int, post *modelDB.Post) error {
	repository.Called(id, post)
	post.Id = uint(id)
	post.Title = fmt.Sprintf("Updated - %s", post.Title)
	post.Body = fmt.Sprintf("Updated - %s", post.Body)
	return nil
}

func (repository *repositoryPostMock) GetPostsList(limit int, offset int) ([]modelDB.Post, error) {
	repository.Called(limit, offset)
	posts := []modelDB.Post{}
	const (
	    ID1 = iota + 1
	    ID2
	    ID3
	)
	posts = append(posts, modelDB.Post{
		Id: uint(ID1),
		Title: "First Story",
		Body: "Mary has a litle lamb",
	})
	posts = append(posts, modelDB.Post{
		Id: uint(ID2),
		Title: "Second Story",
		Body: "Nocturne",
	})
	posts = append(posts, modelDB.Post{
		Id: uint(ID3),
		Title: "Third Story",
		Body: "Fur Elise",
	})
	return posts, nil
}


// TEST
func TestPostServiceGetPostByIdMocked(t *testing.T) {
	t.Parallel()
	dbMockData := repositoryPostMock{}
	
	post := &modelDB.Post{}
	var testId int = 1
	dbMockData.On("GetPostById", testId, post).Return(nil)
	
	postService := PostService{&dbMockData}
	resultFuncService := postService.GetPostById(testId)
	assert.Equal(t, uint(testId), resultFuncService.Id, "It should be same ID")
	assert.Equal(t, "Mary story", resultFuncService.Title, "It should be same Title")
	assert.Equal(t, "Mary has a litle lamb", resultFuncService.Body, "It should be same Body")
}

func TestPostServiceGetAllPostMocked(t *testing.T) {
	t.Parallel()
	dbMockData := repositoryPostMock{}
	limit := 1
	offset := 3
	dbMockData.On("GetPostsList", limit, offset).Return([]modelDB.Post{}, nil)
	postService := PostService{&dbMockData}
	resultFuncService := postService.GetAllPosts(limit, offset)
	assert.Equal(t, len(resultFuncService), 3, "It should be same length as Mock Data")
	assert.Equal(t, resultFuncService[0].Title, "First Story", "It should be same NAME as Mock Data")
	assert.Equal(t, resultFuncService[1].Title, "Second Story", "It should be same NAME as Mock Data")
	assert.Equal(t, resultFuncService[2].Title, "Third Story", "It should be same NAME as Mock Data")
}
