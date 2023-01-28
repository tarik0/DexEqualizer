#!/usr/bin/env sh

# Remove the old one.
rm -f dexeq_out

# Build the executable.
go build -gcflags="all=-N -l" -o ../dexeq_out ../main.go