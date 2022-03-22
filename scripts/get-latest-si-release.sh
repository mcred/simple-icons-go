#!/bin/sh

set -e

if [[ -z "$GITHUB_PAT" ]]; then
  echo "Set the GITHUB_PAT env variable."
  exit 1
fi

LATEST_TAG=$(curl -s -X GET -H "Authorization: token $GITHUB_PAT" \
  "https://api.github.com/repos/simple-icons/simple-icons/releases" \
  | jq -r 'first(.[].tag_name)')

echo $LATEST_TAG