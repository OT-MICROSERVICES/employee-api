# Employee REST API

Employee REST API is a golang based microservice which is responsible for all the employee related transactions in the [OT-Microservices](https://github.com/OT-MICROSERVICES). This application is completely platform independent and can be run on any kind of platform.

Supported features in the applications are:-

- Gin REST API for web transactions
- ScyllaDB as primary database for storing information
- Redis as a cache management system for quick response
- Swagger integration for the documentation of the API
- Prometheus's metrics support to monitor application health and performance

## Pre-Requisites

The application doesn't have any specific pre-requisites except the database connectivity. Additionally, we can add `Redis` as cache system but it's not part of the mandatory setup.

For running the application, we need following things configured:

- [ScyllaDB](https://www.scylladb.com/)
- [Redis](https:/redis.com/)

## Architecture

![](./static/employee.png)


