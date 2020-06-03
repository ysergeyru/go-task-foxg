#!/bin/bash
source cmd/service/env.sh
go run pg/migrations/*.go init
go run pg/migrations/*.go up
