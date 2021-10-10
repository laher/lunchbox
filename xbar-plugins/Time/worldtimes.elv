#!/usr/local/bin/elvish
# <xbar.title>World times</xbar.title>
# <xbar.author>Am Laher</xbar.author>
# <xbar.author.github>laher</xbar.author.github>
# <xbar.desc>Show times across multiple locations</xbar.desc>
# <xbar.dependencies>elvish,lunchbox</xbar.dependencies>
# <xbar.os>linux,windows,darwin</xbar.os>

use str
use github.com/laher/lunchbox/elvish/lunchbox
use github.com/laher/lunchbox/elvish/util/dotenv

# load some environment variables
# example:
# LOCATIONS=UTC,Europe/London,America/Los_Angeles
# FORMAT="%a %d %b %H:%M" 
dotenv:load env/worldtimes.env
var locations = $E:LOCATIONS
var format = $E:FORMAT
lunchbox:bin date -format $format" :globe_with_meridians:"
echo "---"
str:split , $locations | each [loc]{
 lunchbox:bin date -location $loc -format $format" - "$loc" | shell=ls"
}
