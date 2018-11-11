package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

import (
	"gcrud/run"
	"gcrud/model"
)

/* ================================================================================
 * 生成数据
 * author: jcheng
 * ================================================================================ */

func main() {
	run.DatabaseConfig = &model.DatabaseConfig{
		DatabaseName: "test",
		Host:         "127.0.0.1",
		Port:         "3306",
		UserName:     "test",
		PassWord:     "test",
	}
	run.Run()
}
