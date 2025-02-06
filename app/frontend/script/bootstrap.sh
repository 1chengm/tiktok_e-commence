#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=./app/frontend
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}