use github.com/laher/lunchbox/elvish/lunchbox

fn get [url]{
  lunchbox:bin http GET $url
}

fn do [method url @rest]{
  lunchbox:bin http $method $url $@rest
}

fn is_error [resp]{
  put $true
}
