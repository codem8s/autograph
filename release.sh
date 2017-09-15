#!/usr/bin/env bash

export BUILD_TAG="travis-$TRAVIS_BUILD_NUMBER-$TRAVIS_BRANCH-$COMMIT-go$TRAVIS_GO_VERSION"

if [ "${TRAVIS_BRANCH}" == "master" ] && [ "${TRAVIS_GO_VERSION}" == "1.8" ]; then
  export LATEST_TAG="latest"
  export VERSION_TAG="$VERSION"
else
  export LATEST_TAG="latest-$TRAVIS_GO_VERSION"
  export VERSION_TAG="$COMMIT-$TRAVIS_GO_VERSION"
fi

echo "REPO=$REPO, COMMIT=$COMMIT, BUILD_TAG=$BUILD_TAG, VERSION_TAG=$VERSION_TAG, LATEST=$LATEST_TAG"

docker build -t ${REPO}:${COMMIT} .
docker tag ${REPO}:${COMMIT} quay.io/${REPO}:${VERSION_TAG}
docker tag ${REPO}:${COMMIT} quay.io/${REPO}:${BUILD_TAG}
docker tag ${REPO}:${COMMIT} quay.io/${REPO}:${LATEST_TAG}

docker images

docker login -u="$QUAY_USER" -p="$QUAY_PASS" quay.io
docker push quay.io/${REPO}:${VERSION_TAG}
docker push quay.io/${REPO}:${BUILD_TAG}
docker push quay.io/${REPO}:${LATEST_TAG}
