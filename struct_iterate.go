package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-sql-driver/mysql"
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
	Number    *int
	Egg       *int
	Name      string  `gorm:"default:'Name'"`
	Alias     *string `gorm:"default:'Alias WOWOW'"`
	IsMan     bool
	Handsome  *bool
	UpdatedAt mysql.NullTime
	CreatedAt *mysql.NullTime
}

func StructIterate() {
	db, _ := getDbConnection()
	luqman := &DummyUser{}
	db.First(&luqman, 1)

	cliff := &DummyUser{}
	db.First(&cliff, 2)

	full := &DummyUser{}
	db.First(&full, 3)

	user := &DummyUser{}

	db.Create(user)

	fmt.Printf("%v\n\n", luqman)
	fmt.Printf("%v\n\n", cliff)
	fmt.Printf("%v\n\n", full)
	// fmt.Printf("luqman alias %s egg %d\n", *luqman.Alias, *luqman.Egg)
	// fmt.Printf("luqman alias %s alias %s\n", full.Name, *full.Alias)

	v := reflect.ValueOf(*full)

	values := make([]interface{}, v.NumField())

	p := &mysql.NullTime{}
	fmt.Printf("%s %v\n", reflect.TypeOf(p), p)
	fmt.Println(p == nil)

	fmt.Println("done")
	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		values[i] = v.Field(i).Interface()

		tipe := reflect.TypeOf(values[i])
		fmt.Printf("%d %s\n", i, reflect.TypeOf(values[i]))
		if strings.HasPrefix(tipe.String(), "*") && !v.Field(i).IsNil() {
			values[i] = v.Field(i).Elem().Interface()
		}
		fmt.Printf("ini %s\n", reflect.TypeOf(values[i]))
		fmt.Printf("%d type %s %s: %v\n", i, name, tipe, values[i])
		if strings.HasSuffix(tipe.String(), "mysql.NullTime") {

			fmt.Printf("time %d %v\n", i, values[i])
			fmt.Printf("%s %v\n", reflect.TypeOf(values[i]), values[i])
			fmt.Println(values[i] == (*mysql.NullTime)(nil))
			if values[i] != (*mysql.NullTime)(nil) {
				fmt.Printf("time %d %v\n", i, values[i])
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
	}

	fmt.Println(values)

}
