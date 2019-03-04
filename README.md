# StudySpaceFinder

A simple app to find a study space for you needs.

## Development

### Prerequisites

- Docker (+ docker-compose)
- go (https://golang.org/doc/install)
- npm (https://www.npmjs.com/get-npm)
- vue-cli (https://cli.vuejs.org/guide/installation.html)
- buffalo (https://gobuffalo.io/en/docs/installation)

### Frontend

The Frontend is written using **vue.js** with **typescript**.

In order to start the frontend run

```
npm install
npm run serve
```

in the frontend directory.

### Backend

Backend is written using **golang** and **buffalo**.

To run the backend in development mode just run:

```
buffalo db create -a -d
buffalo db migrate -d
buffalo dev
```

This will:
1. create all necessary databases
2. create the tables in the databases
3. start the buffalo app in dev mode


### Database

For running the app you also need a database. We use cockroachdb. The easiest way to set it up is using the docker-compose.yml file in infrastructure/cockroachdb.

Simply change the path accordingly and make sure you manually create the path before starting. 

Then start the database using:

```
docker-compose up -d
```


