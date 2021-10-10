use github.com/laher/lunchbox/elvish/lunchbox
use str

fn plugin [locations format]{
  lunchbox:bin date -format $format" :globe_with_meridians:"
  echo "---"
  str:split , $locations | each [loc]{
   lunchbox:bin date -location $loc -format $format" - "$loc" | shell=ls"
  }
}
