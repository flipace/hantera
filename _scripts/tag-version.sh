if [ "$1" != "" ]; then
    git tag -a "v${1}" -m "version v${1}"
else
    echo "must provide version parameter (e.g. ./version.sh 1.0.1)"
fi
