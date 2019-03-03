# StudySpaceFinder

A simple app to find a study space for you needs.

## Development

### Frontend

The Frontend is written using **vue.js** with **typescript**.

In order to start the frontend run

```
npm run serve
```

in the frontend directory.

### Backend

Backend is written using **golang** and **buffalo**.



### Database

For running the app you also need a database. We use cockroachdb. The easiest way to set it up is using the docker-compose.yml file in infrastructure/cockroachdb.

Simply change the path accordingly and make sure you manually create the path before starting. 

Then start the database using:

```
docker-compose up -d
```


