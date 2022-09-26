# Incrowd code test

Done by Niels Sanchez van den Beuken.

## Instructions to start the application on localhost

- install docker and docker-compose
- create variables.env file inside env folder (right now there are no critical values but here we could save apikey and sensitive stuff needed for real environments) with following values:

```[env]
ENVIRONMENT="LOCAL"
VERSION = "1.0.0"
DBURL="mongodb"
SPORTNEWSCOLLECTIONNAME="news"
DATABASENAME="incrowd"
NEWSURL="https://www.wearehullcity.co.uk/api/incrowd/getnewlistinformation?count=5"
ARTICLEURL="https://www.wearehullcity.co.uk/api/incrowd/getnewsarticleinformation?id="
```

- run following command in main folder of the repo: docker-compose up --build
- In order to check mongoDB: http://localhost:8081/ no password / user is required for this project
- In order to run the test manually in console run: go test -v ./...
- Postman collection has been done, please see attached in the repository. In order to make it work you can:
  - News
  - News by ID

## An explanation of the choices taken and assumptions made during development

A repository pattern an a hexagonal arquitecture has been applied. It is a very simple microservice but with this structure it should be easy to scalate with other services and handlers.
Moreover, with the docker image and docker compose is easy to implement in different environments using K8 and AWS stuff. Multistage has been used in docker to have a more lightweighted image, this has been done as a good practice.

Code is mainly synchronous because with defined requirements there is no need to over engineer things. The only asynchronous routine is the go cron routine to pull new information from the feed URL, process it and save it to the mongo DB. If necessary it can be implemented communication between routines to check if all the information has been gathered correctly then allow the server to process api calls. This has not been implemented.

Logger has been implemented in a very simple way just calling it as an internal global var initialized in main.

Unit test coverage is a little bit low, because it has been covered mainly the logic of the microservice and avoided many repetitions in similar patterns/logic . Should be higher.

## Possible extensions or improvements to the service (focusing on scalability and deployment to production)

- Configure env variables for the different environments
- Add proper authentication middlewares for cybersecurity
- Add K8 configurations to be able to scalate easy the application if high traffic or requests
- More sensitive approach for cron go routines being able to communicate and block between go routines (server handler and cron routines)
