# lunchbox - batteries-included binary tools for Go shells (e.g. elvish)

 * general-purpose utilities which work similarly across different platforms (mainly Linux, MacOS, Windows)
 * Invoke them with subcommands, similar to busybox (`lunchbox http`, etc)
 * The ability to bundle lunchbox tools within a go binary
 * Include some elvish wrapper libraries 
 * Provides support for xbar/lunchbar plugins

## Goals

Initial goal is to build an xbar clone which is comfortable to use cross-platformally. Elvish script writers should be able to download stuff, manipulate files, run stuff, without too much custom code for cross-platform support.

 * Go-based extended featureset with buysbox-style subcommands
  * [x] Create the binary and first elvish wrapper scripts
  * [x] Get date/time. 
    * [x] timezone support
  * [ ] File operations (mv/cp/rm, etc)
  * [ ] HTTP
    * [ ] Make API calls
    * [ ] download stuff
  * [ ] compression/decompression tooling
 * [ ] Cross-platform scripting utilities
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
 * [ ] World clock
 * [ ] RSS
 * [ ] Some other API stuff

### xbar/lunchbar Usage

e.g. to use the docker-compose plugin ...

```
#!/usr/local/bin/lunchbox

use github.com/laher/lunchbox/elvish/util/dotenv
use github.com/laher/lunchbox/elvish/xbar/plugin/dev/docker-compose

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
