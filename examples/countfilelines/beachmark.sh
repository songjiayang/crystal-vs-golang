#!/bin/bash

crystal build countfilelines.cr -o ../../bin/countfilelines-crystal
go build -o ../../bin/countfilelines-go  countfilelines.go

time for (( i=1; i<=1000; i++ ))
do
   ../../bin/countfilelines-crystal
done
echo "finish countfilelines-crystal test"

time for (( i=1; i<=1000; i++ ))
do
   ../../bin/countfilelines-go
done
echo "finish countfilelines-go test"
