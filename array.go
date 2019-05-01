package misc

import (
	"io/ioutil"
	"os"
	"reflect"
)

var (
	ReverseString func([]string)
)

func init() {
	MakeReverse(&ReverseString)
}

// ReadFileAll 很多地方需要用到[]byte的输入(例如json.Unmarshal等)，那么从文件读取数据也是很常用的方式，于是把从文件读取到[]byte写成一个函数
//
// 返回值中的[]byte和error绝不可能同时有效，也不可能同时为nil
func ReadFileAll(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	ret, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return ret, nil
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
		if len(s) == 0 {
			do = true
			break
		}
	}
	if !do {
		return in
	}
	ret := make([]string, 0, len(in))
	for _, s := range in {
		if len(s) > 0 {
			ret = append(ret, s)
		}
	}
	return ret
}
