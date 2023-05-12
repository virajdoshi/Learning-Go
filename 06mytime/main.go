package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	//printing present time using time package
	presentTime := time.Now()
	fmt.Println("Present time is ", presentTime)

	//formated date using time package and we need to use exact this time as an example to format the date otherwise it will be printed a wrong date
	fmt.Println("Present time in formated date is ", presentTime.Format("01-02-2006 15:04:05 Monday January"))

	//created date has different format
	//year is type int
	//month is type time so we need to provide something like time.march
	//day, hour, min, sec, nsec is all int
	//timezone is type time and we need to provide something like time.IST
	createdDate := time.Date(1998, time.December, 2, 12, 12, 0, 0, time.Local)
	fmt.Println("Created date is ", createdDate) //2020-01-01 00:00:00 +0000 UTC
	fmt.Println("Created date in formated date is ", createdDate.Format("01-02-2006 15:04:05 Monday January"))

	//print epoch time
	fmt.Println("Epoch time is ", presentTime.Unix())

	//build an application using go lang
	//GOOS="windows" means you are telling golang to build your application for windows
	//go build will build your program for windows
	//so command will be like GOOS="windows" go build
	//if you already on windows then go build can be used directly
}
