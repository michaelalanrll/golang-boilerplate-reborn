package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	connection "example_app/util/helper/mysqlconnection"
	dbEntity "example_app/entity/db"
)

type PostRepository struct {
	DB gorm.DB
}

func PostRepositoryHandler() *PostRepository {
	return &PostRepository{DB: *connection.GetConnection()}
}

type PostRepositoryInterface interface {
	GetPostById(id int, postData *dbEntity.Post) error
	UpdatePostById(id int, postData *dbEntity.Post) error
	GetPostsList(limit int, offset int) ([]dbEntity.Post, error)
	StorePost(postData *dbEntity.Post) error
	DeletePost(id int, postData *dbEntity.Post) error
}

func (repository *PostRepository) GetPostById(id int, postData *dbEntity.Post) error {
	query := repository.DB.Preload("User")
	query = query.Where("id=?", id)
	query = query.First(postData)
	return query.Error
}

func (repository *PostRepository) UpdatePostById(id int, postData *dbEntity.Post) error {
	query := repository.DB.Table("posts")
	query = query.Where("id=?", id)
	success := query.Updates(postData).RowsAffected
	if success < 1 {
		return errors.New("No data affected")
	}
	return query.Error
}

func (repository *PostRepository) GetPostsList(limit int, offset int) ([]dbEntity.Post, error) {
	posts := []dbEntity.Post{}
	query := repository.DB.Table("posts")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&posts)
	return posts, query.Error
}

func (repository *PostRepository) StorePost(postData *dbEntity.Post) error {
	query := repository.DB.Table("posts")
	query = query.Create(postData)
	return query.Error
}

func (repository *PostRepository) DeletePost(id int, postData *dbEntity.Post) error {
	post := &dbEntity.Post{}
	query := repository.DB.Table("posts")
	query = query.Where("id=?", id)
	query = query.First(postData)
	query = query.Delete(post)
	return query.Error
}