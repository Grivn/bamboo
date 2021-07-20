#!/usr/bin/env zsh
    echo "sending exutables into remote sites"
    go build ../../server/
    go build ../../client/
    echo "Servers are already started in this folder."
