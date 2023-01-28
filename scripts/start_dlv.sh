#!/usr/bin/env bash

# Build it again.
./build.sh

# Kill old delve.
pkill dlv

# Export IS_DEV to true.
export IS_DEV=true

# Attach delve.
cd ..
dlv --listen=:2345 --continue --check-go-version=false --headless=true --api-version=2 --accept-multiclient exec ./dexeq_out ws://localhost:8546
# dlv debug --headless --listen=0.0.0.0:2345 --api-version=2 --accept-multiclient --check-go-version=false