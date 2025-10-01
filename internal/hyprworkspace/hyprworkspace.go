package hyprworkspace

import (
	"fmt"
	"strings"

	"github.com/Xkonti/aeromanager/internal/aerospace"
)

// Execute performs intelligent workspace switching based on cursor position
func Execute(workspaceNum int) error {
	// Validate workspace number (1-5 or 6-0, where 0 is treated as 10)
	if workspaceNum < 0 || workspaceNum > 10 {
		return fmt.Errorf("invalid workspace number: %d (must be 1-5 or 6-0)", workspaceNum)
	}

	// Get current workspace and monitor configuration
	workspaces, monitors, err := aerospace.ListWorkspacesAndMonitors()
	if err != nil {
		return fmt.Errorf("failed to get workspace and monitor info: %w", err)
	}

	// Get the monitor where the mouse cursor is
	mouseMonitorID, err := aerospace.GetMouseMonitorID()
	if err != nil {
		return fmt.Errorf("failed to get mouse monitor: %w", err)
	}

	// Determine which workspace to switch to based on monitor count and cursor position
	var targetWorkspace string

	switch len(monitors) {
	case 1:
		targetWorkspace = mapWorkspaceForSingleMonitor(workspaceNum)
	case 2:
		targetWorkspace = mapWorkspaceForTwoMonitors(workspaceNum, mouseMonitorID, monitors)
	case 3:
		targetWorkspace = mapWorkspaceForThreeMonitors(workspaceNum, mouseMonitorID, monitors)
	default:
		return fmt.Errorf("unsupported monitor configuration: %d monitors", len(monitors))
	}

	// Validate that the target workspace exists
	if !workspaceExists(targetWorkspace, workspaces) {
		return fmt.Errorf("workspace %s does not exist", targetWorkspace)
	}

	fmt.Printf("Switching to workspace %s on monitor %d\n", targetWorkspace, mouseMonitorID)

	// Switch to the target workspace
	return aerospace.SwitchWorkspace(targetWorkspace)
}

// mapWorkspaceForSingleMonitor maps workspace numbers for 1-monitor setup
// 1-5 -> L1-L5, 6-0 -> R1-R5
func mapWorkspaceForSingleMonitor(num int) string {
	if num >= 1 && num <= 5 {
		return fmt.Sprintf("L%d", num)
	}
	// 6-0 maps to R1-R5
	if num == 0 {
		return "R5"
	}
	return fmt.Sprintf("R%d", num-5)
}

// mapWorkspaceForTwoMonitors maps workspace numbers for 2-monitor setup
// Built-in: 1-5 -> A-E, 6-0 -> F-J
// External: 1-5 -> L1-L5, 6-0 -> R1-R5
func mapWorkspaceForTwoMonitors(num int, mouseMonitorID int, monitors []aerospace.Monitor) string {
	// Find if the mouse is on the built-in monitor
	isBuiltIn := false
	for _, mon := range monitors {
		if mon.ID == mouseMonitorID && strings.Contains(mon.Name, "Built-in") {
			isBuiltIn = true
			break
		}
	}

	if isBuiltIn {
		// Built-in monitor: 1-5 -> A-E, 6-0 -> F-J
		if num >= 1 && num <= 5 {
			return string(rune('A' + num - 1))
		}
		// 6-0 maps to F-J
		if num == 0 {
			return "J"
		}
		return string(rune('F' + num - 6))
	}

	// External monitor: same as single monitor
	return mapWorkspaceForSingleMonitor(num)
}

// mapWorkspaceForThreeMonitors maps workspace numbers for 3-monitor setup
// Built-in: 1-5 -> A-E, 6-0 -> F-J
// Left external: 1-5 -> L1-L5, 6-0 -> L1-L5 (wraps around, but typically just 1-5)
// Right external: 1-5 -> R1-R5, 6-0 -> R1-R5 (wraps around, but typically just 1-5)
func mapWorkspaceForThreeMonitors(num int, mouseMonitorID int, monitors []aerospace.Monitor) string {
	// Find the monitor type
	var builtInID int
	var leftExternalID, rightExternalID int
	var externalIDs []int

	for _, mon := range monitors {
		if strings.Contains(mon.Name, "Built-in") {
			builtInID = mon.ID
		} else {
			externalIDs = append(externalIDs, mon.ID)
		}
	}

	// Determine left and right external monitors
	if len(externalIDs) == 2 {
		leftExternalID = externalIDs[0]
		rightExternalID = externalIDs[1]
		if externalIDs[1] < externalIDs[0] {
			leftExternalID = externalIDs[1]
			rightExternalID = externalIDs[0]
		}
	}

	// Map based on which monitor the mouse is on
	if mouseMonitorID == builtInID {
		// Built-in monitor: 1-5 -> A-E, 6-0 -> F-J
		if num >= 1 && num <= 5 {
			return string(rune('A' + num - 1))
		}
		if num == 0 {
			return "J"
		}
		return string(rune('F' + num - 6))
	} else if mouseMonitorID == leftExternalID {
		// Left external: 1-5 -> L1-L5, 6-0 wraps to L1-L5
		if num >= 1 && num <= 5 {
			return fmt.Sprintf("L%d", num)
		}
		if num == 0 {
			return "L5"
		}
		return fmt.Sprintf("L%d", num-5)
	} else if mouseMonitorID == rightExternalID {
		// Right external: 1-5 -> R1-R5, 6-0 wraps to R1-R5
		if num >= 1 && num <= 5 {
			return fmt.Sprintf("R%d", num)
		}
		if num == 0 {
			return "R5"
		}
		return fmt.Sprintf("R%d", num-5)
	}

	// Fallback (should not happen)
	return mapWorkspaceForSingleMonitor(num)
}

// workspaceExists checks if a workspace with the given name exists
func workspaceExists(name string, workspaces []aerospace.Workspace) bool {
	for _, ws := range workspaces {
		if ws.Name == name {
			return true
		}
	}
	return false
}
