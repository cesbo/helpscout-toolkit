# Helpscout Toolkit

[![Go Reference](https://pkg.go.dev/badge/github.com/cesbo/helpscout-toolkit/pkg/helpscout.svg)](https://pkg.go.dev/github.com/cesbo/helpscout-toolkit/pkg/helpscout)

Import documentation structure from HelpScout to JSON

## HelpScout Library

### Installation

```
go get github.com/cesbo/helpscout-toolkit/pkg/helpscout
```

## Tool

Tool implemented for DigitalOcean Functions

### Environment Variables

Create file `.helpscout.env` in the root of the project with the following variables:

```bash
HELPSCOUT_API_KEY="..."
```

### Deployment

```
doctl serverless deploy . --include <funcName> --env .helpscout.env
```

funcName is a path to the function file, e.g. `helpscout/summary`

### Function URL

```
doctl sls fn get <funcName> --url
```
