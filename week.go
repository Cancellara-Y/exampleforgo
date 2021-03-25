package main

import (
	"fmt"
	"time"
)

//判断时间是当年的第几周

func WeekByDate(t time.Time) string {

	yearDay := t.YearDay()

	yearFirstDay := t.AddDate(0, 0, -yearDay+1)

	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天

	firstWeekDays := 1

	if firstDayInWeek != 0 {

		firstWeekDays = 7 - firstDayInWeek + 1

	}

	var week int

	if yearDay <= firstWeekDays {

		week = 1

	} else {

		week = (yearDay-firstWeekDays)/7 + 2

	}

	return fmt.Sprintf("%d第%d周", t.Year(), week)

}

func main() {
	fmt.Println(WeekByDate(time.Now()))
	fmt.Println(WeekByDate(time.Now().AddDate(0, 0, 3)))

}
