#!/bin/bash

echo "Cloning github state"

./make.gits.sh


echo "Copying static assests"

rm -rf ../gosocket-publish
mkdir -p ../gosocket-publish/www

cp index.html ../gosocket-publish/www/
mkdir ../gosocket-publish/www/workspace


cp -ra ../gosocket-publish/www ../gosocket-publish/linux.amd64/


echo "Making .gits"

./copy.gits.sh


echo "Building for Linux"

env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../gosocket-publish/linux.amd64/gosocket .

./pullgits.sh $1

echo "gosocket-publish completed"