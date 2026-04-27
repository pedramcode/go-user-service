#!/bin/bash

if [ "$1" = "" ]; then
    echo "migration file name is required as first parameter"
    exit 1
fi

migrate create -ext sql -dir migrations -seq $1