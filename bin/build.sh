#!/usr/bin/env bash

# directory configurations
workflow_dir=$(grep -i -r -l super-momo-tools /Users/MoMo/Dropbox/配置/Alfred-inuse-20190522/Alfred.alfredpreferences/workflows | sed "s/info\.plist//g")

echo "workflow_dir : ${workflow_dir}"
# constant vars
current_dir=$(cd "$(dirname "$0")";pwd)

# build flags
go build -ldflags "-s -w"  -o ${workflow_dir}/exec/momo-util  ${current_dir}/../main

# upx compress the files
if [[ ! $(command -v upx) ]]; then
  brew install upx
fi

cd ${workflow_dir}/exec/
ls | xargs upx --best