package main

import (
	"fmt"

	GoLayouts "github.com/artziel/go-layouts"
)

type MySampleRow struct {
	GoLayouts.Row
	ID       int    `excelLayout:"column:A,required,min:1,unique"`
	Username string `excelLayout:"column:B,required,minLength:6"`
	Password string `excelLayout:"column:C,required,minLength:8"`
	Avatar   string `excelLayout:"column:D,url"`
	Fullname string `excelLayout:"column:E,required"`
	Email    string `excelLayout:"column:F,required,email"`
	Age      int    `excelLayout:"column:G,required,min:18,max:50"`
}

func main() {

	fmt.Println("Sample....")

	l := GoLayouts.ExcelLayout{}

	err := l.ReadFile(MySampleRow{}, "./sample.xlsx")
	if err != nil {
		for _, e := range l.GetErrors() {
			fmt.Printf("Error >> Cell[%v:%v] %s\n", e.Column, e.RowIndex, e.Error.Error())
		}
	} else {
		rows := l.GetRows()
		for i, r := range rows {
			row := r.(*MySampleRow)
			fmt.Printf("%d) ID:%v, Username: %v\n", i, row.ID, row.Username)
		}
	}
}
