package hyprworkspace

import (
	"fmt"

	"github.com/Xkonti/aeromanager/internal/aerospace"
	"github.com/Xkonti/aeromanager/internal/workspacemap"
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
	targetWorkspace := workspacemap.MapWorkspaceNumber(workspaceNum, mouseMonitorID, monitors)

	if len(monitors) > 3 {
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


// workspaceExists checks if a workspace with the given name exists
func workspaceExists(name string, workspaces []aerospace.Workspace) bool {
	for _, ws := range workspaces {
		if ws.Name == name {
			return true
		}
	}
	return false
}
