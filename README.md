# Helpscout Toolkit

Import documentation structure from HelpScout to JSON

## Environment Variables

Create file `.helpscout.env` in the root of the project with the following variables:

```bash
HELPSCOUT_API_KEY="..."
```

## Deployment

```
doctl serverless deploy . --include <funcName> --env .helpscout.env
```

funcName is a path to the function file, e.g. `helpscout/list-docs`

## Function URL

```
doctl sls fn get <funcName> --url
```
