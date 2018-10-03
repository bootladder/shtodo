#!/bin/bash
set -e

echo ========= UNIT TESTS ================
echo running.... go test
go test
echo ========= ACCEPTANCE TESTS ================
echo running.... go build.....
go build
docker run -t --rm --volume=$(pwd):/opt hello /opt/bats/bin/bats /opt/tests/test_shtodo1.bats
docker run -t --rm --volume=$(pwd):/opt hello /opt/bats/bin/bats /opt/tests/test_shtodo2.bats
docker run -t --rm --volume=$(pwd):/opt hello /opt/bats/bin/bats /opt/tests/test_shtodo3.bats
#docker run -t --rm --volume=$(pwd):/tmp hello /bin/bash -c "cd /tmp && bundle exec cucumber"
#docker run -t --rm --volume=$(pwd):/opt hello ls -al /tmp/
rm shtodo
