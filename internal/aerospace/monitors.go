package aerospace

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Monitor represents an Aerospace monitor with its properties
type Monitor struct {
	ID   int    // 1-based sequential number of the monitor
	Name string // Name of the monitor
}

// ListMonitors executes the aerospace list-monitors command and returns
// an array of Monitor objects ordered from left to right as arranged in macOS settings
func ListMonitors() ([]Monitor, error) {
	cmd := exec.Command("aerospace", "list-monitors", "--format", "%{monitor-id}|%{monitor-name}")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute aerospace list-monitors: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	monitors := make([]Monitor, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid monitor output format: %s", line)
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid monitor ID: %s", parts[0])
		}

		monitors = append(monitors, Monitor{
			ID:   id,
			Name: parts[1],
		})
	}

	return monitors, nil
}

// GetMouseMonitorID returns the ID of the monitor that currently has the mouse cursor
func GetMouseMonitorID() (int, error) {
	cmd := exec.Command("aerospace", "list-monitors", "--mouse", "--format", "%{monitor-id}")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to execute aerospace list-monitors --mouse: %w", err)
	}

	idStr := strings.TrimSpace(string(output))
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid monitor ID: %s", idStr)
	}

	return id, nil
}
