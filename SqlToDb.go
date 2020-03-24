package eorm

import (
	"./logger"
	"./util"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type SqlService interface {
	insert(tableName string, batchParams [][]eormData) (int64, error)
}

type SqlServiceImpl struct {
	db *sql.DB
}

func newSqlService(db *sql.DB) *SqlServiceImpl {
	return &SqlServiceImpl{
		db: db,
	}
}

func (sqlService *SqlServiceImpl) insert(tableName string, batchParams [][]eormData) (int64, error) {
	// sql 前缀
	prefix := "insert into " + tableName

	// sql 列名
	colums := make([]string, 0, len(batchParams[0]))

	// 解析参数
	for _, param := range batchParams[0] {
		colums = append(colums, param.colum)
		colums = append(colums, ",")
	}

	// 实际参数
	var insertParams []interface{}
	// 占位参数
	var preParams []string

	for _, params := range batchParams {
		paramItem := make([]rune, 0, len(params))
		paramItem = append(paramItem, '(')
		for _, param := range params {
			paramItem = append(paramItem, '?')
			paramItem = append(paramItem, ',')

			insertParams = append(insertParams, param.value)
		}
		paramItem = paramItem[:len(paramItem)-1]
		paramItem = append(paramItem, ')')

		s := string(paramItem)
		preParams = append(preParams, s)
	}

	// 得到实际列名
	colums = colums[:len(colums)-1]
	realColume := strings.Join(colums, "")
	// 实际占位符
	realPreParamStr := strings.Join(preParams, ",")

	// 获取sql
	realSql := util.GetSql(prefix, " (", realColume, ") values ", realPreParamStr)

	// 执行sql

	stmt, err2 := sqlService.db.Prepare(realSql)
	if err2 != nil {
		logger.Error(err2.Error())
	}

	// 执行sql
	if stmt != nil {
		result, err2 := stmt.Exec(insertParams...)
		if err2 != nil {
			logger.Error(err2.Error())
			return -1, errors.New("执行sql失败")
		}
		if result != nil {
			return result.RowsAffected()
		}
	}
	return -1, errors.New("执行sql失败")

}
