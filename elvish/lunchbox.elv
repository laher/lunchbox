use platform
use str

fn do [@args]{
  if (!=s $E:LUNCHBOX_BIN "") {
    # use an alternative lunchbox provider, if set
    # note that the standard elvish distribution doesn't allow bare variable invocation (needs a slash).
    # Best to set full path
    $E:LUNCHBOX_BIN $@args
  } else {
    lunchbox $@args
  }
}

fn http [@args]{
  do http $@args
}

fn jq [@args]{
  do jq $@args
}

fn date [@args]{
  do date $@args
}

# this is really primitive. I'm considering dropping it
fn dotenv [@args]{
  do dotenv $@args | each [x]{
    var @parts = (str:split '=' $x)
    set-env $parts[0] $parts[1]
  }
}
