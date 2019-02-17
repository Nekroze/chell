#!/bin/sh
set -euf
make \
    --warn-undefined-variables \
    --keep-going \
    --makefile "$@"
