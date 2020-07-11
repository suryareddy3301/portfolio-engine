#!/bin/bash

echo "Copying assets ..."
sudo cp -r  assets/  /etc/portfolio/

echo "Building project"
go build

echo "running the app ..."
./portfolio
