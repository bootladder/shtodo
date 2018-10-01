echo running.... go test
go test
echo running.... go build.....
go build
docker run -t --rm --volume=$(pwd):/tmp hello /tmp/bats/bin/bats /tmp/tests/hello.bats
rm shtodo
