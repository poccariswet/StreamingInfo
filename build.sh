#!/bin/zsh
echo Building counter...
cd counter
go build -o counter

echo Building twittervotes...
cd ../twittervotes
go build -o twittervotes

echo Building api...
cd ../api
go build -o api

echo Building public...
cd ../public
go build -o public

echo Done.
