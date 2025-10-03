package hyprmove

import (
	"fmt"

	"github.com/Xkonti/aeromanager/internal/aerospace"
	"github.com/Xkonti/aeromanager/internal/workspacemap"
)

// Execute performs intelligent window movement based on cursor position
// If workspaceNum is -1, moves the window to the visible workspace on the monitor with the mouse
// Otherwise, moves the window to the workspace that corresponds to the given number
func Execute(workspaceNum int) error {
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

	var targetWorkspace string

	if workspaceNum == -1 {
		// Move to the visible workspace on the monitor with the mouse
		targetWorkspace = findVisibleWorkspaceOnMonitor(mouseMonitorID, workspaces)
		if targetWorkspace == "" {
			return fmt.Errorf("no visible workspace found on monitor %d", mouseMonitorID)
		}
		fmt.Printf("Moving focused window to visible workspace %s on monitor %d\n", targetWorkspace, mouseMonitorID)
	} else {
		// Validate workspace number (1-5 or 6-0, where 0 is treated as 10)
		if workspaceNum < 0 || workspaceNum > 10 {
			return fmt.Errorf("invalid workspace number: %d (must be 1-5 or 6-0)", workspaceNum)
		}

		// Determine which workspace to move the window to based on monitor count and cursor position
		targetWorkspace = workspacemap.MapWorkspaceNumber(workspaceNum, mouseMonitorID, monitors)

		if len(monitors) > 3 {
			return fmt.Errorf("unsupported monitor configuration: %d monitors", len(monitors))
		}

		// Validate that the target workspace exists
		if !workspaceExists(targetWorkspace, workspaces) {
			return fmt.Errorf("workspace %s does not exist", targetWorkspace)
		}

		fmt.Printf("Moving focused window to workspace %s on monitor %d\n", targetWorkspace, mouseMonitorID)
	}

	// Move the focused window to the target workspace
	// We don't use focus-follows-window because the user might want to keep working
	// on the current workspace after moving a window
	return aerospace.MoveNodeToWorkspace(targetWorkspace, false)
}

// findVisibleWorkspaceOnMonitor finds the visible workspace on a specific monitor
func findVisibleWorkspaceOnMonitor(monitorID int, workspaces []aerospace.Workspace) string {
	for _, ws := range workspaces {
		if ws.MonitorID == monitorID && ws.IsVisible {
			return ws.Name
		}
	}
	return ""
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