#!/usr/local/bin/elvish

use path
use github.com/laher/lunchbox/elvish/util/dotenv

fn dcplugin [WD scriptName]{
  cd $WD
  var name = (path:base $WD)
  echo ":whale: "$name
  echo "---"
  echo "Running services"
  docker-compose ps --filter status=running --services 2>/dev/null | each [x]{ echo '-- '$x" :arrow_down: | terminal=true shell='"$scriptName"' param1=stop param2="$x }
  echo "Stopped services"
  docker-compose ps --filter status=stopped --services 2>/dev/null | each [x]{ echo '-- '$x" :arrow_up: | terminal=true shell='"$scriptName"' param1=stop param2="$x }
}

# load environment 
# at the minimum, just specify WD (working directory for docker-compose file)
# may need to also adjust PATH
dotenv:load env/dc.env
cd $E:WD

if (> (count $args) 0) {
  docker-compose $@args
} else {
  #scriptName is used for self-referencing this script
  var scriptName = (src)[name]
  dcplugin $E:WD $scriptName
}

