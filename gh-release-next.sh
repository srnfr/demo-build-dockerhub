#!/bin/bash


LASTVERSION=$(gh release view --json name --jq '.name')
## Remove first char
LASTVERSION="${LASTVERSION:1}"
echo "Last release was : $LASTVERSION"

NEWVERSION=$((LASTVERSION+1))
echo "New version to release is : $NEWVERSION"

NEWVERSION="v$NEWVERSION"

gh release create $NEWVERSION --generate-notes
