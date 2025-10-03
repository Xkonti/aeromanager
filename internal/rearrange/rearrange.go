package rearrange

import (
	"fmt"
	"strings"

	"github.com/Xkonti/aeromanager/internal/aerospace"
)

// Execute performs the workspace rearrangement based on monitor setup
func Execute() error {
	// Get current workspace and monitor configuration
	workspaces, monitors, err := aerospace.ListWorkspacesAndMonitors()
	if err != nil {
		return fmt.Errorf("failed to get workspace and monitor info: %w", err)
	}

	fmt.Printf("Found %d monitors and %d workspaces\n", len(monitors), len(workspaces))

	// Determine arrangement based on monitor count
	switch len(monitors) {
	case 1:
		// Single monitor - no rearrangement needed
		aerospace.SwitchWorkspace("A")
		return nil

	case 2:
		return rearrangeTwoMonitors(workspaces, monitors)

	case 3:
		return rearrangeThreeMonitors(workspaces, monitors)

	default:
		return fmt.Errorf("unsupported monitor configuration: %d monitors", len(monitors))
	}
}

// rearrangeTwoMonitors handles the 2-monitor setup:
// Built-in monitor gets A-L workspaces, the other gets L1-L5 and R1-R5
func rearrangeTwoMonitors(workspaces []aerospace.Workspace, monitors []aerospace.Monitor) error {
	// Find the built-in monitor
	var builtInID int
	var externalID int

	for _, mon := range monitors {
		if strings.Contains(mon.Name, "Built-in") {
			builtInID = mon.ID
		} else {
			externalID = mon.ID
		}
	}

	if builtInID == 0 || externalID == 0 {
		return fmt.Errorf("could not identify built-in and external monitors")
	}

	fmt.Printf("Built-in monitor: %d, External monitor: %d\n", builtInID, externalID)

	// Move workspaces to appropriate monitors
	for _, ws := range workspaces {
		var targetMonitor int

		// A-L go to built-in, L1-L5 and R1-R5 go to external
		if ws.Name >= "A" && ws.Name <= "L" {
			targetMonitor = builtInID
		} else if (ws.Name >= "L1" && ws.Name <= "L5") || (ws.Name >= "R1" && ws.Name <= "R5") {
			targetMonitor = externalID
		} else {
			// Skip unexpected workspace names
			continue
		}

		// Only move if not already on the correct monitor
		if ws.MonitorID != targetMonitor {
			fmt.Printf("Moving workspace %s to monitor %d\n", ws.Name, targetMonitor)
			if err := aerospace.MoveWorkspaceToMonitor(ws.Name, targetMonitor); err != nil {
				return fmt.Errorf("failed to move workspace %s: %w", ws.Name, err)
			}
		}
	}

	aerospace.SwitchWorkspace("A")
	aerospace.SwitchWorkspace("L1")

	return nil
}

// rearrangeThreeMonitors handles the 3-monitor setup:
// Built-in gets A-L, left external (smaller ID) gets L1-L5, right external gets R1-R5
func rearrangeThreeMonitors(workspaces []aerospace.Workspace, monitors []aerospace.Monitor) error {
	// Find built-in and external monitors
	var builtInID int
	var externalIDs []int

	for _, mon := range monitors {
		if strings.Contains(mon.Name, "Built-in") {
			builtInID = mon.ID
		} else {
			externalIDs = append(externalIDs, mon.ID)
		}
	}

	if builtInID == 0 || len(externalIDs) != 2 {
		return fmt.Errorf("could not identify monitors properly for 3-monitor setup")
	}

	// The external monitors are already sorted by ID (left to right)
	// Smaller ID is left, larger ID is right
	leftExternalID := externalIDs[0]
	rightExternalID := externalIDs[1]
	if externalIDs[1] < externalIDs[0] {
		leftExternalID = externalIDs[1]
		rightExternalID = externalIDs[0]
	}

	fmt.Printf("Built-in: %d, Left external: %d, Right external: %d\n", builtInID, leftExternalID, rightExternalID)

	// Move workspaces to appropriate monitors
	for _, ws := range workspaces {
		var targetMonitor int

		// A-L go to built-in, L1-L5 go to left external, R1-R5 go to right external
		if ws.Name >= "A" && ws.Name <= "L" {
			targetMonitor = builtInID
		} else if ws.Name >= "L1" && ws.Name <= "L5" {
			targetMonitor = leftExternalID
		} else if ws.Name >= "R1" && ws.Name <= "R5" {
			targetMonitor = rightExternalID
		} else {
			// Skip unexpected workspace names
			continue
		}

		// Only move if not already on the correct monitor
		if ws.MonitorID != targetMonitor {
			fmt.Printf("Moving workspace %s to monitor %d\n", ws.Name, targetMonitor)
			if err := aerospace.MoveWorkspaceToMonitor(ws.Name, targetMonitor); err != nil {
				return fmt.Errorf("failed to move workspace %s: %w", ws.Name, err)
			}
		}
	}

	aerospace.SwitchWorkspace("A")
	aerospace.SwitchWorkspace("L1")
	aerospace.SwitchWorkspace("R1")

	return nil
}
