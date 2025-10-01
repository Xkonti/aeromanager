package aerospace

import (
	"testing"
)

func TestListWorkspacesAndMonitors(t *testing.T) {
	workspaces, monitors, err := ListWorkspacesAndMonitors()
	if err != nil {
		t.Fatalf("ListWorkspacesAndMonitors() error = %v", err)
	}

	// Expected workspaces based on the example output
	expectedWorkspaces := []Workspace{
		{Name: "0", IsFocused: false, IsVisible: false, MonitorID: 3, MonitorName: "XZ272U P (1)"},
		{Name: "1", IsFocused: true, IsVisible: true, MonitorID: 1, MonitorName: "XZ272U P (2)"},
		{Name: "2", IsFocused: false, IsVisible: false, MonitorID: 1, MonitorName: "XZ272U P (2)"},
		{Name: "3", IsFocused: false, IsVisible: false, MonitorID: 1, MonitorName: "XZ272U P (2)"},
		{Name: "4", IsFocused: false, IsVisible: false, MonitorID: 1, MonitorName: "XZ272U P (2)"},
		{Name: "5", IsFocused: false, IsVisible: false, MonitorID: 1, MonitorName: "XZ272U P (2)"},
		{Name: "6", IsFocused: false, IsVisible: false, MonitorID: 1, MonitorName: "XZ272U P (2)"},
		{Name: "7", IsFocused: false, IsVisible: false, MonitorID: 3, MonitorName: "XZ272U P (1)"},
		{Name: "8", IsFocused: false, IsVisible: true, MonitorID: 3, MonitorName: "XZ272U P (1)"},
		{Name: "9", IsFocused: false, IsVisible: false, MonitorID: 3, MonitorName: "XZ272U P (1)"},
		{Name: "A", IsFocused: false, IsVisible: true, MonitorID: 2, MonitorName: "Built-in Retina Display"},
		{Name: "B", IsFocused: false, IsVisible: false, MonitorID: 2, MonitorName: "Built-in Retina Display"},
		{Name: "C", IsFocused: false, IsVisible: false, MonitorID: 2, MonitorName: "Built-in Retina Display"},
		{Name: "D", IsFocused: false, IsVisible: false, MonitorID: 2, MonitorName: "Built-in Retina Display"},
		{Name: "E", IsFocused: false, IsVisible: false, MonitorID: 2, MonitorName: "Built-in Retina Display"},
	}

	// Expected monitors extracted from workspace data
	expectedMonitors := []Monitor{
		{ID: 1, Name: "XZ272U P (2)"},
		{ID: 2, Name: "Built-in Retina Display"},
		{ID: 3, Name: "XZ272U P (1)"},
	}

	// Test workspaces
	if len(workspaces) != len(expectedWorkspaces) {
		t.Errorf("Got %d workspaces, expected %d", len(workspaces), len(expectedWorkspaces))
	}

	for i, ws := range workspaces {
		if i >= len(expectedWorkspaces) {
			break
		}
		exp := expectedWorkspaces[i]
		if ws.Name != exp.Name {
			t.Errorf("Workspace[%d].Name = %q, expected %q", i, ws.Name, exp.Name)
		}
		if ws.IsFocused != exp.IsFocused {
			t.Errorf("Workspace[%d].IsFocused = %v, expected %v", i, ws.IsFocused, exp.IsFocused)
		}
		if ws.IsVisible != exp.IsVisible {
			t.Errorf("Workspace[%d].IsVisible = %v, expected %v", i, ws.IsVisible, exp.IsVisible)
		}
		if ws.MonitorID != exp.MonitorID {
			t.Errorf("Workspace[%d].MonitorID = %d, expected %d", i, ws.MonitorID, exp.MonitorID)
		}
		if ws.MonitorName != exp.MonitorName {
			t.Errorf("Workspace[%d].MonitorName = %q, expected %q", i, ws.MonitorName, exp.MonitorName)
		}
	}

	// Test monitors
	if len(monitors) != len(expectedMonitors) {
		t.Errorf("Got %d monitors, expected %d", len(monitors), len(expectedMonitors))
	}

	for i, mon := range monitors {
		if i >= len(expectedMonitors) {
			break
		}
		exp := expectedMonitors[i]
		if mon.ID != exp.ID {
			t.Errorf("Monitor[%d].ID = %d, expected %d", i, mon.ID, exp.ID)
		}
		if mon.Name != exp.Name {
			t.Errorf("Monitor[%d].Name = %q, expected %q", i, mon.Name, exp.Name)
		}
	}
}
