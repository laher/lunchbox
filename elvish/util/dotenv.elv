

# TODO: this is very primitive ... 
# I think dotenv files should be parsed via luncbhox 
fn load [f]{ 
  use file
  use str
  var fh = (file:open $f)
  each [x]{ 
    var @parts = (str:split '=' $x)
    #echo $parts
    set-env $parts[0] $parts[1] 
  } < $fh
  file:close $fh
}
