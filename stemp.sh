#!/bin/sh
echo "Triggered from .go file"
rm $2
ln -s $1 $2