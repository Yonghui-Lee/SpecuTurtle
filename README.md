## SpecuTurtle

SpecuTurtle is a  forum for Situation-Puzzle. Please visit [https://specuturtle.xyz](https://specuturtle.xyz) for more details. 

## How to deployment

1. A VPS
2. Install Nginx `sudo apt install nginx -y`
3. Install Postgresql `sudo apt install postgresql -y`, after create the database, you need import the database schema
4. Deploy the api server and web
5. Then you can visit you website after restart nginx

## Features

1. REST API back-end written in Golang
2. React-based frontend
3. PostgreSQL, one of the best open source, flexible database
4. JSON Web Tokens (JWT) are used for user authentication in the API
5. Markdown supported topic and comment
6. Model tested


## Built With

1. go version go1.15 darwin/amd64
2. postgres (PostgreSQL) 12.3
3. react ^16.13.1

## Structure

1. `./` is back-end service, we followed [golang-standards project-layout](https://github.com/golang-standards/project-layout).
2. `./app` is front-end service, contains React, Parcel and etc.
2. `./deploy` contains example of deploy, nginx and systemd.



## Getting Started

### Backend

1. `cd ./internal`, copy `config/config.example` to `config/config.yaml`. Replace config with yours.
2. Prepare and start database, the database schema under `./internal/models/schema.sql`
3. `cd ./ && go build && ./satellity` to start Golang server

### Frontend

1. Copy `env.example` to `.env`
   
    ```
    SITE_NAME=your site name
    ```
2. run `yarn install`, then `yarn start`

## License

![https://opensource.org/licenses/MIT](https://img.shields.io/github/license/mashape/apistatus.svg)
