# GOrm

https://github.com/jinzhu/gorm

## 示例

main.go
```go
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"lastn_ame"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "./gorm.db")
	defer db.Close()

	db.AutoMigrate(&Person{})

	p1 := Person{FirstName: "John", LastName: "Doe"}
	p2 := Person{FirstName: "Jane", LastName: "Smith"}

	db.Create(&p1)
	db.Create(&p2)

	fmt.Println(p1.FirstName, p1.LastName, )
	fmt.Println(p2.FirstName, p2.LastName, )

	var p3 Person
	db.First(&p3)

	fmt.Println(p3.FirstName, p3.LastName, )
}
```