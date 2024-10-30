package convert

import (
	"github.com/pkg/errors"
	"reflect"
)

const (
	base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func IntToBase62(n int) string {
	if n == 0 {
		return string(base62[0])
	}

	var result []byte
	for n > 0 {
		result = append(result, base62[n%62])
		n /= 62
	}

	// 反转字符串
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Convert 是一个通用函数，用于将结构体 A 转换为结构体 B
func Convert(src interface{}, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// 确保 src 和 dst 都是指针
	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
		return errors.New("source and destination must be pointers")
	}

	// 确保 src 和 dst 指向的值都是结构体
	srcVal = srcVal.Elem()
	dstVal = dstVal.Elem()

	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		return errors.New("source and destination must be structs")
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		fieldName := srcVal.Type().Field(i).Name
		dstField := dstVal.FieldByName(fieldName)

		// 确保目标字段有效且可以设置
		if dstField.IsValid() && dstField.CanSet() {
			if srcField.Type() == dstField.Type() {
				dstField.Set(srcField)
			} else {
				return errors.New("source and destination field types do not match")
			}
		}
	}

	return nil
}
