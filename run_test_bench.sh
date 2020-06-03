#!/bin/bash
source cmd/service/env.sh
go test -bench=. cmd/service/*.go
