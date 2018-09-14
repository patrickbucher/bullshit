#!/bin/sh

rm -rf examples/
go build cmd/bullshit.go
for f in input/*
do
    folder=`basename "$f"`
    mkdir -p "examples/$folder"
    ./bullshit -n 5 < "$f"
    mv output/*.html "examples/${folder}/"
done
