#! /bin/bash
mv go.* ../;
SERVICE_NAME=greeting-svc go run server/main.go || mv ../go.* .