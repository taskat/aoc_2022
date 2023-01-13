package timer

import (
	"fmt"
	"time"
)

func Start(s string) (string, time.Time) {
    return s, time.Now()
}

func TraceSingleCall(s string, startTime time.Time) {
    endTime := time.Now()
    fmt.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

type aggregateCall struct {
	times map[string]data
}

type data struct {
	d time.Duration
	calls int
}

func (d *data) update(t time.Duration) {
	d.d += t
	d.calls++
}

func (d data) String() string {
	return fmt.Sprintf("ElapsedTime in seconds: %v, calls: %v", d.d, d.calls)
}

var aggregator = aggregateCall{times: make(map[string]data)}

func TraceMultipleCalls(s string, startTime time.Time) {
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	d := aggregator.times[s]
	d.update(elapsedTime)
	aggregator.times[s] = d
}

func PrintAggregateTimes() {
	fmt.Println("AGGREGATE TIMES:")
	for k, v := range aggregator.times {
		fmt.Println("  ", k, ":", v)
	}
}
