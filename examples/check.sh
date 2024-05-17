#!/bin/bash

testFile="/tmp/bullet"

if [[ -f $testFile ]]
then
	echo "$testFile found"
	exit 0
else
	echo "$testFile not found"
	exit 1
fi

