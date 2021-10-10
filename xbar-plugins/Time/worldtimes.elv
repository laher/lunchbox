#!/usr/local/bin/elvish

use github.com/laher/lunchbox/elvish/lunchbox
use str
use github.com/laher/lunchbox/elvish/util/dotenv

fn plugin [locations format]{
}

# load some environment variables
# example:
# LOCATIONS=
# FORMAT=
dotenv:load env/worldtimes.env
var locations = $E:LOCATIONS
var format = $E:FORMAT
lunchbox:bin date -format $format" :globe_with_meridians:"
echo "---"
str:split , $locations | each [loc]{
 lunchbox:bin date -location $loc -format $format" - "$loc" | shell=ls"
}
