package macos

import (
	"fmt"
	"os/exec"
)

// New creates a new macos module
func New() *Module {
	return &Module{}
}

// Module is a darksun.Module which can change the os appearance. First time
// this is used you will see a prompt asking access to System Events.
type Module struct{}

// Config is module configuration for macos
type Config struct {
	Enabled bool
}

const macOSSetDarkMode = `
tell application "System Events"
	tell appearance preferences
		set dark mode to %s
	end tell
end tell
`

// Dark implements darksun.Module
func (m *Module) Dark() error {
	if isDark() {
		return nil
	}
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(macOSSetDarkMode, "true"))
	return cmd.Run()
}

// Light implements darksun.Module
func (m *Module) Light() error {
	if !isDark() {
		return nil
	}
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(macOSSetDarkMode, "false"))
	return cmd.Run()
}

func isDark() bool {
	cmd := exec.Command("defaults", "read", "-g", "AppleInterfaceStyle")
	return cmd.Run() == nil
}
