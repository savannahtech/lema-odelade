# accessment
Golang

Docker is used to containerize the application

The root of the project contains docker-compose.yml and Dockerfile

To start the app run 

Install docker on your machine

cd into the root directory and run 

`docker build .`

`docker compose up`

Please you have to provide the following env variables in the docker-compose.yml file
`GITACCESSTOKEN`
`OWNER`
`REPONAME`

these are currently empty strings, they are however required

Please use your personal Git Access token for `GITACCESSTOKEN`

## Endpoints 
Method is GET for bot endpoints
1.
http://localhost:8085/api/services/repo

2.
http://localhost:8085/api/services/commit

## NOTE
Unit Test file can be found in the rest_test.go in the service package

Rate limiting was implemented in the cron package



