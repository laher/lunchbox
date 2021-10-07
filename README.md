# valinor.elv - cross-platform support for elvish

 * general-purpose utilities which work similarly across different platforms (mainly Linux, MacOS, Windows)
 * Some go binary/library plus some elvish wrapper libraries 
 * The ability to bundle valinor along with an elvish interpreter
 * Some support for xbar/crossbar plugins

## Goals

Initial goal is to build an xbar clone which is comfortable to use cross-platformly. plugins should be able to download stuff, manipulate files, run stuff, without too much custom code.

 * [ ] Cross-platform core utilities (a busybox-ish Go app)
  * [ ] find HOME dir
  * [ ] File operations (mv/cp, etc)
  * [ ] HTTP
    * [ ] Make API calls
    * [ ] download stuff
  * [ ] Get date/time. 
    * [ ] timezone support
  * [ ] compression/decompression tooling
 * [ ] General purpose scripting utilities
   * [x] Support for external (dotenv-style) config files
   * [ ] Open a file, directory
 * [ ] Bundle valinor as part of an elvish interpreter
   * [ ] A sample here
   * [ ] bundling with crossbar
 * [ ] xbar/crossbar plugin support
   * [ ] Easy to reuse and compose together (why not have multiple docker-composes stacked into a single top-level menu?)
   * [ ] Batteries included (wrappers for valinor features)
   * [ ] Examples for using these

## Xbar/Crossbar Plugins

 * [x] Docker-compose
   * [ ] Stackable
 * [ ] Managing processes
 * [ ] World clock
 * [ ] RSS
 * [ ] Some other API stuff

### xbar/crossbar Usage

e.g. to use the docker-compose plugin ...

```
#!/usr/local/bin/valinor

use github.com/laher/valinor.elv/util/dotenv
use github.com/laher/valinor.elv/plugin/dev/docker-compose

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

Crossbar is an experimental xbar clone which is cross-platform and bundles elvish (and valinor) support. I haven't shared it yet ...
