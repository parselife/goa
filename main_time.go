package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()

	fmt.Println(now)

	fmt.Printf("today is %d \n", now.Day())
	fmt.Printf("month is %s\n", now.Month().String())
	fmt.Printf("weekday is %s\n", now.Weekday().String())
	tomorrow := now.Add(time.Hour*24)
	fmt.Printf("明天: %s\n", tomorrow.String())
	yesterday := now.AddDate(0, 0, -1)
	fmt.Println(yesterday.Date())
	//格式化时间为字符串
	t3 := now.Format("2006-01-02 15:04:05")
	fmt.Println(t3)
	//格式化时间为字符串,只保留年月日
	t3 = now.Format("2006-01-02")
	fmt.Println(t3)

}
