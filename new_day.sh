#!/bin/bash

DAY=$1
echo "Creating DAY $DAY"
set -x
cp adventofcode/day2/ "adventofcode/day$DAY" -r
mv adventofcode/day$DAY/day2.go "adventofcode/day$DAY/day$DAY.go"
mv adventofcode/day$DAY/day2_test.go "adventofcode/day$DAY/day$DAY""_test.go"
sed -i  "s/Day2/Day$DAY/g" "adventofcode/day$DAY/day$DAY.go"
sed -i  "s/Day2/Day$DAY/g" "adventofcode/day$DAY/day$DAY""_test.go"
sed -i  "s/day2/day$DAY/g" "adventofcode/day$DAY/day$DAY.go"
sed -i  "s/day2/day$DAY/g" "adventofcode/day$DAY/day$DAY""_test.go"
