package eorm

import (
	"./options"
	"reflect"
	"strings"
)

type eormData struct {
	property string
	colum    string
	value    interface{}
}

// 解析参数
func parseParam(param interface{}, option *options.Options) (result []eormData, tableName string) {
	// 反射获取类型
	_type := reflect.TypeOf(param)

	// 获取表名并转化成下划线caml to underline
	tableName = camlToUnderline(_type.Name())

	// 反射获取值
	_value := reflect.ValueOf(param)

	// 遍历获取 key, value, tag
	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i)

		val := _value.Field(i).Interface()

		// 列命 默认驼峰法获取
		colum := camlToUnderline(field.Name)

		// 如果是指定 tag 列命
		if option.ColumAssign == options.CUSTOM_COLUM {
			colum = string(field.Tag.Get("colum"))
		}

		var eormData = eormData{
			property: field.Name,
			colum:    colum,
			value:    val,
		}
		result = append(result, eormData)
	}

	return
}

// 将驼峰法命名的字段转化为数据库字段 ex:NameMine -> name_mine
func camlToUnderline(property string) string {
	data := make([]byte, 0, len(property)*2)
	j := false
	num := len(property)
	for i := 0; i < num; i++ {
		d := property[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		// 防止第一个下划线
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
