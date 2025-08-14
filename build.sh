#! /bin/bash

# build landing page
cd client/ && test -f package.json && npm run build; cd ..;

# build each tool
for dir in tools/*; do
    cd "$dir" && test -f package.json && npm run build; cd -;
done

# build go server
cd server/ && go build -o out  ../cmd/app;
