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

// MoveNodeToWorkspace moves the focused window to a specific workspace
// If focusFollows is true, the focus will follow the window to its new workspace
func MoveNodeToWorkspace(workspaceName string, focusFollows bool) error {
	args := []string{"move-node-to-workspace"}
	if focusFollows {
		args = append(args, "--focus-follows-window")
	}
	args = append(args, workspaceName)

	cmd := exec.Command("aerospace", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to move window to workspace %s: %w (output: %s)", workspaceName, err, string(output))
	}

	// If there's any output, something likely went wrong
	if len(output) > 0 {
		return fmt.Errorf("unexpected output while moving window to workspace %s: %s", workspaceName, string(output))
	}

	return nil
}
