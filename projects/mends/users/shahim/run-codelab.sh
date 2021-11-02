#!/usr/bin/env bash
#set -x
set -e
set -u
set -o pipefail
set -o noclobber
#set -f # no globbing
shopt -s failglob # fail if glob doesn't expand

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do
    DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE"
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"

cd "$DIR"
SHAHIM_DIR="$(pwd)"

cd "../../mapping_engine/main"

go run .    -input_file_spec="${SHAHIM_DIR}/codelab/codelab.json" \
            -output_dir="${SHAHIM_DIR}/codelab" \
            -mapping_file_spec="${SHAHIM_DIR}/codelab/codelab.wstl" \
            -verbose=true  #TODO: not working
