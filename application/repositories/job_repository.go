package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func (repository JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {

	err := repository.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (repository JobRepositoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job
	repository.Db.Preload("Video").First(&job, "id = ?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	return &job, nil
}

func (repository JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := repository.Db.Save(&job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}
