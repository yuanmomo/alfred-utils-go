#!/usr/bin/env bash

# directory configurations
workflow_dir=$(grep -i -r -l super-momo-tools /Users/MoMo/Dropbox/配置/Alfred-inuse-20190522/Alfred.alfredpreferences/workflows | sed "s/info\.plist//g")

# constant vars
current_dir=$(cd "$(dirname "$0")";pwd)

go build -o ${workflow_dir}/exec/show-local-ip  ${current_dir}/../network/local
go build -o ${workflow_dir}/exec/search-ip-location  ${current_dir}/../network/wlan
