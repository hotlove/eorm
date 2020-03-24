package util

import (
	"../options"
	"strconv"
	"strings"
)

func GetPath(options options.Options) string {
	portStr := strconv.Itoa(options.Port)

	pathSlice := []string{
		options.UserName,
		":",
		options.Password,
		"@tcp(", options.Host, ":", portStr, ")/",
		options.Database,
		"?charset=utf8mb4"}
	path := strings.Join(pathSlice, "")
	return path
}

func GetSql(vals ...string) string {
	sql := make([]string, 0, len(vals))
	for _, val := range vals {
		sql = append(sql, val)
	}

	return strings.Join(sql, "")
}
