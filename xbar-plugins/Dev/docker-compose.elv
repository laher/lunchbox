#!/usr/local/bin/elvish
# <xbar.title>docker-compose</xbar.title>
# <xbar.author>Am Laher</xbar.author>
# <xbar.author.github>laher</xbar.author.github>
# <xbar.desc>Show status for a given docker-compose</xbar.desc>
# <xbar.dependencies>elvish,lunchbox,docker-compose</xbar.dependencies>
# <xbar.os>linux,windows,darwin</xbar.os>

use path
use str
use github.com/laher/lunchbox/elvish/lunchbox

fn dcplugin [WD scriptName]{
  cd $WD
  var ups = 0
  var downs = 0
  var @up = (docker-compose ps --filter status=running --services 2>/dev/null | each [x]{ echo '-- '$x" :arrow_down: | terminal=true shell='"$scriptName"' param1=stop param2="$x })
  var @down = (docker-compose ps --filter status=stopped --services 2>/dev/null | each [x]{ echo '-- '$x" :arrow_up: | terminal=true shell='"$scriptName"' param1=stop param2="$x })
  var counts = $ups
  var name = (path:base $WD)
  echo ":whale: "$name" "(count $up)"/"(count $down)
  echo "---"
  echo "Running services"
  echo (str:join "\n" $up)
  echo "Stopped services"
  echo (str:join "\n" $down)
}

# load environment
# at the minimum, just specify WD (working directory for docker-compose file)
# may need to also adjust PATH
lunchbox:dotenv env/dc.env

if (not-eq $E:PATH_EXTRA "") {
  set-env PATH $E:PATH":"$E:PATH_EXTRA
}
cd $E:WD

if (> (count $args) 0) {
  docker-compose $@args
} else {
  #scriptName is used for self-referencing this script
  var scriptName = (src)[name]
  dcplugin $E:WD $scriptName
}

