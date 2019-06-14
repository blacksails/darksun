# Darksun 🌓

Developers work in diverse environments – Sometimes it's in a dark basement
other times at a sunny office. Darksun is a program which lets you toggle
dark/light mode across all your applications in one go.

## Installation

Currently the only way of installing this is with the go tool chain:

```
go get -u github.com/blacksails/darksun/cmd/darksun
```

## Configuration

Darksun looks for a config file in the following places:

```
/etc/darksun/config.yaml
$HOME/.config/darksun/config.yaml
$HOME/.darksun/config.yaml 
./config.yaml
```

The following is a sample configuration:

```yaml
modules:
  macos:
    enabled: true
  iterm2:
    enabled: true
    dark: ~/.config/iterm2/dark.json
    sun: ~/.config/iterm2/sun.json
  vscode:
    enabled: true
```

| Key | Default | Required when enabled | Description |
| --- | ------- | --------------------- | ----------- |
| modules.macos.enabled | false | | Enables macos module |
| modules.iterm2.enabled | false | | Enables iterm2 module |
| modules.iterm2.dark | | true | Path to iterm2 dark profile |
| modules.iterm2.sun | | true | Path to iterm2 sun profile |
| modules.iterm2.guid | | false | GUID used for iterm2 Darksun profile |
| modules.vscode.enabled | false | | Enables vscode module |
| modules.vscode.dark | Default Dark+ | false | Name of dark vscode workbench color theme |
| modules.vscode.dark | Default Sun+ | false | Name of sun vscode workbench color theme |

All modules are off by default and you will need to create the config file and
set the enabled field on the modules that you want. The following sections
contain details about how to configure each of the modules.

### iTerm2

Darksun relies on the DynamicProfile functionality of iTerm. Based on the
configured dark and sun profiles it generates a dynamic profile called Darksun,
this profile should be set as the default in iTerm.

In order to generate the dark.json and sun.json files it is easiest to create
two manual profiles in iTerm and configure them to your liking you can then
click "Other actions..." and select "Copy Profile as JSON". Paste the json for
each of the profiles to each of the files.

## Contributing

Any contribution is welcome 🙏

Ideas for contribution:
- option for creating a launchd/cron schedule
- vim module (might require a vim plugin)
- tmux module
- intellij module (might require an intellij plugin)
- A module for a tool that you use that supports dark/light mode
- Better documentation
- Improve code quality
- Make it easier to add new modules
