#!/bin/bash

echo "Syncing with github"

cd ../gosocket-publish/



cd ./linux.amd64
git add .
git update-index --add --chmod=+x gosocket
git commit -m "build-$1"
git push origin HEAD

