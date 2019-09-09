#!/bin/sh

# If GOPATH is not set, exit
if [ -z "$GOPATH" ]; then
  echo "ERROR: GOPATH is not set"
  return 1
fi

echo "INFO: GOPATH=$GOPATH"

# If BUILD_VERSION is not set in args, set default build version
BUILD_VERSION=$1
if [ -z "$BUILD_VERSION" ]; then
  export BUILD_VERSION="local-dev"
  echo "WARN: Build Version is not set in args. Set to default version"
fi

echo "INFO: BUILD_VERSION=$BUILD_VERSION"

# If build folder is not set, make folder
if [ ! -d "$GOPATH/build" ]; then
  mkdir "$GOPATH/build"
  echo "INFO: build folder created in $GOPATH"
fi

# Copy error codes to build folder
cp error_codes.yml "$GOPATH/build"
echo "INFO: error_codes.yml copied"

# Compile
echo "INFO: building api..."
CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -ldflags "-X /flags.AppVersion=$BUILD_VERSION" \
    -installsuffix cgo \
    -o "$GOPATH/build/api" .

echo "INFO: Done"