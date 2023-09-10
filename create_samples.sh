#!/bin/bash

mkdir samples
for i in {1..20}; do truncate -s $(( i*1000*1000 )) ./samples/file$i.txt;done;