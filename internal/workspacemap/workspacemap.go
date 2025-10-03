package workspacemap

import (
	"fmt"
	"strings"

	"github.com/Xkonti/aeromanager/internal/aerospace"
)

// MapWorkspaceNumber maps a workspace number (1-5 or 6-0) to a workspace name
// based on the monitor configuration and which monitor is targeted.
func MapWorkspaceNumber(workspaceNum int, targetMonitorID int, monitors []aerospace.Monitor) string {
	switch len(monitors) {
	case 1:
		return mapForSingleMonitor(workspaceNum)
	case 2:
		return mapForTwoMonitors(workspaceNum, targetMonitorID, monitors)
	case 3:
		return mapForThreeMonitors(workspaceNum, targetMonitorID, monitors)
	default:
		// Fallback to single monitor mapping
		return mapForSingleMonitor(workspaceNum)
	}
}

// mapForSingleMonitor maps workspace numbers for 1-monitor setup
// 1-5 -> L1-L5, 6-0 -> R1-R5
func mapForSingleMonitor(num int) string {
	if num >= 1 && num <= 5 {
		return fmt.Sprintf("L%d", num)
	}
	// 6-0 maps to R1-R5
	if num == 0 {
		return "R5"
	}
	return fmt.Sprintf("R%d", num-5)
}

// mapForTwoMonitors maps workspace numbers for 2-monitor setup
// Built-in: 1-5 -> B1-B5, 6-0 wraps to B1-B5
// External: 1-5 -> L1-L5, 6-0 -> R1-R5
func mapForTwoMonitors(num int, targetMonitorID int, monitors []aerospace.Monitor) string {
	// Find if the target is the built-in monitor
	isBuiltIn := false
	for _, mon := range monitors {
		if mon.ID == targetMonitorID && strings.Contains(mon.Name, "Built-in") {
			isBuiltIn = true
			break
		}
	}

	if isBuiltIn {
		// Built-in monitor: 1-5 -> B1-B5, 6-0 wraps to B1-B5
		if num >= 1 && num <= 5 {
			return fmt.Sprintf("B%d", num)
		}
		// 6-0 wraps to B1-B5
		if num == 0 {
			return "B5"
		}
		return fmt.Sprintf("B%d", num-5)
	}

	// External monitor: same as single monitor
	return mapForSingleMonitor(num)
}

// mapForThreeMonitors maps workspace numbers for 3-monitor setup
// Built-in: 1-5 -> B1-B5, 6-0 wraps to B1-B5
// Left external: 1-5 -> L1-L5, 6-0 wraps to L1-L5
// Right external: 1-5 -> R1-R5, 6-0 wraps to R1-R5
func mapForThreeMonitors(num int, targetMonitorID int, monitors []aerospace.Monitor) string {
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

	// Map based on which monitor the target is
	if targetMonitorID == builtInID {
		// Built-in monitor: 1-5 -> B1-B5, 6-0 wraps to B1-B5
		if num >= 1 && num <= 5 {
			return fmt.Sprintf("B%d", num)
		}
		if num == 0 {
			return "B5"
		}
		return fmt.Sprintf("B%d", num-5)
	} else if targetMonitorID == leftExternalID {
		// Left external: 1-5 -> L1-L5, 6-0 wraps to L1-L5
		if num >= 1 && num <= 5 {
			return fmt.Sprintf("L%d", num)
		}
		if num == 0 {
			return "L5"
		}
		return fmt.Sprintf("L%d", num-5)
	} else if targetMonitorID == rightExternalID {
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
	return mapForSingleMonitor(num)
}