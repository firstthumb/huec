# HUE CLI

<img src="https://user-images.githubusercontent.com/66023/123544560-8bf53c80-d75c-11eb-9d04-0f19fd729a77.png" width="400">

## Description
Control your Hue Lights within your terminal.
- Support for Linux, MacOS, and Windows
- Download easily with homebrew
- Turn on and off the ligths
- See the latest status of the lights

## Installation
Get the latest version <a href="https://github.com/firstthumb/huec/releases/latest">here</a>.

```
$ brew install firstthumb/homebrew-tap/huec
```


## Usage

After installation you need to authorize the client for controlling hue bridge.

```
$ huec init
```

Get all lights attached to your bridge.
```
$ huec ligths 
```

Turn on
```
$ huec ligths on 1 
```

Turn off
```
$ huec ligths off 1 
```

## Commands

```
$ huec help
huec controls a Philips Hue

Usage:
  huec [command]

Available Commands:
  help        Help about any command
  init        Initializes huec
  lights      Manage Hue light bulbs
  version     Prints the version

Flags:
  -h, --help   help for huec
```

## Missing features

There are still lots of features missing and waiting to be implemented.
* Set color of the lights
* Get color of the lights
* Set scene
