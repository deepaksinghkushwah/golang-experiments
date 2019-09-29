package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type Blog struct {
	ID      int
	Title   string
	Content string
}

func main() {
	db, err := gorm.Open("mysql", "root:deepak@/go_orm_test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}

	blogs := []Blog{}
	db.Where("id > 4").Limit(10).Find(&blogs)
	for _, v := range blogs {
		fmt.Println(v.Title)
		fmt.Println("----------------------------------")
	}
}
