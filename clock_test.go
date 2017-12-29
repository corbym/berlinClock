package berlinclock_test

import (
	"testing"
	"berlinclock"
	"gocrest/then"
	"gocrest/is"
)

func TestGivenAClockWhenTimeIsEnteredThenCorrectClock(testing *testing.T) {
	var clockParams = []struct {
		time     string // input
		expected string // expected result
	}{
		{time: "00:00:00", expected: "YOOOOOOOOOOOOOOOOOOOOOOO"},
		{time: "23:59:59", expected: "ORRRRRRROYYRYYRYYRYYYYYY"},
		{time: "16:50:06", expected: "YRRROROOOYYRYYRYYRYOOOOO"},
		{time: "11:37:01", expected: "ORROOROOOYYRYYRYOOOOYYOO"},
	}
	for _, test := range clockParams {
		clock, err := berlinclock.Clock(test.time)
		then.AssertThat(testing, err, is.Nil())
		then.AssertThat(testing, clock, is.
			EqualTo(test.expected).
			Reasonf("time incorrect for %s", test.time))
	}
}

func TestGivenAClockWhenInvalidTimeIsEnteredThenClockError(testing *testing.T) {
	var clockParams = []struct {
		time     string // input
		expected string // expected result
	}{
		{time: "a:b:c", expected: "invalid argument, a:b:c is invalid time format"},
		{time: "99:23:21", expected: "invalid argument"},
		{time: "16:99:06", expected: "invalid argument"},
		{time: "11:37:99", expected: "invalid argument"},
	}
	for _, test := range clockParams {
		clock, err := berlinclock.Clock(test.time)
		then.AssertThat(testing, clock, is.Empty().Reasonf("clock was not empty for given time %s", test.time))
		then.AssertThat(testing, err.Error(), is.EqualTo(test.expected).Reasonf("incorrect error for given time %s", test.time))
	}
}
