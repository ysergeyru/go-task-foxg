#!/bin/bash
export CONFIG_PATH=~/go/src/github.com/ysergeyru/go-task-foxg/config
STAGE=development go test -bench=. cmd/service/*.go
