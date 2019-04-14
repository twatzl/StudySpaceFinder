FROM node:lts-alpine

ENV PORT=8080

# install simple json server for serving static content
RUN npm install -g json-server

# make the 'app' folder the current working directory
WORKDIR /app

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

EXPOSE ${PORT} 
CMD json-server --host 0.0.0.0 -p $PORT db.json
