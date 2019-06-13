package main

import (
	"github.com/blacksails/darksun"
	"github.com/blacksails/darksun/iterm2"
	"github.com/blacksails/darksun/macos"
)

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

	return modules, nil
}
