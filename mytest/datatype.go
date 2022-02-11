package mytest

import (
	"commpkg/exdatatype"
	"fmt"
)

func Test_map() {
	data := make([]interface{}, 2)
	//var data [2]interface{}
	for i := 0; i < 2; i++ {

		datamap := make(map[string]interface{})
		datamap["name"] = "ddd"
		datamap["age"] = 20 + i
		data[i] = datamap
	}

	tmparr, _ := exdatatype.ToArray(data)
	//exdatatype.ToMap(data)
	for _, datadic := range tmparr {
		datamap, _ := exdatatype.ToMap(datadic)
		fmt.Println("name:", datamap["name"], "age:", datamap["age"])
	}
}
