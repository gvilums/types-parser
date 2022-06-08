#!/usr/bin/bash
rm ./parser/*
java -jar ./antlr-4.10.1-complete.jar -Dlanguage=Go -o parser -visitor -no-listener Types.g4