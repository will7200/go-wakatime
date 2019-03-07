#!/bin/bash

TOTAL_TEST=$(cat run.txt)
PASSED_TEST=$(cat passed.txt)
TOTAL_P=$(awk -v p="$PASSED_TEST" -v t="$TOTAL_TEST" 'BEGIN{printf("%.00f\n",p / t * 100 )}')
BIN=5c81e66a2e4731596f19ea57
json=$(curl --request GET https://api.jsonbin.io/b/"$BIN"/latest)
if [ $TOTAL_P -lt 50 ]; then
   COLOR="red"
elif [ $TOTAL_P -lt 75 ]; then
   COLOR="yellow"
else
   COLOR="blue"
fi
NEW_JSON=$(echo $json | jq ".message = \"$TOTAL_P\"" | jq ".color = \"$COLOR\"")
curl --header "Content-Type: application/json" \
       --request PUT \
       --data "$NEW_JSON" \
       "https://api.jsonbin.io/b/$BIN"
