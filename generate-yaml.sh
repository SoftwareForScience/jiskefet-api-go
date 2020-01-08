#!/usr/bin/bash

TXT_RED='\033[0;31m'
TXT_NC='\033[0m'
TXT_GREEN='\033[0;32m'

url=$1
if [ ! $url ]; then
	echo -ne "\t${TXT_RED}Give the path of the doc-json as parameter.${TXT_NC}\n"
	exit
fi

snap list yq >/dev/null
if [ $? -ne 0 ]; then
	snap install yq
fi

echo -ne "\tRetrieving doc-json from: $url"

# Retrieve json from url
curl --silent  $url > doc.json 

# Transform json to yml
yq r doc.json > doc.yaml

# Remove all isArray: false lines
sed -i '/isArray: false/d' ./doc.yml

rm doc.json



echo -ne "\n\t${TXT_GREEN}Completed.${TXT_NC}\n"
