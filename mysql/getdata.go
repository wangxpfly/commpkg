package mysql

import (
	"database/sql"
	"fmt"
	"reflect"
)

// 返回一条记录 map[string]interface{}
func GetRecord(rows *sql.Rows) (map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{}, len(columns))
	for rows.Next() {
		rowdatas := make([]interface{}, len(columns))
		for i := 0; i < len(rowdatas); i++ {
			var sss interface{}
			rowdatas[i] = &sss //makeEmptyStr()
		}
		rows.Scan(rowdatas...)
		for i := 0; i < len(rowdatas); i++ {
			k := columns[i]
			val := rowdatas[i].(*interface{})
			v := *val
			fmt.Println("data type:", reflect.TypeOf(v))
			sk := v.([]uint8)
			sk1 := string([]byte(sk))
			fmt.Println("data:", sk1)

			result[k] = sk1
			fmt.Println(result[k])
		}
	}
	return result, nil
}

// 返回多条记录 []interface{}
func GetRecords(rows *sql.Rows) ([]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var items []interface{}
	for rows.Next() {
		result := make(map[string]interface{}, len(columns))
		rowdatas := make([]interface{}, len(columns))
		for i := 0; i < len(rowdatas); i++ {
			var sss interface{}
			rowdatas[i] = &sss //makeEmptyStr()
		}
		rows.Scan(rowdatas...)
		for i := 0; i < len(rowdatas); i++ {
			k := columns[i]
			val := rowdatas[i].(*interface{})
			v := *val
			fmt.Println("data type:", reflect.TypeOf(v))
			sk := v.([]uint8)
			sk1 := string([]byte(sk))
			fmt.Println("data:", sk1)

			result[k] = sk1
			fmt.Println(result[k])
		}
		items = append(items, result)
	}
	return items, nil
}
