package dao

import "gitee.com/haiyux/kratos-layout/internal/models"

// mock, should be operation in database
func GetTest() []*models.TestModels {
	return []*models.TestModels{
		{
			Id:   1,
			Name: "test1",
			Age:  1,
		},
		{
			Id:   2,
			Name: "test2",
			Age:  2,
		},
	}
}
