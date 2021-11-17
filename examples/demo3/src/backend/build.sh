#!/bin/sh

# Get full path to script directory
SCRIPT=$(readlink -f "$0")
BASEDIR=$(dirname "$SCRIPT")

# Create build directory
if [ "$#" -ne 1 ]; then
    echo "Missing bin directory"
    exit
fi
BINDIR=$1
mkdir -p $BINDIR

# Copy config file to bin
cp $BASEDIR/util/app_instance.yaml $BINDIR/demo-server

# Build demo App server
cd $BASEDIR

go build -o $BINDIR/demo-server .
