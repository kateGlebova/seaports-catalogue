# Seaports catalogue

This application provides information about seaports. 

## Description 

The app consists of two services:

1. ClientAPI (`api`)
   -  parse JSON file with ports data 
   -  interact with PortDomainService to save it
   -  provide REST API 
2. PortDomainService (`repository`)
   -  saving ports data and retrieving it from the database
   
ClientAPI communicates with PortDomainService via gRPC.

## Getting started

1. To populate database with initial seaports data, add `ports.json` file to `data` directory. It should contain JSON object of the following format:
    ```json
    {
      "AEAJM": {
        "name": "Ajman",
        "city": "Ajman",
        "country": "United Arab Emirates",
        "alias": [],
        "regions": [],
        "coordinates": [
          55.5136433,
          25.4052165
        ],
        "province": "Ajman",
        "timezone": "Asia/Dubai",
        "unlocs": [
          "AEAJM"
        ],
        "code": "52000"
      },
      ...
    }  
    ``` 

2. Docker images of both services are on DockerHub, so you only need `docker-compose` to run this project:
    ```bash
    make run
    ```
 3. Now, you can use ClientAPI to CRUD seaports
    ```bash
    curl '127.0.0.1:8080/ports?limit=2&offset=12'
    ```
    
## ClientAPI

**Fetching ports:**
``` 
GET /ports?limit={limit}&offset={offset}
```
_required_ limit is a maximum number of ports returned
_optional_ offset is a number of ports to skip (0 if not set)

**Creating the port:**
``` 
POST /ports
```
Port info should be supplied in the request body. Example:
```json
{
  "id": "AEAJM",
  "name": "Ajman",
  "city": "Ajman",
  "country": "United Arab Emirates",
  "alias": [],
  "regions": [],
  "coordinates": [
    55.5136433,
    25.4052165
  ],
  "province": "Ajman",
  "timezone": "Asia/Dubai",
  "unlocs": [
    "AEAJM"
  ],
  "code": "52000"
}
```
**Fetching the port:**
``` 
GET /ports/{id}
```
**Updating the port:**
``` 
PUT /ports/{id}
```
Fields you want to update should be supplied in the request body. Example:
```json
{
  "country": "United Arab Emirates",
  "province": "Ajman"
}
```
**Deleting the port:**
``` 
DELETE /ports/{id}
```

## Building images manually

You can build images manually:

1. To build `client-api` image:
    ```bash
    make api
    ```
2. To build `port-domain-service` image:
   ```bash
   make repository
   ```
   
## Services configuration

You can control services configuration with environment variables.

### ClientAPI

Environment variable | Default value | Description 
--- | --- | ---
REPOSITORY | repository:8080 | PortDomainService address 
DATA_FILE | /opt/api/data/ports.json | file with ports data for initial population 
PORT | 8080 | port service should listen

### PortDomainService

Environment variable | Default value | Description 
--- | --- | ---
MONGO_URL | mongo:27017 | MongoDB address
MONGO_DB | ports | MongoDB database to use
MONGO_COLLECTION | ports | MongoDB collection
PORT | 8080 | port service should listen

## Possible improvements

1. Tests for gRPC client and server
2. Add logging for PortDomainService