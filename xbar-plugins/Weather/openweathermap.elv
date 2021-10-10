#!/usr/local/bin/elvish
# <xbar.title>Openweathermap</xbar.title>
# <xbar.author>Am Laher</xbar.author>
# <xbar.author.github>laher</xbar.author.github>
# <xbar.desc>Show the weather for a given location</xbar.desc>
# <xbar.dependencies>elvish,lunchbox</xbar.dependencies>
# <xbar.os>linux,windows,darwin</xbar.os>

use github.com/laher/lunchbox/elvish/util/dotenv
use github.com/laher/lunchbox/elvish/lunchbox
use str

dotenv:load "env/openweathermap.env"

var url = "http://api.openweathermap.org/data/2.5/weather?q="$E:CITY"&units="$E:UNITS"&appid="$E:API_KEY

try {
  #echo $url
  var resp = [(lunchbox:bin http GET $url)]
  var body = $resp[0]
  #echo $body

  var desc = (echo $body | lunchbox:bin jq -raw -query ".weather[0].main")
  # echo $desc
  var icon = ":question:"
  if (str:has-prefix $desc 'Cloud') {
    icon = ":cloud:"
  } elif (str:has-prefix $desc 'Sun') {
    icon = ":sun:"
  } elif (str:has-prefix $desc 'Rain') {
    icon = ":umbrella:"
  } elif (str:has-prefix $desc 'Cyclone') {
    icon = ":cyclone:"
  } elif (str:has-prefix $desc 'Snow') {
    icon = ":snow:"
  }

  echo $icon" "$desc
  echo "---"
  echo "Temp: "(echo $body | lunchbox:bin jq -query ".main.temp")
  echo "Humidity: "(echo $body | lunchbox:bin jq -query ".main.humidity")
  echo "Pressure: "(echo $body | lunchbox:bin jq -query ".main.pressure")
  echo "Visiblity: "(echo $body | lunchbox:bin jq -query ".main.visibility")

} except e {
  echo ':cry:'
  echo '---'
  echo $e
}

