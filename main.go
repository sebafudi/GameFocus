package main

import (
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func main() {
	processName := "witcher3.exe"

	// Initialize a variable to keep track of the last known state
	lastState := "extend"

	for {
		cmd := exec.Command("tasklist", "/fi", "imagename eq "+processName)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // Hide the window
		output, err := cmd.CombinedOutput()
		currentState := "extend"
		if err == nil && strings.Contains(string(output), processName) {
			currentState = "internal"
		}

		if currentState != lastState {
			// Display mode has changed, perform the switch
			displaySwitchCmd := exec.Command("DisplaySwitch.exe", "/"+currentState)
			displaySwitchCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // Hide the window
			displaySwitchCmd.Run()
			lastState = currentState
		}

		// Sleep for a while before checking again
		time.Sleep(2 * time.Second)
	}
}
