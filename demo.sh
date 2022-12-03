#!/bin/bash

source .helpscout.env
HELPSCOUT_API_KEY="$HELPSCOUT_API_KEY" go run ./cmd/demo/main.go
