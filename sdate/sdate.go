package sdate

import (
	"time"

	xtime "github.com/fynxiu/go-common/time"
)

type SDate string

func (d SDate) String() string {
	return string(d)
}

func Now() SDate {
	return SDate(time.Now().Format(xtime.LayoutStandardShort))
}

func Yesterday() SDate {
	return Now().AddDays(-1)
}

func AWeekBefore() SDate {
	return Now().AddDays(-7)
}

func AMonthBefore() SDate {
	return Now().AddMonths(-1)
}

func AYearBefore() SDate {
	return Now().AddYears(-1)
}

// AddYears
func (d SDate) AddYears(years int) SDate {
	return SDate(d.ToTime().AddDate(years, 0, 0).Format(xtime.LayoutStandardShort))
}

// AddMonths 增加若干月
func (d SDate) AddMonths(months int) SDate {
	return SDate(d.ToTime().AddDate(0, months, 0).Format(xtime.LayoutStandardShort))
}

// AddDays 增加若干天
func (d SDate) AddDays(days int) SDate {
	return SDate(d.ToTime().AddDate(0, 0, days).Format(xtime.LayoutStandardShort))
}

// LessThanNow 是否是在今天之前
func (d SDate) LessThanNow() bool {
	return d.lessThan(Now())
}

// HowManyDaysBeforeNow 距离现在多少天， 昨天 1， 前天 2， 明天 -1
func (d SDate) HowManyDaysBeforeNow() int {
	return int(time.Now().Sub(d.ToTime()).Hours() / 24)
}

func (d SDate) lessThan(date SDate) bool {
	return d < SDate(date)
}

// ToTime 这里要求SDate必须满足 "YYYY-MM-DD" 格式
func (d SDate) ToTime() time.Time {
	if t, err := time.Parse(xtime.LayoutStandardShort, string(d)); err != nil {
		panic(err)
	} else {
		return t
	}
}
