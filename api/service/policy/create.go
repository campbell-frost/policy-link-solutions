package policy

import (
	"github.com/campbell-frost/policy-link-solutions/database"
	"github.com/campbell-frost/policy-link-solutions/model"
)

func create(req *model.Policy) (model.Nothing, error) {
	db, err := database.Connect()
	if err != nil {
		return model.Nothing{}, err
	}

	result := db.Create(req)
	if result.Error != nil {
		return model.Nothing{}, result.Error
	}
	return model.Nothing{}, nil
}
