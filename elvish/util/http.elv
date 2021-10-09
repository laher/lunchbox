use github.com/laher/lunchbox/elvish/lunchbox

fn get [url]{
  lunchbox:bin http -method GET $url
}

fn is_error [resp]{
  put $true
}
