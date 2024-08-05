package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

func findYearAndRemainingDays(ts int) (int, int) {
	secondsInDay := 24 * 60 * 60
	days := ts / secondsInDay
	year := 1970

	for (days - 365) > 0 {
		days -= 365
		if isLeapYear(year) {
			days -= 1
		}
		year++
	}

	return year, days
}

func findMonthAndDay(year int, daysOfYear int) (int, int) {
	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if isLeapYear(year) {
		monthDays[1] = 29
	}

	month := 0
	for i, days := range monthDays {
		if daysOfYear < days {
			month = i + 1
			break
		}
		daysOfYear -= days
	}

	return month, daysOfYear + 1
}

func convertUnixTimestamp(ts int) string {
	year, dayOfYear := findYearAndRemainingDays(ts)
	month, day := findMonthAndDay(year, dayOfYear)

	remainingSeconds := ts % (24 * 60 * 60)
	hour := remainingSeconds / 3600
	remainingSeconds %= 3600
	minute := remainingSeconds / 60
	second := remainingSeconds % 60

	return fmt.Sprintf("Year: %d, Month: %02d, Day: %02d, Time: %02d:%02d:%02d",
		year, month, day, hour, minute, second)
}

func main() {
	ts := 1722841336
	formattedDate := convertUnixTimestamp(ts)
	fmt.Println(formattedDate)
}
