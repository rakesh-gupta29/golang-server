## Golang Web Server

Blog management project for learning how to build Golang projects

> _This is still WORK IN PROGRESS_

> _Please do not use it unless your backend is client agnostic: when it does not care about the client. Otherwise, go with NextJS or NuxtJS. They are better options in that case. This backend is useful only beyond CRUD._

## Prerequisites

To clone and run project, you must have Golang installed and set up on your local machine.

## Folder Structure:

- [bin](https://www.github.com/rakesh-gupta29/golang-server/tree/main/bin): directory will contains compiled application binaries, ready for deployment to a production server.
- [cmd/api](https://www.github.com/rakesh-gupta29/golang-server/tree/main/cmd/api): directory will contain the application-specific code for our API
- [internal](https://www.github.com/rakesh-gupta29/golang-server/tree/main/internal): contains various ancillary packages used by our API.
- [migrations](https://www.github.com/rakesh-gupta29/golang-server/tree/main/migrations): migrations directory will contain the SQL migration files for our database.
- [remote](https://www.github.com/rakesh-gupta29/golang-server/tree/main/remote): The remote directory : will contain the configuration files and setup scripts for our production server.
- [go.mod](https://www.github.com/rakesh-gupta29/golang-server/tree/main/go.mod): The go.mod file will declare our project dependencies, versions and module path.
- [Makefile](https://www.github.com/rakesh-gupta29/golang-server/tree/main/Makefile): Makefile will contain recipesfor automating common administrative tasks

## Features

- routing - using [httprouter](https://github.com/julienschmidt/httprouter) package; Making the
- static file serving: server static files using h[httprouter](https://github.com/julienschmidt/httprouter) package: serve all files under the [static](https://www.github.com/rakesh-gupta29/golang-server/tree/main/static) directory
- Utils for handling all errors with proper headers and messages.

## Installation

To clone and run the following project, run the following commands

```
 $ git clone https://www.github.com/rakesh-gupta29/golang-web-server.git
 $ cd golang-web-server
 $ go  mod tidy
 $ go run
```

Before we can run the project, ceate a `.env` file ar the root of the project and fill appropriate values for following properties.

```
port = runtime port for server (number)
```

## Deployment

To build the code for deployment, build the code for deployment using following command

```
go  build
```

## Author

[Rakesh Gupta](https://www.github.com/rakesh-gupta29)
