package eorm

import (
	"./logger"
	"./options"
	"./util"
	"database/sql"
	"errors"
)

type EormEngine struct {
	db     *sql.DB
	option *options.Options // 数据库配置
	path   string           // 请求路径
}

var eormEngine *EormEngine

var sqlService SqlService

func InitEorm(option *options.Options) {
	// 将配置项中的东西转化为路径
	path := util.GetPath(*option)

	// 获取到连接
	db, err := sql.Open(option.DriverName, path)
	if err != nil {
		logger.Error(err.Error())
	}

	// 初始化erom对象
	eormEngine = &EormEngine{
		db:     db,
		option: option,
		path:   path,
	}

	sqlService = newSqlService(db)
}

// 单个插入
func Insert(entity interface{}) (int64, error) {
	// 错误返回
	if eormEngine == nil {
		return -1, errors.New("未发现eorm")
	}

	// 通过反射解析出来的结果
	params, tableName := parseParam(entity, eormEngine.option)

	var param [][]eormData
	param = append(param, params)

	return sqlService.insert(tableName, param)
}

// 批量插入
func BatchInsert(entitys []interface{}) (int64, error) {
	// 错误返回
	if eormEngine == nil {
		return -1, errors.New("未发现eorm")
	}

	// 通过反射解析出来的结果
	var params [][]eormData
	var tableName = ""
	for _, entity := range entitys {
		param, name := parseParam(entity, eormEngine.option)
		tableName = name
		params = append(params, param)
	}

	return sqlService.insert(tableName, params)
}
