#!/bin/bash
# source ./env.sh # for export to current shell
export TF_VAR_YANDEX_TOKEN=$(yc iam create-token)