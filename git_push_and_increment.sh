#!/bin/bash

# Git upload
git add .
if [[ -z "${2//  }" ]]
then
  git commit
else
  git commit -m "$2"
fi
git pull origin $1 || exit
git push origin $1 || exit

sh ./git_increment_version.sh