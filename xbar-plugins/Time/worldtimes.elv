#!/usr/local/bin/elvish
# <xbar.title>World times</xbar.title>
# <xbar.author>Am Laher</xbar.author>
# <xbar.author.github>laher</xbar.author.github>
# <xbar.desc>Show times across multiple locations</xbar.desc>
# <xbar.dependencies>elvish,lunchbox</xbar.dependencies>
# <xbar.os>linux,windows,darwin</xbar.os>

use str
use github.com/laher/lunchbox/elvish/lunchbox

# load some environment variables
# example:
# LOCATIONS=UTC,Europe/London,America/Los_Angeles
# FORMAT="%a %d %b %H:%M" 
lunchbox:dotenv env/worldtimes.env
var locations = $E:LOCATIONS
if (eq $locations "") {
  locations = "UTC,Europe/London,America/Los_Angeles" 
}
var format = $E:FORMAT
if (eq $format "") {
  format = "%a %d %b %H:%M" 
}
#echo $format
loc = (lunchbox:date -format "%a %H:%M")
echo $loc" :globe_with_meridians:"
echo "---"
str:split , $locations | each [loc]{
 var @parts = (str:split "/" $loc)
 var idx = (- (count $parts) 1)
 var shortloc = $parts[$idx]
 lunchbox:date -location $loc -format $format" - "$shortloc
 echo "-- "$loc
}
