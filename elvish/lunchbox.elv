use platform

fn do [@args]{
  if (!=s $E:LUNCHBOX_BIN "") {
    # use an alternative lunchbox provider, if set
    # note that the standard elvish distribution doesn't allow this style of variable invocation (needs a slash). 
    # TODO is there any equivalent safety check which should be applied here?
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
