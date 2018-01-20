package ld

import (
	"context"
	"time"
)

const (
	location   string = "Asia/Jakarta"
	dateFormat string = "2006-01-02"
	offset     int    = 1

	myLDDate LDDate = "2018-02-09"
)

type (
	LDDate string
)

var (
	myLD LDDate
)

func (d LDDate) GetTime(ctx context.Context) (time.Time, error) {
	var ldTime time.Time

	loc, err := time.LoadLocation(location)
	if err != nil {
		return ldTime, err
	}

	ldTime, err = time.ParseInLocation(dateFormat, string(myLDDate), loc)
	if err != nil {
		return ldTime, err
	}

	return ldTime, nil
}

func (d LDDate) CountLeft(ctx context.Context, current time.Time) (int, error) {
	ldTime, err := d.GetTime(ctx)
	if err != nil {
		return 0, err
	}

	dur := ldTime.Sub(current)
	return int(dur.Hours()/24) + offset, nil
}

func init() {
	myLD = myLDDate
}

func Get() LDDate {
	return myLD
}
