package exdatatype

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ToInt 转换为int类型
func ToInt(value interface{}) (int, error) {
	if value == nil {
		return 0, errors.New("params is nil")
	}
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.Atoi(v)
	case int:
		v, _ := value.(int)
		return v, nil
	case int32:
		v, _ := value.(int32)
		return int(v), nil
	case int64:
		v, _ := value.(int64)
		return int(v), nil
	case float32:
		v, _ := value.(float32)
		return int(v), nil
	case float64:
		v, _ := value.(float64)
		return int(v), nil
	default:
		return int(0), errors.New("unknown type")
	}
}

// ToInt32 转换为Int32
func ToInt32(value interface{}) (int32, error) {
	if value == nil {
		return 0, errors.New("params is nil")
	}
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseInt(v, 10, 32)
		return int32(result), err
	case int:
		v, _ := value.(int)
		return int32(v), nil
	case int32:
		v, _ := value.(int32)
		return int32(v), nil
	case int64:
		v, _ := value.(int64)
		return int32(v), nil
	case float32:
		v, _ := value.(float32)
		return int32(v), nil
	case float64:
		v, _ := value.(float64)
		return int32(v), nil
	default:
		return int32(0), errors.New("unknown type")
	}
}

// ToInt64 转换为Int64
func ToInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, errors.New("params is nil")
	}
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseInt(v, 10, 64)
	case int:
		v, _ := value.(int)
		return int64(v), nil
	case int32:
		v, _ := value.(int32)
		return int64(v), nil
	case int64:
		v, _ := value.(int64)
		return v, nil
	case float32:
		v, _ := value.(float32)
		return int64(v), nil
	case float64:
		v, _ := value.(float64)
		return int64(v), nil
	default:
		return int64(0), errors.New("unknown type")
	}
}

// ToFloat32 转换为ToFloat32
func ToFloat32(value interface{}) (float32, error) {
	if value == nil {
		return 0, errors.New("params is nil")
	}
	switch value.(type) {
	case string:
		v, _ := value.(string)
		result, err := strconv.ParseFloat(v, 32)
		return float32(result), err
	case int:
		v, _ := value.(int)
		return float32(v), nil
	case int32:
		v, _ := value.(int32)
		return float32(v), nil
	case int64:
		v, _ := value.(int64)
		return float32(v), nil
	case float32:
		v, _ := value.(float32)
		return v, nil
	case float64:
		v, _ := value.(float64)
		return float32(v), nil
	default:
		return float32(0), errors.New("unknown type")
	}
}

// ToFloat64 转换为 float64
func ToFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0, errors.New("params is nil")
	}
	switch value.(type) {
	case string:
		v, _ := value.(string)
		return strconv.ParseFloat(v, 64)
	case int:
		v, _ := value.(int)
		return float64(v), nil
	case int32:
		v, _ := value.(int32)
		return float64(v), nil
	case int64:
		v, _ := value.(int64)
		return float64(v), nil
	case float32:
		v, _ := value.(float32)
		return float64(v), nil
	case float64:
		v, _ := value.(float64)
		return v, nil
	default:
		return float64(0), errors.New("unknown type")
	}
}

// 转换成 Mongo 的ID对象
func ToObjectID(value interface{}) (primitive.ObjectID, error) {
	if value == nil {
		return primitive.NilObjectID, errors.New("params is nil")
	}
	switch value.(type) {
	case string:
		v, err := primitive.ObjectIDFromHex(value.(string))
		if err != nil {
			return primitive.NilObjectID, err
		}
		return v, nil
	case primitive.ObjectID:
		v, _ := value.(primitive.ObjectID)
		return v, nil
	default:
		return primitive.NilObjectID, errors.New("unknown type")
	}
}

// ToString args: value,
func ToString(str interface{}) (string, error) {
	if str == nil {
		return "", errors.New("params is nil")
	}
	value := str
	//fmt.Println("value:", value)
	precision := 12 //浮点型，默认最多保持小数12位
	switch value.(type) {
	case string:
		fmt.Println("string:", value)
		v, _ := value.(string)
		//fmt.Println("stringv:", v, err)
		return v, nil
	case int:
		//fmt.Println("int:", value)
		v, _ := value.(int)
		return strconv.Itoa(v), nil
	case int32:
		//fmt.Println("int32:", value)
		v, _ := value.(int32)
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		//fmt.Println("int64:", value)
		v, _ := value.(int64)
		return strconv.FormatInt(v, 10), nil
	case float32:
		//fmt.Println("float32:", value)
		v, _ := value.(float32)
		return strconv.FormatFloat(float64(v), 'f', precision, 32), nil
	case float64:
		v, _ := value.(float64)
		return strconv.FormatFloat(v, 'f', precision, 64), nil
	case primitive.ObjectID:
		v, _ := value.(primitive.ObjectID)
		return v.Hex(), nil
	default:
		//fmt.Println("default:", value)
		return "", errors.New("unknown type")
	}
}

// 转换成MAP map[string]interface{}
func ToMap(data interface{}) (map[string]interface{}, error) {
	if data == nil {
		return nil, errors.New("params is nil")
	}
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		tmpmap := data.(map[string]interface{})
		//fmt.Println("pp:", tmpmap)
		return tmpmap, nil
	default:
		return nil, errors.New("unknown type")
	}
}

// 转换成数组 []interface{}
func ToArray(data interface{}) ([]interface{}, error) {
	if data == nil {
		return nil, errors.New("params is nil")
	}
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice, reflect.Array:
		//fmt.Println("元素的值是: ", data, "反射类型是:", reflect.TypeOf(data))
		tmparr := data.([]interface{})
		return tmparr, nil
	default:
		return nil, errors.New("unknown type")
	}
}
