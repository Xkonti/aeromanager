package aerospace

import (
	"testing"
)

func TestListMonitors(t *testing.T) {
	monitors, err := ListMonitors()
	if err != nil {
		t.Fatalf("ListMonitors() error = %v", err)
	}

	expected := []Monitor{
		{ID: 1, Name: "XZ272U P (2)"},
		{ID: 2, Name: "Built-in Retina Display"},
		{ID: 3, Name: "XZ272U P (1)"},
	}

	if len(monitors) != len(expected) {
		t.Errorf("ListMonitors() returned %d monitors, expected %d", len(monitors), len(expected))
	}

	for i, monitor := range monitors {
		if i >= len(expected) {
			break
		}
		if monitor.ID != expected[i].ID {
			t.Errorf("Monitor[%d].ID = %d, expected %d", i, monitor.ID, expected[i].ID)
		}
		if monitor.Name != expected[i].Name {
			t.Errorf("Monitor[%d].Name = %q, expected %q", i, monitor.Name, expected[i].Name)
		}
	}
}

func TestGetMouseMonitorID(t *testing.T) {
	id, err := GetMouseMonitorID()
	if err != nil {
		t.Fatalf("GetMouseMonitorID() error = %v", err)
	}

	if id != 1 {
		t.Errorf("GetMouseMonitorID() = %d, expected 1", id)
	}
}
