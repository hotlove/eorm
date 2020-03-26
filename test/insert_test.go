package test

import (
	"fmt"
	"github.com/hotlove/eorm"
	"github.com/hotlove/eorm/options"
	"testing"
)

type MyProfile struct {
	Id   int
	Name string
	Sex  string
}

func TestInsert(t *testing.T) {
	newOptions :=
		options.NewOptions(options.Host("192.168.20.51"),
			options.Password("123456"),
			options.UserName("root"),
			options.Port(3306),
			options.Database("spring-boot"),
			options.DriverName(options.MYSQL))

	eorm.InitEorm(newOptions)

	var batchInsert []interface{}
	for i := 20; i < 30; i++ {
		data := MyProfile{
			Id:   i,
			Name: "xiaoqiang",
			Sex:  "2",
		}
		batchInsert = append(batchInsert, data)
	}
	aa, err := eorm.BatchInsert(batchInsert)
	fmt.Printf("%+v", aa)
	fmt.Printf("%+v", err)
}
