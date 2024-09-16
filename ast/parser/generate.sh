#!/bin/bash
# alias antlr4='java -Xmx500M -cp "./antlr-4.13.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -visitor -package parsing ./*.g4 -o ../parsing
# antlr4 -Dlanguage=Go -package parsing ./*.g4 -o ../parsing