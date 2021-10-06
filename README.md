# elvish support for xbar plugins

These scripts will provide elvish-based utilities for use in xbar (and crossbar) plugins

## Goals

 * [x] Some basic cross-platform support (mainly for crossbar. eventually for xbar, hopefully).
 * [x] Support for external (dotenv-style) config files
 * [ ] Easy to reuse and compose together (why not have multiple docker-composes stacked into a single top-level menu?)
 * [ ] Batteries included
 * [ ] Examples for using these

## Plugins

 * [x] Docker-compose
   * [ ] Stackable
 * [ ] Managing processes
 * [ ] World clock
 * [ ] RSS
 * [ ] Some other API stuff

## Usage

e.g. to use the docker-compose plugin ...

```
#!/usr/local/bin/elvish

use github.com/laher/xbar-elvish/util/dotenv
use github.com/laher/xbar-elvish/dev/docker-compose

# You can use the dotenv module to manage config separately. If you want to...
dotenv:load ../config/dc.env
var WD = $E:WD

var scriptName = (src)[name]

if (> (count $args) 0) {
  cd $WD
  # the plugin's click handlers will invoke this script to perform docker-compose operations ...
  docker-compose $@args
} else {
  docker-compose:plugin $WD $scriptName
}
```


## About Crossbar

Crossbar is an experimental xbar clone which is cross-platform and bundles elvish support. I haven't shared it yet ...
