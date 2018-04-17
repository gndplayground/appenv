package appenv

import (
	"errors"
	"os"
	"reflect"
	"testing"
	"fmt"
)

func loopConvert(mapKey map[reflect.Kind]string, finalType string, t *testing.T) {

	for k, v := range mapKey {
		result, e := stringConv(v, k)
		if e != nil {
			t.Error("got error when convert type " + k.String())
			fmt.Println(e)
		} else if result == nil {
			t.Errorf("value return should not be nil (Test case %v)", k)
		} else {
			if reflect.ValueOf(result).Type().String() != finalType {
				t.Errorf("failled convert type %v. Result %v", k, reflect.ValueOf(result).Type().String())
			}
		}
	}
}

func TestConTypeFunctionShouldSuccess(t *testing.T) {

	var uintKey = map[reflect.Kind]string{
		reflect.Uint:   "7000",
		reflect.Uint8:  "80",
		reflect.Uint16: "808",
		reflect.Uint32: "80888",
		reflect.Uint64: "804581205",
	}

	var intKey = map[reflect.Kind]string{
		reflect.Int:   "7000",
		reflect.Int8:  "80",
		reflect.Int16: "808",
		reflect.Int32: "80888",
		reflect.Int64: "804581205",
	}

	loopConvert(uintKey, "uint64", t)
	loopConvert(intKey, "int64", t)

	var v interface{}
	var e error

	v, e = stringConv("true", reflect.Bool)

	if e != nil {
		t.Error("error when convert string true to bool")
		fmt.Println(e)
	} else if v == nil {
		t.Errorf("value return should not be nil (Test case %v)", "true")
	} else if reflect.ValueOf(v).Type().String() != "bool" && v != true {
		t.Errorf("failled convert type bool with value %v. Result %v", "true", reflect.ValueOf(v).Type().String())
	}

	v, e = stringConv("false", reflect.Bool)

	if e != nil {
		t.Error("error when convert string false to bool")
		fmt.Println(e)
	} else if v == nil {
		t.Errorf("value return should not be nil (Test case %v)", "false")
	} else if reflect.ValueOf(v).Type().String() != "bool" && v != false {
		t.Errorf("failled convert type bool with value %v. Result %v", "false", reflect.ValueOf(v).Type().String())
	}

	v, e = stringConv("6.5", reflect.Float32)

	if e != nil {
		t.Error("error when convert string 6.5 to float32")
		fmt.Println(e)
	} else if v == nil {
		t.Errorf("value return should not be nil (Test case %v)", "6.5")
	} else if reflect.ValueOf(v).Type().String() != "float64" && v != false {
		t.Errorf("failled convert type float32 with value %v. Result %v", "6.5", reflect.ValueOf(v).Type().String())
		fmt.Println(reflect.ValueOf(v).Type().String())
	}

	v, e = stringConv("6.5", reflect.Float64)

	if e != nil {
		t.Error("error when convert string 6.5 to float64")
		fmt.Println(e)
	} else if v == nil {
		t.Errorf("value return should not be nil (Test case %v)", "6.5")
	} else if reflect.ValueOf(v).Type().String() != "float64" && v != false {
		t.Errorf("failled convert type float64 with value %v. Result %v", "6.5", reflect.ValueOf(v).Type().String())
		fmt.Println(reflect.ValueOf(v).Type().String())
	}
}

func TestConTypeFunctionShouldFail(t *testing.T) {
	var v interface{}
	var e error

	v, e = stringConv("true", reflect.Complex128)

	if e == nil && v != nil {
		t.Error("Should return error when try convert unknown type")
	}

	v, e = stringConv("abc", reflect.Bool)

	if e == nil && v != nil {
		t.Error("Should return error when convert string not true or false to bool")
	}
}


type testListUint struct {
	Uint   uint
	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64
}

type testListInt struct {
	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64
}

func TestInitShouldSuccess(t *testing.T) {

	type configTestSuccess struct {
		Uint         uint    `env:"TEST_ENV_Uint" envDefault:"10"`
		Uint8        uint8   `env:"TEST_ENV_Uint8" envDefault:"10"`
		Uint16       uint16  `env:"TEST_ENV_Uint16" envDefault:"10"`
		Uint32       uint32  `env:"TEST_ENV_Uint32" envDefault:"10"`
		Uint64       uint64  `env:"TEST_ENV_Uint64" envDefault:"10"`
		Int          int     `env:"TEST_ENV_Int" envDefault:"-10"`
		Int8         int8    `env:"TEST_ENV_Int8" envDefault:"-10"`
		Int16        int16   `env:"TEST_ENV_Int16" envDefault:"-10"`
		Int32        int32   `env:"TEST_ENV_Int32" envDefault:"-10"`
		Int64        int64   `env:"TEST_ENV_Int64" envDefault:"-10"`
		Float32      float32 `env:"TEST_ENV_FLOAT32" envDefault:"3.2"`
		Float64      float64 `env:"TEST_ENV_FLOAT32" envDefault:"6.5"`
		String       string  `env:"TEST_ENV_STRING" envDefault:"hello"`
		StringFromOS string  `env:"TEST_ENV_STRINGFromOS" envDefault:"hello"`
	}

	os.Setenv("TEST_ENV_STRINGFromOS", "bonjour")

	c := configTestSuccess{}

	Init(&c)

	if c.StringFromOS != "bonjour" || reflect.ValueOf(c.String).Kind().String() != "string" {
		t.Error("Failed set StringFromOS")
	}

	if c.String != "hello" || reflect.ValueOf(c.String).Kind().String() != "string" {
		t.Error("Failed set String")
	}

	if c.Float32 != 3.2 || reflect.ValueOf(c.Float32).Kind().String() != "float32" {
		t.Error("Failed set Float32")
	}

	if c.Float64 != 6.5 || reflect.ValueOf(c.Float64).Kind().String() != "float64" {
		t.Error("Failed set Float64")
	}

	if c.Uint != 10 || reflect.ValueOf(c.Uint).Kind().String() != "uint" {
		t.Error("Failed set Uint")
	}

	if c.Uint8 != 10 || reflect.ValueOf(c.Uint8).Kind().String() != "uint8" {
		t.Error("Failed set Uint8")
	}

	if c.Uint16 != 10 || reflect.ValueOf(c.Uint16).Kind().String() != "uint16" {
		t.Error("Failed set Uint16")
	}

	if c.Uint32 != 10 || reflect.ValueOf(c.Uint32).Kind().String() != "uint32" {
		t.Error("Failed set Uint32")
	}

	if c.Uint64 != 10 || reflect.ValueOf(c.Uint64).Kind().String() != "uint64" {
		t.Error("Failed set Uint64")
	}

	if c.Int != -10 || reflect.ValueOf(c.Int).Kind().String() != "int" {
		t.Error("Failed set Int")
	}

	if c.Int8 != -10 || reflect.ValueOf(c.Int8).Kind().String() != "int8" {
		t.Error("Failed set Int8")
	}

	if c.Int16 != -10 || reflect.ValueOf(c.Int16).Kind().String() != "int16" {
		t.Error("Failed set Int16")
	}

	if c.Int32 != -10 || reflect.ValueOf(c.Int32).Kind().String() != "int32" {
		t.Error("Failed set Int32")
	}

	if c.Int64 != -10 || reflect.ValueOf(c.Int64).Kind().String() != "int64" {
		t.Error("Failed set Int64")
	}

}

func TestInitShouldSuccessCaseArray(t *testing.T) {

	os.Setenv("TEST_ENV_ArrayStringFromOS", "a,c,d")

	type configTestSuccess struct {
		ArrayStringFromOS []string  `env:"TEST_ENV_ArrayStringFromOS" envDefault:"a,b,c"`
		ArrayString       []string  `env:"TEST_ENV_ArrayString" envDefault:"a,b,c"`
		ArrayBool         []bool    `env:"TEST_ENV_ArrayBool" envDefault:"true,true,false"`
		ArrayFloat64      []float64 `env:"TEST_ENV_ArrayFloat64" envDefault:"5.4,5.1,5.3"`
		ArrayFloat32      []float32 `env:"TEST_ENV_ArrayFloat32" envDefault:"5.4,5.1,5.3"`
		ArrayUint         []uint    `env:"TEST_ENV_ArrayUint" envDefault:"1,3,5"`
		ArrayUint8        []uint8   `env:"TEST_ENV_ArrayUint8" envDefault:"1,3,5"`
		ArrayUint16       []uint16  `env:"TEST_ENV_ArrayUint16" envDefault:"1,3,5"`
		ArrayUint32       []uint32  `env:"TEST_ENV_ArrayUint32" envDefault:"1,3,5"`
		ArrayUint64       []uint64  `env:"TEST_ENV_ArrayUint64" envDefault:"1,3,5"`
		ArrayInt          []int     `env:"TEST_ENV_ArrayInt" envDefault:"1,-3,5"`
		ArrayInt8         []int8    `env:"TEST_ENV_ArrayInt8" envDefault:"1,-3,5"`
		ArrayInt16        []int16   `env:"TEST_ENV_ArrayInt16" envDefault:"1,-3,5"`
		ArrayInt32        []int32   `env:"TEST_ENV_ArrayInt32" envDefault:"1,-3,5"`
		ArrayInt64        []int64   `env:"TEST_ENV_ArrayInt64" envDefault:"1,-3,5"`
	}

	c := configTestSuccess{}

	Init(&c)

	if !reflect.DeepEqual(c.ArrayStringFromOS, []string{"a", "c", "d"}) {
		t.Error("Failed set ArrayStringFromOS")
	}

	if !reflect.DeepEqual(c.ArrayString, []string{"a", "b", "c"}) {
		t.Error("Failed set ArrayString")
	}

	if !reflect.DeepEqual(c.ArrayBool, []bool{true, true, false}) {
		t.Error("Failed set ArrayBool")
	}

	if !reflect.DeepEqual(c.ArrayFloat64, []float64{5.4, 5.1, 5.3}) {
		t.Error("Failed set ArrayFloat64")
	}

	if !reflect.DeepEqual(c.ArrayFloat32, []float32{5.4, 5.1, 5.3}) {
		t.Error("Failed set ArrayFloat32")
	}

	if !reflect.DeepEqual(c.ArrayUint, []uint{1, 3, 5}) {
		t.Error("Failed set ArrayUint")
	}

	if !reflect.DeepEqual(c.ArrayUint8, []uint8{1, 3, 5}) {
		t.Error("Failed set ArrayUint8")
	}

	if !reflect.DeepEqual(c.ArrayUint16, []uint16{1, 3, 5}) {
		t.Error("Failed set ArrayUint16")
	}

	if !reflect.DeepEqual(c.ArrayUint32, []uint32{1, 3, 5}) {
		t.Error("Failed set ArrayUint32")
	}

	if !reflect.DeepEqual(c.ArrayUint64, []uint64{1, 3, 5}) {
		t.Error("Failed set ArrayUint64")
	}

	if !reflect.DeepEqual(c.ArrayInt, []int{1, -3, 5}) {
		t.Error("Failed set ArrayInt")
	}

	if !reflect.DeepEqual(c.ArrayInt8, []int8{1, -3, 5}) {
		t.Error("Failed set ArrayInt8")
	}

	if !reflect.DeepEqual(c.ArrayInt16, []int16{1, -3, 5}) {
		t.Error("Failed set ArrayInt16")
	}

	if !reflect.DeepEqual(c.ArrayInt32, []int32{1, -3, 5}) {
		t.Error("Failed set ArrayInt32")
	}

	if !reflect.DeepEqual(c.ArrayInt64, []int64{1, -3, 5}) {
		t.Error("Failed set ArrayInt64")
	}
}

func TestInitShouldPanic(t *testing.T) {

	old := mockAbleStringCov

	mockAbleStringCov = func(str string, typeS reflect.Kind) (v interface{}, err error) {
		err = errors.New("error")
		return v, err
	}

	defer func() { mockAbleStringCov = old }()

	type configTestPanic struct {
		Uint uint `env:"TEST_ENV_Uint" envDefault:"10"`
	}

	var res error

	res = Init(&configTestPanic{})

	if res == nil {
		t.Error("Should return error")
	}

	os.Setenv("TEST_ENV_PANIC", "panic")

	type configTestPanicOS struct {
		Panic string `env:"TEST_ENV_PANIC" envDefault:"string"`
	}

	res = Init(&configTestPanicOS{})

	if res == nil {
		t.Error("Should return error")
	}

}

func TestSetUint(t *testing.T) {
	str := testListUint{}

	list := reflect.ValueOf(&str).Elem()

	for i := 0; i < list.NumField(); i++ {
		f := list.Field(i)
		switch f.Kind().String() {

		case "uint":
			setUint(f, uint64(30))
			if str.Uint != uint(30) {
				t.Error("Failed set uint")
			}

		case "uint8":
			setUint8(f, uint64(30))
			if str.Uint8 != uint8(30) {
				t.Error("Failed set uint8")
			}

		case "uint16":
			setUint16(f, uint64(30))
			if str.Uint16 != uint16(30) {
				t.Error("Failed set uint16")
			}

		case "uint32":
			setUint32(f, uint64(30))
			if str.Uint32 != uint32(30) {
				t.Error("Failed set uint32")
			}

		}
	}
}

func TestSetInt(t *testing.T) {
	str := testListInt{}

	list := reflect.ValueOf(&str).Elem()

	for i := 0; i < list.NumField(); i++ {
		f := list.Field(i)
		switch f.Kind().String() {

		case "int":
			setInt(f, int64(30))
			if str.Int != int(30) {
				t.Error("Failed set int")
			}

		case "int8":
			setInt8(f, int64(30))
			if str.Int8 != int8(30) {
				t.Error("Failed set int8")
			}

		case "int16":
			setInt16(f, int64(30))
			if str.Int16 != int16(30) {
				t.Error("Failed set int16")
			}

		case "int32":
			setInt32(f, int64(30))
			if str.Int32 != int32(30) {
				t.Error("Failed set int32")
			}

		}
	}
}
