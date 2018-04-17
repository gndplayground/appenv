package appenv

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var mockAbleStringCov = stringConv

func isZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func stringConv(str string, typeS reflect.Kind) (v interface{}, err error) {

	switch typeS {

	case reflect.String:
		return str, err

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var bitSize int64

		if typeS == reflect.Uint {
			bitSize = 0
		} else {
			bitSize, err = strconv.ParseInt(typeS.String()[len("uint"):], 10, 8)
		}
		v, err = strconv.ParseUint(str, 10, int(bitSize))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var bitSize int64

		if typeS == reflect.Int {
			bitSize = 0
		} else {
			bitSize, err = strconv.ParseInt(typeS.String()[len("int"):], 10, 8)
		}
		v, err = strconv.ParseInt(str, 10, int(bitSize))

	case reflect.Float32, reflect.Float64:

		var bitSize int64

		bitSize, err = strconv.ParseInt(typeS.String()[len("float"):], 10, 8)

		v, err = strconv.ParseFloat(str, int(bitSize))

	case reflect.Bool:
		if str == "true" {
			v = true
		} else if str == "false" {
			v = false
		} else {
			err = errors.New("string convert type bool must be true or false")
		}

	default:
		err = errors.New("not supported type " + typeS.String())
	}
	return v, err
}

func Init(c interface{}) (err error) {

	defer func() {
		if r := recover(); r != nil {
			if reflect.ValueOf(r).Kind() == reflect.String {
				err = errors.New(r.(string))
			} else {
				err = r.(error)
			}
		}
	}()

	listField := reflect.ValueOf(c).Elem()

	for i := 0; i < listField.NumField(); i++ {

		envName := listField.Type().Field(i).Tag.Get("env")

		// Skip if no found env name
		if envName != "" {

			f := listField.Field(i)

			// Field can be set
			if f.CanSet() {

				fType := f.Type().Kind()

				if f.CanInterface() && isZeroOfUnderlyingType(f.Interface()) {

					// Field is a slice
					if fType == reflect.Slice {

						var list string

						valueFromOsEnv, existInOsEnv := os.LookupEnv(envName)

						if existInOsEnv {
							list = valueFromOsEnv
						} else {
							list = listField.Type().Field(i).Tag.Get("envDefault")
						}

						switch {

						case f.Type() == reflect.TypeOf([]string(nil)):
							setValueSliceString(f, list)

						case f.Type() == reflect.TypeOf([]bool(nil)):
							setValueSliceBool(f, list)

						case f.Type() == reflect.TypeOf([]uint(nil)):
							setValueSliceUint(f, list)
						case f.Type() == reflect.TypeOf([]uint8(nil)):
							setValueSliceUint8(f, list)
						case f.Type() == reflect.TypeOf([]uint16(nil)):
							setValueSliceUint16(f, list)
						case f.Type() == reflect.TypeOf([]uint32(nil)):
							setValueSliceUint32(f, list)
						case f.Type() == reflect.TypeOf([]uint64(nil)):
							setValueSliceUint64(f, list)

						case f.Type() == reflect.TypeOf([]int(nil)):
							setValueSliceInt(f, list)
						case f.Type() == reflect.TypeOf([]int8(nil)):
							setValueSliceInt8(f, list)
						case f.Type() == reflect.TypeOf([]int16(nil)):
							setValueSliceInt16(f, list)
						case f.Type() == reflect.TypeOf([]int32(nil)):
							setValueSliceInt32(f, list)
						case f.Type() == reflect.TypeOf([]int64(nil)):
							setValueSliceInt64(f, list)

						case f.Type() == reflect.TypeOf([]float32(nil)):
							setValueSliceFloat32(f, list)
						case f.Type() == reflect.TypeOf([]float64(nil)):
							setValueSliceFloat64(f, list)
						}

					} else {
						var err error
						var finalVal interface{}

						valueFromOsEnv, existInOsEnv := os.LookupEnv(envName)

						if existInOsEnv {
							finalVal, err = mockAbleStringCov(valueFromOsEnv, fType)

							if err != nil {
								panic(err)
							}

						} else {

							finalVal, err = mockAbleStringCov(listField.Type().Field(i).Tag.Get("envDefault"), fType)

							if err != nil {
								panic(err)
							}
						}

						switch fType {

						case reflect.String:
							f.Set(reflect.ValueOf(finalVal))

						case reflect.Uint64, reflect.Int64, reflect.Float64:
							f.Set(reflect.ValueOf(finalVal))

						default:
							setValue(f, finalVal, fType)
						}
					}

				}
			} else {
				panic(errors.New("a field cannot be set. Field with type" + f.Kind().String()))
			}
		}

	}

	return nil
}

func setValueSliceBool(f reflect.Value, list string) {
	var arr []bool

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.String)

		if err != nil {
			panic(err)
		}

		if valC == "true" {
			arr = append(arr, true)
		} else {
			arr = append(arr, false)
		}

	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceString(f reflect.Value, list string) {
	var arr []string

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.String)

		if err != nil {
			panic(err)
		}

		arr = append(arr, valC.(string))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceUint(f reflect.Value, list string) {
	var arr []uint

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Uint)

		if err != nil {
			panic(err)
		}

		arr = append(arr, uint(valC.(uint64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceUint8(f reflect.Value, list string) {
	var arr []uint8

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Uint8)

		if err != nil {
			panic(err)
		}

		arr = append(arr, uint8(valC.(uint64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceUint16(f reflect.Value, list string) {
	var arr []uint16

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Uint16)

		if err != nil {
			panic(err)
		}

		arr = append(arr, uint16(valC.(uint64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceUint32(f reflect.Value, list string) {
	var arr []uint32

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Uint32)

		if err != nil {
			panic(err)
		}

		arr = append(arr, uint32(valC.(uint64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceUint64(f reflect.Value, list string) {
	var arr []uint64

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Uint64)

		if err != nil {
			panic(err)
		}

		arr = append(arr, uint64(valC.(uint64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceInt(f reflect.Value, list string) {
	var arr []int

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Int)

		if err != nil {
			panic(err)
		}

		arr = append(arr, int(valC.(int64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceInt8(f reflect.Value, list string) {
	var arr []int8

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Int8)

		if err != nil {
			panic(err)
		}

		arr = append(arr, int8(valC.(int64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceInt16(f reflect.Value, list string) {
	var arr []int16

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Int16)

		if err != nil {
			panic(err)
		}

		arr = append(arr, int16(valC.(int64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceInt32(f reflect.Value, list string) {
	var arr []int32

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Int32)

		if err != nil {
			panic(err)
		}

		arr = append(arr, int32(valC.(int64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceInt64(f reflect.Value, list string) {
	var arr []int64

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Int64)

		if err != nil {
			panic(err)
		}

		arr = append(arr, int64(valC.(int64)))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValueSliceFloat32(f reflect.Value, list string) {
	var arr []float32

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Float32)
		if err != nil {
			panic(err)
		}

		arr = append(arr, float32(valC.(float64)))
	}
	f.Set(reflect.ValueOf(arr))
}

func setValueSliceFloat64(f reflect.Value, list string) {
	var arr []float64

	for _, val := range strings.Split(list, ",") {

		valC, err := stringConv(val, reflect.Float64)
		if err != nil {
			panic(err)
		}

		arr = append(arr, valC.(float64))
	}

	f.Set(reflect.ValueOf(arr))
}

func setValue(f reflect.Value, v interface{}, t reflect.Kind) {
	switch t {
	case reflect.Uint:
		setUint(f, v)
	case reflect.Uint8:
		setUint8(f, v)
	case reflect.Uint16:
		setUint16(f, v)
	case reflect.Uint32:
		setUint32(f, v)
	case reflect.Int:
		setInt(f, v)
	case reflect.Int8:
		setInt8(f, v)
	case reflect.Int16:
		setInt16(f, v)
	case reflect.Int32:
		setInt32(f, v)
	case reflect.Float32:
		setFloat32(f, v)
	}
}

func setUint(f reflect.Value, v interface{}) {
	f.Set(reflect.ValueOf(uint(v.(uint64))))
}

func setUint8(f reflect.Value, v interface{}) {
	f.Set(reflect.ValueOf(uint8(v.(uint64))))
}

func setUint16(f reflect.Value, v interface{}) {
	f.Set(reflect.ValueOf(uint16(v.(uint64))))
}

func setUint32(f reflect.Value, v interface{}) {
	f.Set(reflect.ValueOf(uint32(v.(uint64))))
}

func setInt(f reflect.Value, v interface{}) {
	f.Set(reflect.ValueOf(int(v.(int64))))
}

func setInt8(f reflect.Value, v interface{}) {
	f.Set(reflect.ValueOf(int8(v.(int64))))
}

func setInt16(f reflect.Value, v interface{}) {
	f.Set(reflect.ValueOf(int16(v.(int64))))
}

func setInt32(f reflect.Value, v interface{}) {
	f.Set(reflect.ValueOf(int32(v.(int64))))
}

func setFloat32(f reflect.Value, v interface{}) {
	f.Set(reflect.ValueOf(float32(v.(float64))))
}
