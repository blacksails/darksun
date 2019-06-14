package vscode

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

const (
	defaultDark = "Default Dark+"
	defaultSun  = "Default Light+"
)

// New creates a new vscode module
func New(cfg Config) *Module {
	m := &Module{
		dark: defaultDark,
		sun:  defaultSun,
	}
	if cfg.Dark != "" {
		m.dark = cfg.Dark
	}
	if cfg.Sun != "" {
		m.sun = cfg.Sun
	}
	return m
}

// Module is the vscode module
type Module struct {
	dark string
	sun  string
}

// Config is the vscode module configuration
type Config struct {
	Enabled bool
	Dark    string
	Sun     string
}

// Dark implements darksun.Module
func (m *Module) Dark() error {
	return m.readAndWriteSettings(true)
}

// Sun implements darksun.Module
func (m *Module) Sun() error {
	return m.readAndWriteSettings(false)
}

func (m *Module) readAndWriteSettings(dark bool) error {
	settings, err := m.readSettings()
	if err != nil {
		return err
	}
	return m.writeSettings(settings, dark)
}

func (m *Module) readSettings() (map[string]interface{}, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	vscodeCfg := path.Join(home, "Library/Application Support/Code/User/settings.json")

	if _, err := os.Stat(vscodeCfg); err == nil {
		f, err := os.Open(vscodeCfg)
		if err != nil {
			return nil, err
		}
		var cfg map[string]interface{}
		err = json.NewDecoder(f).Decode(&cfg)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	} else if os.IsNotExist(err) {
		return map[string]interface{}{}, nil
	} else {
		return nil, err
	}
}

func (m *Module) writeSettings(settings map[string]interface{}, dark bool) error {
	settings["workbench.colorTheme"] = m.sun
	if dark {
		settings["workbench.colorTheme"] = m.dark
	}
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(settings); err != nil {
		return err
	}
	vscodeCfg := path.Join(home, "Library/Application Support/Code/User/settings.json")
	return ioutil.WriteFile(vscodeCfg, b.Bytes(), 0644)
}
