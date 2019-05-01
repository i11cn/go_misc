package misc

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
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

func (s StringConverter) ToInt() (int, error) {
	if i, err := strconv.ParseInt(string(s), 10, 32); err != nil {
		return 0, err
	} else {
		return int(i), nil
	}
}

func (s StringConverter) ToInt8() (int8, error) {
	if i, err := strconv.ParseInt(string(s), 10, 8); err != nil {
		return 0, err
	} else {
		return int8(i), nil
	}
}

func (s StringConverter) ToInt16() (int16, error) {
	if i, err := strconv.ParseInt(string(s), 10, 16); err != nil {
		return 0, err
	} else {
		return int16(i), nil
	}
}

func (s StringConverter) ToInt32() (int32, error) {
	if i, err := strconv.ParseInt(string(s), 10, 32); err != nil {
		return 0, err
	} else {
		return int32(i), nil
	}
}

func (s StringConverter) ToInt64() (int64, error) {
	if i, err := strconv.ParseInt(string(s), 10, 64); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

func (s StringConverter) ToUint() (uint, error) {
	if i, err := strconv.ParseUint(string(s), 10, 32); err != nil {
		return 0, err
	} else {
		return uint(i), nil
	}
}

func (s StringConverter) ToUint8() (uint8, error) {
	if i, err := strconv.ParseUint(string(s), 10, 8); err != nil {
		return 0, err
	} else {
		return uint8(i), nil
	}
}

func (s StringConverter) ToUint16() (uint16, error) {
	if i, err := strconv.ParseUint(string(s), 10, 16); err != nil {
		return 0, err
	} else {
		return uint16(i), nil
	}
}

func (s StringConverter) ToUint32() (uint32, error) {
	if i, err := strconv.ParseUint(string(s), 10, 32); err != nil {
		return 0, err
	} else {
		return uint32(i), nil
	}
}

func (s StringConverter) ToUint64() (uint64, error) {
	if i, err := strconv.ParseUint(string(s), 10, 64); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

func (s StringConverter) ToFloat32() (float32, error) {
	if i, err := strconv.ParseFloat(string(s), 32); err != nil {
		return 0, err
	} else {
		return float32(i), nil
	}
}

func (s StringConverter) ToFloat64() (float64, error) {
	if i, err := strconv.ParseFloat(string(s), 64); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

func (s StringConverter) ToBool() (bool, error) {
	switch strings.ToUpper(string(s)) {
	case "TRUE", "YES", "Y", "T", "1":
		return true, nil
	case "FALSE", "NO", "N", "F", "0":
		return false, nil
	}
	return false, fmt.Errorf("convert to bool failed")
}

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
