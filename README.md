# golpje

Golpje is a pet project created to automatically search the piratebay for new episodes of tv shows and download them.  
The binary created has zero dependencies, it even has a bittorrent client integrated. 

## minimal effort usage

This setup will create a directory in your current working director called `shows` and automatically download download "new" episodes of the daily show and south park. (minimal season 20 for south park and at least from 2017 for the daily show)

 - Download a suitable release from the releases page and extract it
 - run `./golpje start &`
 - run `./golpje show add --name "daily show" --regexp "daily.*show" --active --minseason 2017`
 - run `./golpje show add --name "south park" --regexp "south.*park" --active --minseason 20 --seasonal`
 - ???
 - profit

 ## configuration

 golpje looks in several places for its configuration:

 - `./config.yml` (current directory)
 - `/etc/golpje/config.yml`
 - `$HOME/.golpje/config.yml`
 - environment variables (prefix with GOLPJE_ and uppercase the following values)

### Configurable:

Default values are shown between parenthesis.
 
 - shows_path (`"./shows/"`)
   - path to download shows into
 - download_path (`"/tmp/golpje/"`)
   - path used for storing the data while downloading (preferably on the same partition for quicker moving the data to the final destination)
 - database_file (`"golpje.db"`)
   - path to store the database (must be writable for the executing user)
 - port (`":3222"`)
   - address:port to listen on for the client
 - metrics_enabled (`true`)
   - enables an prometheus endpoint to scrape some metrics metrics
 - metrics_port (`":8080"`)
   - port to listen on for the metrics
 - metrics_path (`"/metrics"`)
   - path to listen on for the metrics
 - cli_address (`"localhost:3222"`)
   - this is used when issuing commands can be used to talk to a remote golpje instance (most useful when defined with an env var)
 - search_enabled (`true`)
   - set this to false to disable searching completely, can be useful when testing
 - search_interval (`"15m"`)
   - interval to search the piratebay, uses the golang [parseDuration](https://golang.org/pkg/time/#ParseDuration) to parse the duration
 