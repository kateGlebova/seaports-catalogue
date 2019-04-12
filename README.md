# Seaports catalogue

This application provides information about seaports. 

## Description 

The app consists of two services:

1. ClientAPI
   -  parse JSON file with ports data 
   -  interact with PortDomainService to save it
   -  provide REST API 
2. PortDomainService
   -  saving ports data and retrieving it from the database
