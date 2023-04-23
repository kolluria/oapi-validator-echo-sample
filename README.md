# oapi-validator-echo-sample

[![Build Status](https://github.com/kolluria/oapi-validator-echo-sample/actions/workflows/go.yml/badge.svg)](https://github.com/kolluria/oapi-validator-echo-sample/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kolluria/oapi-validator-echo-sample)](https://goreportcard.com/report/github.com/kolluria/oapi-validator-echo-sample)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a sample project showcasing how to use [oapi-codegen](https://github.com/deepmap/oapi-codegen) to generate an [Echo](https://echo.labstack.com/) server and then use [oapi-validator](https://github.com/williamhaley/oapi-codegen/tree/master/pkg/middleware) middleware to validate the OpenAPI specification.

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
