# Golang Layouts Files Package
Golang version 1.18

Artziel Narvaiza <artziel@gmail.com>

### Features
- Simplify Excel Layouts reading 
- Validate read row data 

### Dependencies
- github.com/xuri/excelize/v2

Get the package
```bash
go get github.com/artziel/go-layouts
```

Use example:
```golang
package main

import (
	"fmt"

	DataLayouts "github.com/artziel/go-layouts"
)

type MySampleRow struct {
	DataLayouts.Row
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

	l := DataLayouts.ExcelLayout{}

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
```
