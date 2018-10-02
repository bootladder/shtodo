#!/bin/bash
set -e

echo ========= UNIT TESTS ================
echo running.... go test
go test
echo ========= ACCEPTANCE TESTS ================
echo running.... go build.....
go build
docker run -t --rm --volume=$(pwd):/tmp hello /tmp/bats/bin/bats /tmp/tests/test_shtodo.bats
#docker run -t --rm --volume=$(pwd):/tmp hello /bin/bash -c "cd /tmp && bundle exec cucumber"
#docker run -t --rm --volume=$(pwd):/tmp hello echo hello
rm shtodo
