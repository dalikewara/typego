#!/bin/sh

git pull github master || true
git add . || true
git status || true

if [ "$1" != "" ]; then
    git commit -m "$1" || true
else
    git commit -m "update some scripts" || true
fi

git push github master
