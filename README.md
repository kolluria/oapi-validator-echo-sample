# oapi-validator-echo-sample

[![Build Status](https://github.com/kolluria/oapi-validator-echo-sample/actions/workflows/go.yml/badge.svg)](https://github.com/kolluria/oapi-validator-echo-sample/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kolluria/oapi-validator-echo-sample)](https://goreportcard.com/report/github.com/kolluria/oapi-validator-echo-sample)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a sample project showcasing how to use [oapi-codegen](https://github.com/deepmap/oapi-codegen) to generate an [Echo](https://echo.labstack.com/) server and then use [oapi-validator](https://github.com/williamhaley/oapi-codegen/tree/master/pkg/middleware) middleware to validate the incoming requests.
This project uses the [Petstore API](https://petstore3.swagger.io/) as the API definition and adds validations to the API definition.

## Table of Contents
- [oapi-validator-echo-sample](#oapi-validator-echo-sample)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installing](#installing)
    - [Project structure](#project-structure)
    - [Generating the Server](#generating-the-server)
    - [Running the Server](#running-the-server)
    - [Understanding the Code](#understanding-the-code)

## Getting Started

### Prerequisites

Make sure you have the following tools installed:

- [Go](https://golang.org/) (version 1.20)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen) (version 1.12.3)
- [Taskfile](https://taskfile.dev/#/installation) (version 3.20.0)

### Installing

Clone the repository:

```bash
git clone https://github.com/kolluria/oapi-validator-echo-sample.git
```

### Project structure

The project is structured as follows:

- `server-gen/`: Contains the generated server code.
- `utils/`: Contains the utility functions for the Service.
- `.gitignore`: Contains the list of files to be ignored by Git.
- `main.go`: Contains the main function for the Service.
- `README.md`: Contains the README for the Service.
- `Taskfile.yaml`: Contains the tasks for the Service.
- `swagger.yaml`: Contains the OpenAPI specification for the Service.

### Generating the Server

To generate the server after making any changes to `swagger.yaml`, run the following command:

```bash
task generate-server
```

### Running the Server

To run the server, run the following command:

```bash
task run
```

### Understanding the code generation

The following is a brief explanation of the server generation process:

* The `generate-server` task in `Taskfile.yaml` generates the server code and models using `oapi-codegen` and places it in the `server-gen/` directory.
* The configurations for the server and model generation are specified in `server-gen/server.cfg.yaml` and `server-gen/models.cfg.yaml` respectively.
* The config file for the server generation is as follows:
  ```yaml
  # Path: server-gen/server.cfg.yaml
  package: servergen
  output: server.go
  generate:
    echo-server: true
    embedded-spec: true
    strict-server: true
  ```
* The config file for the model generation is as follows:
  ```yaml
  # Path: server-gen/models.cfg.yaml
  package: servergen
  output: models.go
  generate:
    types: true
  ```
* Upon running the `generate-server` task, the directory structure looks like this:
  ```
  server-gen/
  ├── models.go
  ├── models.cfg.yaml
  ├── server.go
  └── server.cfg.yaml
  ```
### Understanding the validation process

* Since go-playground/validator is a tag-based validator, we need tags on the models to validate them.
* `oapi-codegen` provides a way to add custom tags to the models using `x-oapi-codegen-extra-tags` extension.
* To demonstrate this, I added the following tags to `swagger.yaml` - 
  * `validate: gte=1 lte=100` to quantity in `Order` schema.
  * `validate: datetime` to `shipDate` in `Order` schema.
  * `validate: oneof` to all the enums.
  * `validate: required` to all the required fields.
  * `validate: email` to `email` in `User` schema.
* To see the validation in action, run the server and send a request with invalid data to the `POST /store/order` endpoint.