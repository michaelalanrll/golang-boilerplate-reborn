package controller

import (
	"fmt"
	"time"
	"net/http"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	httpEntity "example_app/entity/http"
)

type PostServiceMock struct{}

func (service *PostServiceMock) GetPostById(id int) *httpEntity.PostDetailResponse {
	t, _ := time.Parse("2006-01-02", "2019-08-10")
	return &httpEntity.PostDetailResponse{
		Id: uint(id),
		Title: "Mary Story",
		Body: "Mary has a little lamb",
		UserId: 35,
		CreatedAt:&t,
		UpdatedAt:&t,
	}
}

func (service *PostServiceMock) GetAllPosts(page int,count int) []httpEntity.PostResponse{
	Posts := []httpEntity.PostResponse{}
	t, _ := time.Parse("2006-01-02", "2019-08-10")
	const (
	    ID1 = iota + 1
	    ID2
	    ID3
	)
	Posts = append(Posts, httpEntity.PostResponse{
		Id: uint(ID1),
		Title: "First Story",
		CreatedAt:&t,
		UpdatedAt:&t,
	})
	Posts = append(Posts, httpEntity.PostResponse{
		Id: uint(ID2),
		Title: "Second Story",
		CreatedAt:&t,
		UpdatedAt:&t,
	})
	Posts = append(Posts, httpEntity.PostResponse{
		Id: uint(ID1),
		Title: "Third Story",
		CreatedAt:&t,
		UpdatedAt:&t,
	})
	return Posts
}

func (service *PostServiceMock) UpdatePostById(id int, payload httpEntity.PostRequest) bool{
	return true
}

func (service *PostServiceMock) DeletePost(id int) *httpEntity.PostResponse{
	t, _ := time.Parse("2006-01-02", "2019-08-10")
	return &httpEntity.PostResponse{
		Id: uint(id),
		Title: "Mary Story",
		CreatedAt:&t,
		UpdatedAt:&t,
	}
}

func (service *PostServiceMock) StorePost(payload httpEntity.PostRequest) bool {
	return true
}

func TestPostGetByIDMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	var idTest uint = 1
	url := "/v1/posts/" + fmt.Sprint(idTest)
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusOK, resp.Code, "Status should be 200")

	res := httpEntity.PostDetailResponse{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)
	
	assert.Equal(err, nil, "should have no error")
	assert.Equal(res.Id, idTest, "It should be same ID")
	assert.Equal(res.Title, "Mary Story", "It should be same Title")
	assert.Equal(res.Body, "Mary has a little lamb", "It should be same Body")
}

func TestGetPostListMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	url := "/v1/posts"
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusOK, resp.Code, "Status should be 200")

	res := []httpEntity.PostResponse{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)
	
	assert.Equal(err, nil, "should have no error")
	assert.Equal(len(res)>=0, true, "length must in minimum value")
	assert.Equal(len(res)==3, true, "length value must match with mock data")
}
