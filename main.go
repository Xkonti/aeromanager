package main

import (
	"fmt"
	"os"

	"github.com/Xkonti/aeromanager/internal/rearrange"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aeromanager <command>")
		fmt.Println("Commands:")
		fmt.Println("  rearrange - Rearrange workspaces based on monitor setup")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "rearrange":
		if err := rearrange.Execute(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
