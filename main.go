package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Xkonti/aeromanager/internal/hyprworkspace"
	"github.com/Xkonti/aeromanager/internal/rearrange"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aeromanager <command> [args]")
		fmt.Println("Commands:")
		fmt.Println("  rearrange            - Rearrange workspaces based on monitor setup")
		fmt.Println("  hyprworkspace <num>  - Switch workspace based on cursor position (num: 1-5 or 6-0)")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "rearrange":
		if err := rearrange.Execute(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "hyprworkspace":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: hyprworkspace requires a workspace number (1-5 or 6-0)\n")
			os.Exit(1)
		}
		workspaceNum, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid workspace number: %s\n", os.Args[2])
			os.Exit(1)
		}
		if err := hyprworkspace.Execute(workspaceNum); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
