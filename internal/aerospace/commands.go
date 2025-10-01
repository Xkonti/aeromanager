package aerospace

import (
	"fmt"
	"os/exec"
)

// MoveWorkspaceToMonitor moves a workspace to a specific monitor
func MoveWorkspaceToMonitor(workspaceName string, monitorID int) error {
	cmd := exec.Command("aerospace", "move-workspace-to-monitor", "--workspace", workspaceName, fmt.Sprintf("%d", monitorID))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to move workspace %s to monitor %d: %w (output: %s)", workspaceName, monitorID, err, string(output))
	}

	// If there's any output, something likely went wrong
	if len(output) > 0 {
		return fmt.Errorf("unexpected output while moving workspace %s to monitor %d: %s", workspaceName, monitorID, string(output))
	}

	return nil
}
