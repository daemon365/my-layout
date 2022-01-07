package config

import (
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
)

func InitConfig(conf string, vs ...interface{}) error {
	logger := log.NewStdLogger(os.Stdout)
	c := config.New(
		config.WithLogger(logger),
		config.WithSource(
			file.NewSource(conf),
		),
	)
	if err := c.Load(); err != nil {
		return err
	}
	for _, v := range vs {
		err := c.Scan(v)
		if err != nil {
			return err
		}
	}
	return nil
}
