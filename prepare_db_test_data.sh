#!/bin/bash
go run pg/migrations/*.go init
go run pg/migrations/*.go up
