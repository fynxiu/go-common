package sdate

import (
	"testing"
)

func TestSDate_LessThan(t *testing.T) {
	const d1 = "2006-06-20"
	const d2 = "2006-06-21"
	const d3 = "2006-06-22"
	type args struct {
		date SDate
	}
	tests := []struct {
		name string
		d    SDate
		args args
		want bool
	}{
		{
			name: "later date should be false",
			d:    SDate(d3),
			args: args{date: d2},
			want: false,
		},
		{
			name: "earlier date should be false",
			d:    SDate(d1),
			args: args{date: d2},
			want: true,
		},
		{
			name: "same date should be false",
			d:    SDate(d1),
			args: args{date: d1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.LessThan(tt.args.date); got != tt.want {
				t.Errorf("SDate.lessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSDate_AddMonths(t *testing.T) {
	const d1 = SDate("2006-06-20")
	type args struct {
		months int
	}
	tests := []struct {
		name string
		d    SDate
		args args
		want SDate
	}{
		{
			name: "add one month",
			d:    d1,
			args: args{1},
			want: SDate("2006-07-20"),
		},
		{
			name: "add multi months in a year",
			d:    d1,
			args: args{5},
			want: SDate("2006-11-20"),
		},
		{
			name: "add multi months across year",
			d:    d1,
			args: args{12},
			want: SDate("2007-06-20"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AddMonths(tt.args.months); got != tt.want {
				t.Errorf("SDate.AddMonths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSDate_AddDays(t *testing.T) {
	const d1 = SDate("2006-07-20")
	type args struct {
		days int
	}
	tests := []struct {
		name string
		d    SDate
		args args
		want SDate
	}{
		{
			name: "add one day",
			d:    d1,
			args: args{1},
			want: SDate("2006-07-21"),
		},
		{
			name: "add multi days in a month",
			d:    d1,
			args: args{5},
			want: SDate("2006-07-25"),
		},
		{
			name: "add multi days across month",
			d:    d1,
			args: args{12},
			want: SDate("2006-08-01"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AddDays(tt.args.days); got != tt.want {
				t.Errorf("SDate.AddDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
