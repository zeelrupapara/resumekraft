package models

import "github.com/doug-martin/goqu/v9"

type Job struct {	
	Id int `json:"id"`
	JobDescription string `json:"job_description"`
}

type JobModel struct {
	db *goqu.Database
}

func InitJobModel(db *goqu.Database) (*JobModel, error) {
	return &JobModel{
		db: db,
	}, nil
}