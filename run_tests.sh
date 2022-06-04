#!/bin/sh

echo ..........................................................................
echo
echo "Running Go tests"
echo
export GO111MODULE="on"
export GOPATH=""
cd ./go
go test -v .
cd ..
echo
echo ..........................................................................
echo
echo "Running Javascript tests"
echo
node js/test_money.js
echo
echo ..........................................................................
echo
echo "Running Python tests"
echo
python3 py/test_money.py -v
echo ..........................................................................
