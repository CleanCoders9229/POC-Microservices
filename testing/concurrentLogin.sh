#!/bin/zsh
for i in {1..10}
do
    curl -d @loginData.txt -X POST http://localhost:8000/user/login
    echo "\n"
done
