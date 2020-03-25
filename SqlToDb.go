package eorm

import (
	"database/sql"
	"errors"
	"strings"

	"./logger"
	"./util"
	_ "github.com/go-sql-driver/mysql"
)

type SqlService interface {
	// 根据参数批量删除
	insert(tableName string, batchParams [][]eormData) (int64, error)

	// 根据条件删除记录
	delete(tableName string, andParam interface{}, orParam interface{})
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

// 根据条件删除记录
func (sqlService *SqlServiceImpl) delete(tableName string, andParam []eormData, orParam []eormData) {
	// sql前缀
	sqlPrefix := "delete from" + tableName;

	// 完整sql
	var sql string

	andParamNum := len(andParam)
	orParamNum := len(orParam)
	if andParamNum == 0 && orParamNum == 0 {
		// 如果两个参数都没有
		sql = sqlPrefix
	}

	// sql 执行参数
	params := make([]interface{}, 0, andParamNum + orParamNum)

	// 如果与请求数据不为空
	if andParamNum > 0 {
		sqlParmas, sqlPlaceHolder := getParamAndSql(andParam)
		if orParamNum > 0 {
			sqlPrefix = sqlPrefix + " where (" + sqlPlaceHolder + ")"
		} else {
			sql = sqlPrefix + " where" + sqlPlaceHolder
		}
		
	}

	// 如果或请求数据不为空
	if orParamNum> 0 {
		andParamStr := make([]string, 0, andParamNum)
		for _, item := range {
			t := item.property + " = ?"
			andParamStr = append(andParamStr, t)
			params = append(params, item.value)
		}
		andStr := strings.join(andParamStr, ",")
		if orParamNum > 0 {
			sqlPrefix = sqlPrefix + " where (" + andStr + ")"
		} else {
			sql = sqlPrefix + " where" + andStr
		}
	}
}

func getParamAndSql(params []interface{}) ([]interface{}, string) {
	sqlParam := make([]interface{}, 0, len(params))
	sqlPlaceHolder := make([]string, 0, len(params))
	for _, item := range params{
		t := item.property + " = ?"
		sqlPlaceHolder = append(sqlPlaceHolder, t)
		sqlParam = append(sqlParam, item.value)
	}
	andStr := strings.join(sqlPlaceHolder, ",")

	return sqlParam, andStr
}