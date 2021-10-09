use github.com/laher/lunchbox/elvish/lunchbox

fn query [q doc]{
  echo $doc | lunchbox:bin json -query $q
}
