#!/bin/bash

echo "Cleaning up"

rm -rf ../gosocket-publishgits
mkdir -p ../gosocket-publishgits

cd ../gosocket-publishgits

mkdir linux.amd64

echo "Cloning"

cd ./linux.amd64
git clone --depth=1 --branch=linux.amd64 --single-branch https://github.com/HasinduLanka/gosocket
