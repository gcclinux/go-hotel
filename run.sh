#!/usr/bin/env sh
go build -o hotel-$(uname)-$(uname -m) cmd/web/*.go && ./hotel-$(uname)-$(uname -m)