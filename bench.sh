#!/bin/bash
mv tests/*_test.go .
go test -run=XXX -bench=.
mv *_test.go tests/
