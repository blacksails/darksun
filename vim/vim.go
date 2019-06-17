package vim

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// New returns a new neovim module
func New(cfg Config) *Module {
	return &Module{
		nvim:       cfg.Neovim,
		file:       cfg.File,
		sunScheme:  cfg.Sun.ColorScheme,
		darkScheme: cfg.Dark.ColorScheme,
		sunBG:      cfg.Sun.Background,
		darkBG:     cfg.Dark.Background,
	}
}

// Module is a darksun.Module
type Module struct {
	nvim       bool
	file       string
	sunScheme  string
	darkScheme string
	sunBG      string
	darkBG     string
}

// Config is the configuration object seen in the config file
type Config struct {
	Enabled bool
	Neovim  bool
	File    string
	Dark    ConfigMode
	Sun     ConfigMode
}

// ConfigMode is the vim config for either dark or sun mode
type ConfigMode struct {
	Background  string
	ColorScheme string
}

// Dark implements darksun.Module
func (m *Module) Dark() error {
	return m.run(true)
}

// Sun implements darksun.Module
func (m *Module) Sun() error {
	return m.run(false)
}

func (m *Module) run(dark bool) error {
	err := m.setConfigFile(dark)
	if err != nil {
		return err
	}
	if m.nvim {
		return m.updateNvim(dark)
	}
	return m.updateVim(dark)
}

func (m *Module) setConfigFile(dark bool) error {
	fileLoc, err := homedir.Expand(m.file)
	if err != nil {
		return err
	}
	cfg, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		return err
	}

	lines := strings.Split(string(cfg), "\n")
	bgFound := false
	schemeFound := false
	for i, line := range lines {
		if strings.Contains(line, "set background=") {
			bgFound = true
			if dark && m.darkBG != "" {
				lines[i] = fmt.Sprintf("set background=%s", m.darkBG)
			}
			if !dark && m.sunBG != "" {
				lines[i] = fmt.Sprintf("set background=%s", m.sunBG)
			}
		}
		if strings.Contains(line, "colorscheme") {
			schemeFound = true
			if dark && m.darkScheme != "" {
				lines[i] = fmt.Sprintf("colorscheme %s", m.darkScheme)
			}
			if !dark && m.sunScheme != "" {
				lines[i] = fmt.Sprintf("colorscheme %s", m.sunScheme)
			}
		}
	}

	if !bgFound {
		if dark && m.darkBG != "" {
			lines = append(lines, fmt.Sprintf("set background=%s", m.darkBG))
		}
		if !dark && m.sunBG != "" {
			lines = append(lines, fmt.Sprintf("set background=%s", m.sunBG))
		}
	}
	if !schemeFound {
		if dark && m.darkScheme != "" {
			lines = append(lines, fmt.Sprintf("colorscheme %s", m.darkScheme))
		}
		if !dark && m.sunScheme != "" {
			lines = append(lines, fmt.Sprintf("colorscheme %s", m.sunScheme))
		}
	}

	newCfg := strings.Join(lines, "\n")

	return ioutil.WriteFile(fileLoc, []byte(newCfg), 0644)
}

func (m *Module) updateNvim(dark bool) error {
	var vcmds []string
	if dark {
		if m.darkBG != "" {
			vcmds = append(vcmds, fmt.Sprintf("set background=%s", m.darkBG))
		}
		if m.darkScheme != "" {
			vcmds = append(vcmds, fmt.Sprintf("colorscheme %s", m.darkScheme))
		}
	} else {
		if m.sunBG != "" {
			vcmds = append(vcmds, fmt.Sprintf("set background=%s", m.sunBG))
		}
		if m.sunScheme != "" {
			vcmds = append(vcmds, fmt.Sprintf("colorscheme %s", m.sunScheme))
		}
	}
	if len(vcmds) == 0 {
		return nil
	}

	vcmd := strings.Join(vcmds, " | ")
	bcmd := fmt.Sprintf("nvr --serverlist | uniq | xargs -I '{}' nvr --servername '{}' -c '%s'", vcmd)
	c := exec.Command("bash", "-c", bcmd)
	return c.Run()
}

func (m *Module) updateVim(dark bool) error {
	panic("not implemented")
}
