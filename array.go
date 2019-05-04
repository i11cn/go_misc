package misc

import (
	"reflect"
)

var (
	ReverseString func([]string)
)

func init() {
	MakeReverse(&ReverseString)
}

func reverse(in []reflect.Value) []reflect.Value {
	arr := in[0]
	if arr.Kind() == reflect.Slice {
		total := arr.Len()
		l := total - 1
		for i := 0; i < total/2; i++ {
			tmp := reflect.New(arr.Index(i).Type())
			tmp.Elem().Set(arr.Index(i))
			arr.Index(i).Set(arr.Index(l - i))
			arr.Index(l - i).Set(tmp.Elem())
		}
	}
	return []reflect.Value{}
}

// MakeFunc 比较通用的创建func的函数，fptr是指向函数类型的指针，根据fptr的类型，把fn转换成fptr，并且赋值给fptr
func MakeFunc(fptr interface{}, fn func(args []reflect.Value) (results []reflect.Value)) {
	f := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(f.Type(), fn)
	f.Set(v)
}

// MakeReverse 算是MakeFunc的一个示例，演示如何把通用的reverse函数改变成fptr类型，并赋值给fptr
func MakeReverse(fptr interface{}) {
	MakeFunc(fptr, reverse)
}

// DropEmpty 丢弃字符串数组中的空字符串
func DropEmpty(in []string) []string {
	do := false
	for _, s := range in {
		if s == "" {
			do = true
			break
		}
	}
	if !do {
		return in
	}
	ret := make([]string, 0, len(in))
	for _, s := range in {
		if s != "" {
			ret = append(ret, s)
		}
	}
	return ret
}
