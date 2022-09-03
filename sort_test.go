package sleepsort

import (
	"reflect"
	"testing"
)

func TestSleepSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test...")
	}
	var (
		input    = report(t, "Input:   ", 3, 1, 4, 1, 5, 9)
		expected = report(t, "Expected:", 1, 1, 3, 4, 5, 9)
		actual   = report(t, "Actual:  ", Sort(input...)...)
	)
	if !reflect.DeepEqual(expected, actual) {
		t.Error("nope")
	}
}
func report(t *testing.T, label string, values ...uint) []uint {
	t.Logf(label+" %+v", values)
	return values
}
