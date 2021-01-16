set -e

echo "*** build"
docker run --rm -v ${PWD}:/local -w=/local --entrypoint /local/bin/build.sh mryan321/utils_gradle_protoc:latest /home/gradle/googleapis /home/gradle/grpc-gateway /local/proto
