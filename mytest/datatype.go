package mytest

import (
	"fmt"

	"github.com/wangxpfly/commpkg/exdatatype"
)

func Test_map() {
	//data := make([]interface{}, 2)  固定数组
	// 动态数组
	var data []interface{}
	//append(data,)
	//var data [2]interface{}
	for i := 0; i < 3; i++ {

		datamap := make(map[string]interface{})
		datamap["name"] = "ddd"
		datamap["age"] = 20 + i
		//data[i] = datamap
		data = append(data, datamap)
	}

	tmparr, _ := exdatatype.ToArray(data)
	//exdatatype.ToMap(data)
	for _, datadic := range tmparr {
		datamap, _ := exdatatype.ToMap(datadic)
		fmt.Println("name:", datamap["name"], "age:", datamap["age"])
	}
}
