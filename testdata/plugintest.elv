use github.com/laher/xbar-elvish/util/plugins

var lst = [a b c]
put $@lst | each [x]{ plugins:as-submenu $x } | each [x]{ plugins:as-submenu $x }
