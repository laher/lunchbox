

fn plugin [WD scriptName]{
  cd $WD
  echo ":whale: docker-compose"
  echo "---"
  echo "Running services"
  docker-compose ps --filter status=running --services 2>/dev/null | each [x]{ echo '-- '$x" :arrow_down: | terminal=true shell='"$scriptName"' param1=stop param2="$x }
  echo "Stopped services"
  docker-compose ps --filter status=stopped --services 2>/dev/null | each [x]{ echo '-- '$x" :arrow_up: | terminal=true shell='"$scriptName"' param1=stop param2="$x }
}
