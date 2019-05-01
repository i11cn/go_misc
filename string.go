package misc

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	// StringConverter 把string类型的各种类型转换都组织起来
	StringConverter string
)

func (s StringConverter) to_int() (reflect.Value, error) {
	if i, err := strconv.ParseInt(string(s), 10, 64); err != nil {
		return reflect.ValueOf(i), err
	} else {
		return reflect.ValueOf(i), nil
	}
}

func (s StringConverter) to_uint() (reflect.Value, error) {
	if i, err := strconv.ParseUint(string(s), 10, 64); err != nil {
		return reflect.ValueOf(i), err
	} else {
		return reflect.ValueOf(i), nil
	}
}

func (s StringConverter) to_float() (reflect.Value, error) {
	if f, err := strconv.ParseFloat(string(s), 64); err != nil {
		return reflect.ValueOf(f), err
	} else {
		return reflect.ValueOf(f), nil
	}
}

// ToInt 字符串转换成 int
func (s StringConverter) ToInt() (int, error) {
	if i, err := strconv.ParseInt(string(s), 10, 32); err != nil {
		return 0, err
	} else {
		return int(i), nil
	}
}

// ToInt 字符串转换成 int8
func (s StringConverter) ToInt8() (int8, error) {
	if i, err := strconv.ParseInt(string(s), 10, 8); err != nil {
		return 0, err
	} else {
		return int8(i), nil
	}
}

// ToInt 字符串转换成 int16
func (s StringConverter) ToInt16() (int16, error) {
	if i, err := strconv.ParseInt(string(s), 10, 16); err != nil {
		return 0, err
	} else {
		return int16(i), nil
	}
}

// ToInt 字符串转换成 int32
func (s StringConverter) ToInt32() (int32, error) {
	if i, err := strconv.ParseInt(string(s), 10, 32); err != nil {
		return 0, err
	} else {
		return int32(i), nil
	}
}

// ToInt 字符串转换成 int64
func (s StringConverter) ToInt64() (int64, error) {
	if i, err := strconv.ParseInt(string(s), 10, 64); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

// ToInt 字符串转换成 uint
func (s StringConverter) ToUint() (uint, error) {
	if i, err := strconv.ParseUint(string(s), 10, 32); err != nil {
		return 0, err
	} else {
		return uint(i), nil
	}
}

// ToInt 字符串转换成 uint8
func (s StringConverter) ToUint8() (uint8, error) {
	if i, err := strconv.ParseUint(string(s), 10, 8); err != nil {
		return 0, err
	} else {
		return uint8(i), nil
	}
}

// ToInt 字符串转换成 uint16
func (s StringConverter) ToUint16() (uint16, error) {
	if i, err := strconv.ParseUint(string(s), 10, 16); err != nil {
		return 0, err
	} else {
		return uint16(i), nil
	}
}

// ToInt 字符串转换成 uint32
func (s StringConverter) ToUint32() (uint32, error) {
	if i, err := strconv.ParseUint(string(s), 10, 32); err != nil {
		return 0, err
	} else {
		return uint32(i), nil
	}
}

// ToInt 字符串转换成 uint64
func (s StringConverter) ToUint64() (uint64, error) {
	if i, err := strconv.ParseUint(string(s), 10, 64); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

// ToInt 字符串转换成 float32
func (s StringConverter) ToFloat32() (float32, error) {
	if i, err := strconv.ParseFloat(string(s), 32); err != nil {
		return 0, err
	} else {
		return float32(i), nil
	}
}

// ToInt 字符串转换成 float64
func (s StringConverter) ToFloat64() (float64, error) {
	if i, err := strconv.ParseFloat(string(s), 64); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

// ToInt 字符串转换成 bool，支持将true、yes、enable、y、t、1转换成true，false、no、disable、n、f、0转换成false，不区分大小写，其他数据一概不支持
func (s StringConverter) ToBool() (bool, error) {
	switch strings.ToUpper(string(s)) {
	case "TRUE", "YES", "ENABLE", "Y", "T", "1":
		return true, nil
	case "FALSE", "NO", "DISABLE", "N", "F", "0":
		return false, nil
	}
	return false, fmt.Errorf("convert to bool failed")
}

// ToType 根据输入的类型，将string转换成对应类型的Value
func (s StringConverter) ToType(t reflect.Type) (reflect.Value, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	ret := reflect.New(t).Elem()
	var err error
	switch t.String() {
	case "string":
		ret = reflect.ValueOf(string(s))
	case "int", "int8", "int16", "int32", "int64":
		if use, err := s.to_int(); err != nil {
			return ret, err
		} else {
			if ret.OverflowInt(use.Int()) {
				return ret, fmt.Errorf("%s 超出 %s 类型的取值范围", string(s), t.String())
			}
			ret.SetInt(use.Int())
		}
	case "uint", "uint8", "uint16", "uint32", "uint64":
		if use, err := s.to_uint(); err != nil {
			return ret, err
		} else {
			if ret.OverflowUint(use.Uint()) {
				return ret, fmt.Errorf("%s 超出 %s 类型的取值范围", string(s), t.String())
			}
			ret.SetUint(use.Uint())
		}
	case "float32", "float64":
		if use, err := s.to_float(); err != nil {
			return ret, err
		} else {
			if ret.OverflowFloat(use.Float()) {
				return ret, fmt.Errorf("%s 超出 %s 类型的取值范围", string(s), t.String())
			}
			ret.SetFloat(use.Float())
		}
	case "bool":
		if use, err := s.ToBool(); err != nil {
			return ret, err
		} else {
			ret.SetBool(use)
		}
	default:
		return ret, fmt.Errorf("type %s not supported by string converterr", t.String())
	}
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// ToValue 根据输入参数v的类型，将string转换成对应的类型，并且赋值给v，该形式方便直接检查返回的错误值
func (s StringConverter) ToValue(v interface{}) error {
	if v == nil {
		return fmt.Errorf("输入参数为 nil，不能用来接收数据")
	}
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("不支持获取到类型 ", value.Type().String(), " , 必须为指针类型")
	}
	value = value.Elem()
	use, err := s.ToType(value.Type())
	if err != nil {
		return err
	}
	value.Set(use)
	return nil
}
