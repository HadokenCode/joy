package main

import (
	"github.com/matthewmueller/golly/testdata/34-keyless-fields/dep/dep"
)

// Phone struct
type Phone struct {
	Type   string
	Number string `js:"number"`
}

// User struct
type User struct {
	Name     string `js:"name"`
	Age      int
	Phone    *Phone `js:"phone"`
	Settings *dep.Settings
}

func main() {
	settings := dep.Settings{Place: "USA"}
	phone := Phone{"Android", "8674205"}
	age := 28
	u := User{"matt", age, &phone, &settings}
	println(u.Name, u.Age, u.Phone.Number, u.Settings.Place)
}
