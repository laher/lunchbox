use github.com/laher/lunchbox/elvish/lunchbox
use str

fn plugin [locations layout]{
  # TODO LAYOUT unless layout is non-empty
  var LAYOUT = "05/04 03:02:01"
  lunchbox:bin date -layout $LAYOUT" :globe_with_meridians:"
  echo "---"
  str:split , $locations | each [loc]{
   lunchbox:bin date -location $loc -layout $LAYOUT" - "$loc" | shell=ls"
  }
}
