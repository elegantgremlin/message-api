# Messages API

## Overview

This is a simple API for storing and retrieving Messages. A Message is valid if it is a non-empty string that is less than 100 characters

## Architecture

This API consists of a single service backed by a database.

### Service

The service is written in Go. It is separated into three separate modules that combine together to process and store the Messages. Each module is built to work independently and have minimal coupling to the other modules. This allows them to be switched out easily; for example, switching from being an API to consuming from a queue, or switching the database to another type of storage.

* `Server`: this module contains the functionality for accepting incoming data. It accepts the incoming data and transforms it into something that can be processed by the other modules of the service. The service is currently setup to expose a simple REST API.
* `Service`: this module contains the business logic for the service. All data mutations and calculations happen in this layer. 
* `Database`: this module contains the functionality for storing and retrieving the data. It is responsible for transforming the data into the format necessary for it to be stored, and for transforming the data back into a format usable by the other modules in the service.

## Running the API

### Starting the Services

All necessary services are contained in the Docker Compose environment in the `compose/` directory.

```
cd compose/
docker-compose up
```

This will build and start the service, create the database, and create the necessary table in the database.

The API service is setup with live reloading so any changes made to the source code will cause the API service to re-compile automatically.

## Specification

A full OpenAPI specification for the service API can be found [here](v1-spec.yaml).

## Tests

### Unit Tests

All unit tests can be run using Go's builtin testing using the following command:

```
cd messageApi/
go test ./...
```