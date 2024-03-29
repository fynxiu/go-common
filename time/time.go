package time

import (
	"context"
	"database/sql/driver"
	"strconv"
	xtime "time"
)

const (
	LayoutStandard       = "2006-01-02 15:04:05"
	LayoutStandardMinute = "2006-01-02 15:04"
	LayoutStandardMil    = "2006-01-02 15:04:05.999999"
	LayoutStandardShort  = "2006-01-02"
	LayoutSlash          = "2006/01/02 15:04:05"
	LayoutSlashShort     = "2006/01/02"
)

// Time be used to MySql timestamp converting.
type Time int64

// Scan scan time.
func (jt *Time) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xtime.Time:
		*jt = Time(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*jt = Time(i)
	}
	return
}

// Value get time value.
func (jt Time) Value() (driver.Value, error) {
	return xtime.Unix(int64(jt), 0), nil
}

// Time get time.
func (jt Time) Time() xtime.Time {
	return xtime.Unix(int64(jt), 0)
}

// Duration be used toml unmarshal string time, like 1s, 500ms.
type Duration xtime.Duration

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := xtime.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

// Shrink will decrease the duration by comparing with context's timeout duration
// and return new timeout\context\CancelFunc.
func (d Duration) Shrink(c context.Context) (Duration, context.Context, context.CancelFunc) {
	if deadline, ok := c.Deadline(); ok {
		if cTimeout := xtime.Until(deadline); cTimeout < xtime.Duration(d) {
			// deliver small timeout
			return Duration(cTimeout), c, func() {}
		}
	}
	ctx, cancel := context.WithTimeout(c, xtime.Duration(d))
	return d, ctx, cancel
}
