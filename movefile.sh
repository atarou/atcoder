#!/bin/bash

echo -n DIR?
read dir
echo -n QUESTION?
read q
mkdir -p _$dir/$q
mv -i main.go _$dir/$q
git add _$dir/$q/main.go
git commit -m "$dir $q"

