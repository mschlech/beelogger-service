#!/bin/bash

echo "wrapper script around make and docker issues"

DOCKER=`which docker`
MAKE=`which make`


echo " make found in "+$MAKE

exit