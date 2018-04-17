# appenv 

[![Build Status](https://travis-ci.org/gndplayground/appenv.svg?branch=master)](https://travis-ci.org/gndplayground/appenv)
[![Coverage Status](https://coveralls.io/repos/github/gndplayground/appenv/badge.svg?branch=master)](https://coveralls.io/github/gndplayground/appenv?branch=master)

## What can it do

When we deploy in some service like AWS we often set our config in environment variables then retrieve it

```go
value, exist := os.LookupEnv(envName)
```

However, we need to group these environment variables to a single struct and also need to handle default environment variables (most likely 
default variables for local developing).

By using struct tag we can define beauty config for our app

```go

type config = struct {
	Port       uint16 `env:"APP_PORT" envDefault:"7000"`
	DBUser     string `env:"DB_USER" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"root"`
	DBDefault  string `env:"DB_DEFAULT" envDefault:"awesome"`
	DBPort     uint16 `env:"DB_PORT" envDefault:"7706"`
}

```

When we use this config with the package, for example, field

```
Port       uint16 `env:"APP_PORT" envDefault:"7000"`
```

The package will set field `Port` with the environment variable `APP_PORT` if not found will use the default `7000`. 
The default value or environment variable will be converted to correct type. In this case, it will be converted to `uint16`

## How to use

Bear in mind the current folder structure just for my app. Your project structure might be different than mine.

### First, define a config

File project/app/config.go

```go
package app

var Config config

type config = struct {
	Port       uint16 `env:"APP_PORT" envDefault:"7000"`
	DBUser     string `env:"DB_USER" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"root"`
	DBDefault  string `env:"DB_DEFAULT" envDefault:"awesome"`
	DBPort     uint16 `env:"DB_PORT" envDefault:"7706"`
}
```

### Then use it in you main file

File project/main.go

```go

package main

import (
	"github.com/gndplayground/appenv"
	"project/app"
)

func main() {

	appenv.Init(&app.Config)

	app.DBInit()

  ...
}


```

Then other package in app can use the config

File project/app/db.go

```go


func DBInit() {
	var err error

	DB, err = gorm.Open("mysql", fmt.Sprintf(
		"%v:%v@(localhost:%v)/%v?charset=utf8&parseTime=True",
		Config.DBUser,
		Config.DBPassword,
		Config.DBPort,
		Config.DBDefault))

	defer func() {
		if err != nil {
			DB.Close()
			panic(err)
		}
	}()
}

```

## Support Type

```go

type config struct {
		Uint         uint      `env:"TEST_ENV_Uint" envDefault:"10"`
		Uint8        uint8     `env:"TEST_ENV_Uint8" envDefault:"10"`
		Uint16       uint16    `env:"TEST_ENV_Uint16" envDefault:"10"`
		Uint32       uint32    `env:"TEST_ENV_Uint32" envDefault:"10"`
		Uint64       uint64    `env:"TEST_ENV_Uint64" envDefault:"10"`
		Int          int       `env:"TEST_ENV_Int" envDefault:"-10"`
		Int8         int8      `env:"TEST_ENV_Int8" envDefault:"-10"`
		Int16        int16     `env:"TEST_ENV_Int16" envDefault:"-10"`
		Int32        int32     `env:"TEST_ENV_Int32" envDefault:"-10"`
		Int64        int64     `env:"TEST_ENV_Int64" envDefault:"-10"`
		Float32      float32   `env:"TEST_ENV_FLOAT32" envDefault:"3.2"`
		Float64      float64   `env:"TEST_ENV_FLOAT32" envDefault:"6.5"`
		String       string    `env:"TEST_ENV_STRING" envDefault:"hello"`
		ArrayString  []string  `env:"TEST_ENV_ArrayString" envDefault:"a,b,c"`
		ArrayBool    []bool    `env:"TEST_ENV_ArrayBool" envDefault:"true,true,false"`
		ArrayFloat64 []float64 `env:"TEST_ENV_ArrayFloat64" envDefault:"5.4,5.1,5.3"`
		ArrayFloat32 []float32 `env:"TEST_ENV_ArrayFloat32" envDefault:"5.4,5.1,5.3"`
		ArrayUint    []uint    `env:"TEST_ENV_ArrayUint" envDefault:"1,3,5"`
		ArrayUint8   []uint8   `env:"TEST_ENV_ArrayUint8" envDefault:"1,3,5"`
		ArrayUint16  []uint16  `env:"TEST_ENV_ArrayUint16" envDefault:"1,3,5"`
		ArrayUint32  []uint32  `env:"TEST_ENV_ArrayUint32" envDefault:"1,3,5"`
		ArrayUint64  []uint64  `env:"TEST_ENV_ArrayUint64" envDefault:"1,3,5"`
		ArrayInt     []int     `env:"TEST_ENV_ArrayInt" envDefault:"1,-3,5"`
		ArrayInt8    []int8    `env:"TEST_ENV_ArrayInt8" envDefault:"1,-3,5"`
		ArrayInt16   []int16   `env:"TEST_ENV_ArrayInt16" envDefault:"1,-3,5"`
		ArrayInt32   []int32   `env:"TEST_ENV_ArrayInt32" envDefault:"1,-3,5"`
		ArrayInt64   []int64   `env:"TEST_ENV_ArrayInt64" envDefault:"1,-3,5"`
	}

```
