#!/usr/bin/env bash

# directory configurations
workflow_dir=$(grep -i -r -l super-momo-tools /Users/MoMo/Dropbox/配置/Alfred-inuse-20190522/Alfred.alfredpreferences/workflows | sed "s/info\.plist//g")

# constant vars
current_dir=$(cd "$(dirname "$0")";pwd)

# build flags
go build -ldflags "-s -w"  -o ${workflow_dir}/exec/show-local-ip  ${current_dir}/../network/local
go build -ldflags "-s -w"  -o ${workflow_dir}/exec/search-ip-location  ${current_dir}/../network/wlan
go build -ldflags "-s -w"  -o ${workflow_dir}/exec/encode-decode  ${current_dir}/../encryption/code

# upx compress the files
if [[ ! $(command -v upx) ]]; then
  brew install upx
fi
cd ${workflow_dir}/exec/
ls | xargs upx