#!/bin/bash
set -e

echo ========= UNIT TESTS ================
echo running.... go test
go test
echo ========= ACCEPTANCE TESTS ================
echo running.... go build.....
go build

echo running.... test loop.....
echo
for TEST_FILE in $(cd tests; ls test_*.bats)
do
  echo ==== $TEST_FILE ====
  docker run -t --rm --volume=$(pwd):/opt hello /opt/bats/bin/bats /opt/tests/$TEST_FILE 
  echo 
done


#docker run -t --rm --volume=$(pwd):/tmp hello /bin/bash -c "cd /tmp && bundle exec cucumber"
#docker run -t --rm --volume=$(pwd):/opt hello ls -al /tmp/
rm shtodo
