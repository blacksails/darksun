package iterm2

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// New returns a new iterm2.Module
func New(config Config) *Module {
	return &Module{
		darkPath:  config.Dark,
		lightPath: config.Light,
		guid:      config.GUID,
	}
}

// Module is a darksun.Module which can change the appearance of iterm.
type Module struct {
	darkPath  string
	lightPath string
	guid      string
}

// Config is the fields avaiable in the configuration
type Config struct {
	Enabled bool
	Dark    string
	Light   string
	GUID    string
}

// Dark implements darksun.Module
func (m *Module) Dark() error {
	profile, err := m.readProfile(true)
	if err != nil {
		return err
	}
	return m.updateDynamicProfile(profile)
}

// Sun implements darksun.Module
func (m *Module) Sun() error {
	profile, err := m.readProfile(false)
	if err != nil {
		return err
	}
	return m.updateDynamicProfile(profile)
}

func (m *Module) readProfile(dark bool) (map[string]interface{}, error) {
	var (
		p   string
		err error
	)
	if dark {
		p, err = homedir.Expand(m.darkPath)
	} else {
		p, err = homedir.Expand(m.lightPath)
	}
	p, err = homedir.Expand(p)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	profile := map[string]interface{}{}
	err = json.NewDecoder(f).Decode(&profile)
	f.Close()
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (m *Module) updateDynamicProfile(profile map[string]interface{}) error {
	profile["Guid"] = m.guid
	profile["Name"] = "Darksun"
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	darksunPath := path.Join(home, "Library/Application Support/iTerm2/DynamicProfiles/Darksun.json")
	var b bytes.Buffer
	output := map[string]interface{}{
		"Profiles": []interface{}{
			profile,
		},
	}
	err = json.NewEncoder(&b).Encode(output)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(darksunPath, b.Bytes(), 0644)
}

func getGUID() (string, error) {
	guid := viper.GetString("iterm2.guid")
	if guid == "" {
		guid = newGUID()
		viper.Set("iterm2.guid", guid)
		err := viper.WriteConfig()
		if err != nil {
			return "", err
		}
	}
	return guid, nil
}

func newGUID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}
