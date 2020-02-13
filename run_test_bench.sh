#!/bin/bash
export CONFIG_PATH=../../config
STAGE=development go test -bench=. cmd/service/*.go
