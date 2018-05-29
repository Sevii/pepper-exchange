

Welcome to Pepper Exchange this is a rough demo cryptocurrency exchange. 

### Directories ###

`cmd` contains the golang code for the backend api of the exchange
`scripts` contains bash used for testing
`vendor` contains golang dependencies
`web` contains the reactjs frontend

### Components ###

Golang executable `exchange`
Redis Server with no authentication configured
React frontend


### Setup ###
Start redis running locally with no authentication, ex on mac `brew install redis` => `redis-server`.
Redis running on `localhost:6379` 
Go to `cmd/exchange` and run the binary `./exchange` to start the golang server.
Go to `web` and run `npm install`, `npm start` to boot up the webserver


### Architecture ###
![Architecture Diagram](Architecture.png)
