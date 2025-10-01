package aerospace

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Workspace represents an Aerospace workspace with its properties
type Workspace struct {
	Name       string // Name of the workspace
	IsFocused  bool   // True if the workspace has focus
	IsVisible  bool   // True if the workspace is visible
	MonitorID  int    // 1-based sequential number of the belonging monitor
	MonitorName string // Name of the belonging monitor
}

// ListWorkspacesAndMonitors executes the aerospace list-workspaces command and returns
// both workspaces and monitors. Monitors are extracted from the workspace data and
// ordered from left to right as arranged in macOS settings (by monitor ID).
func ListWorkspacesAndMonitors() ([]Workspace, []Monitor, error) {
	cmd := exec.Command("aerospace", "list-workspaces", "--all", "--format",
		"%{workspace}|%{workspace-is-focused}|%{workspace-is-visible}|%{monitor-id}|%{monitor-name}")
	output, err := cmd.Output()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute aerospace list-workspaces: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	workspaces := make([]Workspace, 0, len(lines))
	monitorMap := make(map[int]string) // ID -> Name

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "|", 5)
		if len(parts) != 5 {
			return nil, nil, fmt.Errorf("invalid workspace output format: %s", line)
		}

		isFocused, err := strconv.ParseBool(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid workspace-is-focused value: %s", parts[1])
		}

		isVisible, err := strconv.ParseBool(parts[2])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid workspace-is-visible value: %s", parts[2])
		}

		monitorID, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid monitor ID: %s", parts[3])
		}

		workspaces = append(workspaces, Workspace{
			Name:        parts[0],
			IsFocused:   isFocused,
			IsVisible:   isVisible,
			MonitorID:   monitorID,
			MonitorName: parts[4],
		})

		// Collect unique monitors
		monitorMap[monitorID] = parts[4]
	}

	// Convert monitor map to sorted slice (by ID, which represents left-to-right order)
	monitors := make([]Monitor, 0, len(monitorMap))
	for id := 1; id <= len(monitorMap); id++ {
		if name, exists := monitorMap[id]; exists {
			monitors = append(monitors, Monitor{
				ID:   id,
				Name: name,
			})
		}
	}

	return workspaces, monitors, nil
}
