package snowflake

import (
	xtime "github.com/fynxiu/go-common/time"
	"time"
)

func ExampleSnowflake_NextID() {
	startTime, _ := time.Parse(xtime.LayoutStandardShort, "2019-09-25")
	sf, err := NewSnowflake(Settings{
		StartTime: startTime,
		MachineID: func() (u uint16, e error) {
			return 77, nil
		},
	})
	if err != nil {
		panic(err)
	}

	_, err = sf.NextID()
	if err != nil {
		panic(err)
	}

	// output:
}
