#!/bin/bash

crystal build sleep.cr -o ../../bin/sleep-crystal --release
go build -o ../../bin/sleep-go sleep.go

time for (( i=1; i<=10; i++ ))
do
   ../../bin/sleep-crystal
done
echo "finish sleep-crystal test"

time for (( i=1; i<=10; i++ ))
do
   ../../bin/sleep-go
done
echo "finish sleep-go test"
