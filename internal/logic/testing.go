package logic

import (
	"gitee.com/haiyux/kratos-layout/api/testing"
	"gitee.com/haiyux/kratos-layout/internal/dao"
)

func Gettests() (tests []*testing.Test) {
	ms := dao.GetTest()
	for _, v := range ms {
		tests = append(tests, &testing.Test{
			Id:   v.Id,
			Name: v.Name,
			Age:  int64(v.Age),
		})
	}
	return tests
}
