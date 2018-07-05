package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Option holds all necessary options for database.
type MySQLOption struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

type DummyUser struct {
	ID        int
	Number    int
	Egg       *int
	Name      string
	Alias     *string
	IsMan     bool
	Handsome  *bool
	UpdatedAt mysql.NullTime
	CreatedAt *mysql.NullTime
}

func getDbConnection() (db *gorm.DB, err error) {
	d, err := gorm.Open("mysql", "luqman:luqman@tcp(127.0.0.1:3306)/test?charset=utf8")
	return d, err
}

func StructIterate() {
	db, _ := getDbConnection()
	luqman := &DummyUser{}
	db.First(&luqman, 1)

	cliff := &DummyUser{}
	db.First(&cliff, 2)

	full := &DummyUser{}
	db.First(&full, 3)

	fmt.Printf("%v\n\n", luqman)
	fmt.Printf("%v\n\n", cliff)
	fmt.Printf("%v\n\n", full)
	fmt.Printf("luqman alias %s egg %d\n", *luqman.Alias, *luqman.Egg)
	fmt.Printf("luqman alias %s alias %s\n", full.Name, *full.Alias)

	v := reflect.ValueOf(*full)

	values := make([]interface{}, v.NumField())

	fmt.Println("done")
	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		values[i] = v.Field(i).Interface()

		tipe := reflect.TypeOf(values[i])
		if strings.HasPrefix(tipe.String(), "*") {
			values[i] = v.Field(i).Elem().Interface()
		}
		fmt.Printf("%d type %s %s: %v\n", i+1, name, tipe, values[i])
		if strings.HasSuffix(tipe.String(), "mysql.NullTime") {
			vs := reflect.ValueOf(values[i])
			vals := make([]interface{}, vs.NumField())

			for j := 0; j < vs.NumField(); j++ {
				name_d := vs.Type().Field(j).Name
				vals[j] = vs.Field(j).Interface()
				tipe_d := reflect.TypeOf(vals[j])

				fmt.Printf(" %d type %s %s: %v\n", j+1, name_d, tipe_d, vals[j])
			}
		}
	}

	fmt.Println(values)

}
