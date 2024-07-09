package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

type Table struct {
	// 搪行器
	executor *Executor

	// 表名
	name string

	tStruct any

	where []any

	offset int

	limit int

	orderBy string
}

// Init 初始化
func (table *Table) Init() *Table {
	table.where = []any{}
	table.offset = 0
	table.limit = 20
	table.orderBy = ""
	return table
}

// SetExecutor 设置执行品
func (table *Table) SetExecutor(e *Executor) *Table {
	table.executor = e
	return table
}

// SetName 设置名称
func (table *Table) SetName(name string) *Table {
	table.name = name
	return table
}

// SetStruct 设置结构
func (table *Table) SetStruct(s any) *Table {
	table.tStruct = s
	return table
}

// Where 查询条件
func (table *Table) Where(params ...any) *Table {
	if len(table.where) > 0 {
		table.where = append(table.where, "AND")
	}

	l := len(params)
	if l == 1 {
		switch t := params[0].(type) {
		case string:
			table.where = append(table.where, t)
		}
	} else if l == 2 {
		switch t := params[0].(type) {
		case string:
			table.where = append(table.where, [3]any{t, "=", params[1]})
		}
	} else if l == 3 {
		switch t1 := params[0].(type) {
		case string:
			switch t2 := params[1].(type) {
			case string:
				table.where = append(table.where, [3]any{t1, strings.ToUpper(strings.TrimSpace(t2)), params[2]})
			}
		}
	}

	return table
}

// WhereGroup 一组查询条件
func (table *Table) WhereGroup(params []any) *Table {
	if len(table.where) > 0 {
		table.where = append(table.where, "AND")
	}

	table.where = append(table.where, "(")

	for _, v := range params {
		switch t := v.(type) {
		case string:
			table.where = append(table.where, t)
		case [1]string:
			table.where = append(table.where, t[0])
		case [2]string:
			table.where = append(table.where, [3]any{t[0], "=", t[1]})
		case [3]string:
			table.where = append(table.where, [3]any{t[0], t[1], t[2]})
		case []string:
			l := len(t)
			if l == 1 {
				table.where = append(table.where, t[0])
			} else if l == 2 {
				table.where = append(table.where, [3]any{t[0], "=", t[1]})
			} else if l == 3 {
				table.where = append(table.where, [3]any{t[0], t[1], t[2]})
			}
		case [2]any:
			switch field := t[0].(type) {
			case string:
				table.where = append(table.where, [3]any{field, "=", t[1]})
			}
		case [3]any:
			switch field := t[0].(type) {
			case string:
				switch op := t[1].(type) {
				case string:
					table.where = append(table.where, [3]any{field, strings.ToUpper(strings.TrimSpace(op)), t[2]})
				}
			}
		case []any:
			l := len(t)
			if l == 1 {
				switch str := t[0].(type) {
				case string:
					table.where = append(table.where, str)
				}
			} else if l == 2 {
				switch field := t[0].(type) {
				case string:
					table.where = append(table.where, [3]any{field, "=", t[1]})
				}
			} else if l == 3 {
				switch field := t[0].(type) {
				case string:
					switch op := t[1].(type) {
					case string:
						table.where = append(table.where, [3]any{field, strings.ToUpper(strings.TrimSpace(op)), t[2]})
					}
				}
			}
		}
	}

	table.where = append(table.where, ")")
	return table
}

// Offset 编移量
func (table *Table) Offset(offset int) *Table {
	table.offset = offset
	return table
}

// Limit 读取条数，即分页大小
func (table *Table) Limit(limit int) *Table {
	table.limit = limit
	return table
}

// OrderBy 排序
func (table *Table) OrderBy(field string, dir string) *Table {
	dir = strings.ToUpper(dir)
	if dir != "ASC" && dir != "DESC" {
		dir = "ASC"
	}

	table.orderBy = field + dir
	return table
}

// OrderByStr 排序字符串
func (table *Table) OrderByStr(orderBy string) *Table {
	table.orderBy = orderBy
	return table
}

// GetValue 获取值
func (table *Table) GetValue(fields string) string {
	//return table.query("GetValue", fields)
	return ""
}

// Count 获取总数
func (table *Table) Count(fields string) int {
	//return table.query("GetValue", fields)
	return 0
}

func (table *Table) query(fn string, fields string) {

}

func (table *Table) prepareSql() (string, []any) {
	query := ""
	var params []any

	if len(table.where) > 0 {
		query += " WHERE"
		for _, v := range table.where {
			switch t := v.(type) {
			case string:
				query += " " + t
			case [3]any:
				switch field := t[0].(type) {
				case string:
					switch op := t[1].(type) {
					case string:

						switch op {
						case "IN", "NOT IN":

						case "BETWEEN", "NOT BETWEEN":

						default:

							query += " " + field + " " + op + " ?"
							params = append(params, t[2])
						}
					}
				}

			}
		}
	}

	if table.limit > 0 {
		if table.offset > 0 {
			query += " LIMIT " + strconv.Itoa(table.offset) + "," + strconv.Itoa(table.limit)
		} else {
			query += " LIMIT " + strconv.Itoa(table.limit)
		}
	} else {
		if table.offset > 0 {
			query += " OFFSET " + strconv.Itoa(table.offset)
		}
	}

	return query, params
}
