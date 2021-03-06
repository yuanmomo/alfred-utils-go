#!/bin/bash
### thanks for https://www.alfredforum.com/topic/9873-how-to-package-a-workflow-via-the-command-line/

current_dir=$(cd "$(dirname "$0")";pwd)

#readonly workflow_dir="${1}"
workflow_dir=$(grep -i -r -l super-momo-tools /Users/MoMo/Dropbox/配置/Alfred-inuse-20190522/Alfred.alfredpreferences/workflows | sed "s/info\.plist//g")

rm -rf  ${workflow_dir}/cache/*
rm -rf  ${workflow_dir}/cache/*

readonly info_plist="${workflow_dir}/info.plist"

if [[ ! -f "${info_plist}" ]]; then
  echo "The worflow path should be a valid workflow directory : [${workflow_dir}]."
  exit 1
fi

readonly workflow_name="$(/usr/libexec/PlistBuddy -c 'print name' "${info_plist}")"
readonly workflow_file="${current_dir}/../${workflow_name}.alfredworkflow"

if /usr/libexec/PlistBuddy -c 'print variablesdontexport' "${info_plist}" &> /dev/null; then
  readonly workflow_dir_to_package="$(mktemp -d)"
  cp -R "${workflow_dir}/"* "${workflow_dir_to_package}"

  readonly tmp_info_plist="${workflow_dir_to_package}/info.plist"
  /usr/libexec/PlistBuddy -c 'Print variablesdontexport' "${tmp_info_plist}" | grep '    ' | sed -E 's/ {4}//' | xargs -I {} /usr/libexec/PlistBuddy -c "Set variables:'{}' ''" "${tmp_info_plist}"
else
  readonly workflow_dir_to_package="${workflow_dir}"
fi

ditto -ck "${workflow_dir_to_package}" "${workflow_file}"
echo "Exported worflow to ${workflow_file}."