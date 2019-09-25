package snowflake

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/deckarep/golang-set"
)

var sf *Snowflake

var startTime int64
var machineID uint64

func init() {
	var st Settings
	st.StartTime = time.Now()

	var err error
	sf, err = NewSnowflake(st)
	if err != nil {
		panic("Snowflake not created")
	}

	startTime = toSnowflakeTime(st.StartTime)

	ip, _ := lower16BitPrivateIP()
	machineID = uint64(ip)
}

func nextID(t *testing.T) uint64 {
	id, err := sf.NextID()
	if err != nil {
		t.Fatal("id not generated")
	}
	return id
}

func TestSnowflakeOnce(t *testing.T) {
	sleepTime := uint64(50)
	time.Sleep(time.Duration(sleepTime) * 10 * time.Millisecond)

	id := nextID(t)
	parts := Decompose(id)

	actualMSB := parts["msb"]
	if actualMSB != 0 {
		t.Errorf("unexpected msb: %d", actualMSB)
	}

	actualTime := parts["time"]
	if actualTime < sleepTime || actualTime > sleepTime+1 {
		t.Errorf("unexpected time: %d", actualTime)
	}

	actualSequence := parts["sequence"]
	if actualSequence != 0 {
		t.Errorf("unexpected sequence: %d", actualSequence)
	}

	actualMachineID := parts["machine-id"]
	if actualMachineID != machineID {
		t.Errorf("unexpected machine id: %d", actualMachineID)
	}

	fmt.Println("Snowflake id:", id)
	fmt.Println("decompose:", parts)
}

func currentTime() int64 {
	return toSnowflakeTime(time.Now())
}

func TestSnowflakeFor10Sec(t *testing.T) {
	var numID uint32
	var lastID uint64
	var maxSequence uint64

	initial := currentTime()
	current := initial
	for current-initial < 1000 {
		id := nextID(t)
		parts := Decompose(id)
		numID++

		if id <= lastID {
			t.Fatal("duplicated id")
		}
		lastID = id

		current = currentTime()

		actualMSB := parts["msb"]
		if actualMSB != 0 {
			t.Errorf("unexpected msb: %d", actualMSB)
		}

		actualTime := int64(parts["time"])
		overtime := startTime + actualTime - current
		if overtime > 0 {
			t.Errorf("unexpected overtime: %d", overtime)
		}

		actualSequence := parts["sequence"]
		if maxSequence < actualSequence {
			maxSequence = actualSequence
		}

		actualMachineID := parts["machine-id"]
		if actualMachineID != machineID {
			t.Errorf("unexpected machine id: %d", actualMachineID)
		}
	}

	if maxSequence != 1<<BitLenSequence-1 {
		t.Errorf("unexpected max sequence: %d", maxSequence)
	}
	fmt.Println("max sequence:", maxSequence)
	fmt.Println("number of id:", numID)
}

func TestSnowflakeInParallel(t *testing.T) {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Println("number of cpu:", numCPU)

	consumer := make(chan uint64)

	const numID = 10000
	generate := func() {
		for i := 0; i < numID; i++ {
			consumer <- nextID(t)
		}
	}

	const numGenerator = 10
	for i := 0; i < numGenerator; i++ {
		go generate()
	}

	set := mapset.NewSet()
	for i := 0; i < numID*numGenerator; i++ {
		id := <-consumer
		if set.Contains(id) {
			t.Fatal("duplicated id")
		} else {
			set.Add(id)
		}
	}
	fmt.Println("number of id:", set.Cardinality())
}

func TestNilSnowflake(t *testing.T) {
	var startInFuture Settings
	startInFuture.StartTime = time.Now().Add(time.Duration(1) * time.Minute)
	if _, err := NewSnowflake(startInFuture); err == nil {
		t.Errorf("Snowflake starting in the future")
	}

	var noMachineID Settings
	noMachineID.MachineID = func() (uint16, error) {
		return 0, fmt.Errorf("no machine id")
	}
	if _, err := NewSnowflake(noMachineID); err == nil {
		t.Errorf("Snowflake with no machine id")
	}

	var invalidMachineID Settings
	invalidMachineID.CheckMachineID = func(uint16) bool {
		return false
	}
	if _, err := NewSnowflake(invalidMachineID); err == nil {
		t.Errorf("Snowflake with invalid machine id")
	}
}

func pseudoSleep(period time.Duration) {
	sf.startTime -= int64(period) / snowflakeTimeUnit
}

func TestNextIDError(t *testing.T) {
	year := time.Duration(365*24) * time.Hour
	pseudoSleep(time.Duration(174) * year)
	nextID(t)

	pseudoSleep(time.Duration(1) * year)
	_, err := sf.NextID()
	if err == nil {
		t.Errorf("time is not over")
	}
}
