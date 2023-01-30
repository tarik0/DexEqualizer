#!/usr/bin/env sh

# Remove the old one.
rm -f ./scripts/output/dexeq_out
rm -f ./scripts/output/dexeq_out.zip

# Build the executable.
go build -gcflags="all=-N -l" -o ./scripts/output/dexeq_out ./main.go

# Zip the executable.
cd ./scripts/output || exit
zip ./dexeq_out.zip ./dexeq_out