#!/usr/bin/env bash
set -euxo pipefail

mkdir -p ./ops/docs/coverage/
go install github.com/boumenot/gocover-cobertura@latest
go test -p 1 -count=1 -cover -coverprofile ./ops/docs/coverage/coverage-profile.txt ./...
go tool cover -func ./ops/docs/coverage/coverage-profile.txt | awk '/^total/{print "{\"total\":\""$3"\"}"}' > ./ops/docs/coverage/coverage.json
go tool cover -html ./ops/docs/coverage/coverage-profile.txt -o ./ops/docs/coverage/coverage.html
gocover-cobertura < ./ops/docs/coverage/coverage-profile.txt > ./ops/docs/coverage/coverage-cobertura.xml