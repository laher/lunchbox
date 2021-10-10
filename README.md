# lunchbox - batteries-included tooling for bundling into Go shells (e.g. elvish)

*VERY EARLY PROJECT. CHANGING ALL THE TIME*

 * general-purpose utilities which work similarly across different platforms (mainly Linux, MacOS, Windows)
 * Invoke them with subcommands, similar to busybox (`lunchbox http`, etc)
 * The ability to bundle lunchbox tools within a go binary
 * Include some elvish wrapper libraries 
  * Some support for writing xbar/lunchbar plugins

## Goals

This project's goal is host some common tooling which can be bundled into another go binary (along with an interpreter of some sort), thereby simplifying installation of plugins across platforms.

My underlying goal is to support an xbar clone which is comfortable to use cross-platformally. Elvish script writers should be able to do common tasks - call APIs, download stuff, manipulate files, run stuff - without too much custom code for installation or for cross-platform support.


### Go-based subcommands

Go-based cli utils with buysbox-style subcommands ... 

| subcommand | status  | 3p tools  | description |
| ---------- | ------- | --------- | ----------- |
| date       | initial | [strftime](github.com/lestrrat-go/strftime) | Get date/time - timezone support; strftime format specification, e.g. %y/%m/%d %H:%M:%S |
| jq         | initial | [gojq](github.com/itchyny/gojq)      | JSON querying, similar to jq |
| http       | TODO    | httpie-go | HTTP (similar to httpie?) |
| [un]compress | TODO  |           | [un]compression tools |
| mv,cp,rm,... | TODO  |           | file manipulation with consistent flags across platforms | 
| pluginstall | TODO   |           | copy plugins (bundled into this binary with embedfs) into plugins dir |

### Elvish wrappers

 * Cross-platform scripting utilities
   * [x] Support for external (dotenv-style) config files
   * [ ] Open a file, directory (GUI)
   * [x] find HOME dir
 * Bundle elvish interpreter with lunchbox executable
   * [ ] A sample usage here
   * [ ] bundling from lunchbar too
 * xbar/lunchbar plugin support
   * [x] Batteries included (wrappers for lunchbox features)
   * [ ] Easy to reuse and compose together (why not have multiple docker-composes stacked into a single top-level menu?)
   * [ ] Examples for using these

## Xbar/Lunchbar Plugins

 * [x] Docker-compose
   * [ ] Stackable
 * [ ] Managing processes
 * [x] World clock
 * [ ] RSS
 * [ ] Some other API stuff

### xbar/lunchbar Usage

e.g. to use the docker-compose plugin ...

```
#!/usr/local/bin/lunchbox

use github.com/laher/lunchbox/elvish/util/dotenv
use github.com/laher/lunchbox/xbar-plugins/dev/docker-compose

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

## About Lunchbar

Lunchbar is an experimental xbar clone which is cross-platform and bundles elvish (and lunchbar) support. I haven't shared it yet ...
