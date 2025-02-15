package battery_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mq/packages/battery"
)

func TestGetBatteryPercent(t *testing.T) {
	want := battery.Status{
		ChargePercent: 40,
	}
	data, e := os.ReadFile("testdata/pmset.txt")
	if e != nil {
		t.Fatal(e)
	}
	got, e := battery.ParsePmsetOutput(string(data))
	if e != nil {
		t.Fatal(e)
	}
	if !cmp.Equal(want, got) {
		//if got := want{
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetStatus(t *testing.T) {
	t.Skip()
	got, e := battery.GetStatus()
	if e != nil {
		t.Fatal(e)
	}
	want := battery.Status{
		ChargePercent: 10,
	}
	if got != want {
		t.Errorf("got %v want %v", got, "")
	}
}
