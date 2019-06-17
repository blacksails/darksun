package main

import (
	"github.com/blacksails/darksun"
	"github.com/blacksails/darksun/iterm2"
	"github.com/blacksails/darksun/macos"
	"github.com/blacksails/darksun/vim"
	"github.com/blacksails/darksun/vscode"
)

type config struct {
	Dark    bool
	Modules modules
}

type modules struct {
	ITerm2 iterm2.Config
	MacOS  macos.Config
	VSCode vscode.Config
	Vim    vim.Config
}

// GetModules reads the configuration and returns a list of the enabled
// modules.
func GetModules() ([]darksun.Module, error) {
	var config config
	if err := cfg.UnmarshalExact(&config); err != nil {
		return nil, err
	}

	modCfg := config.Modules
	var modules []darksun.Module

	if modCfg.MacOS.Enabled {
		modules = append(modules, macos.New())
	}
	if modCfg.ITerm2.Enabled {
		modules = append(modules, iterm2.New(modCfg.ITerm2))
	}
	if modCfg.VSCode.Enabled {
		modules = append(modules, vscode.New(modCfg.VSCode))
	}
	if modCfg.Vim.Enabled {
		modules = append(modules, vim.New(modCfg.Vim))
	}

	return modules, nil
}
