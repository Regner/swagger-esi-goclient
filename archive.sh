#!/bin/bash

function archive {
  pushd .
  cd $1
  curl "https://esi.tech.ccp.is/$1/swagger.json" | python3.5 -m json.tool --sort-keys > swagger.json && git add swagger.json && git commit --author "ESI archive bot <archive@pizza.moe>" -m "checked in new $1 definition"
  popd
}

git pull

archive latest

git push
