#!/usr/bin/env bash

export TAG=${VERSION}-${COMMIT}-${TRAVIS_BRANCH}-go${TRAVIS_GO_VERSION}
echo REPO=${REPO}, TAG=${TAG}, COMMIT=${COMMIT}, VERSION=${VERSION}, TRAVIS_BUILD_NUMBER=${TRAVIS_BUILD_NUMBER}, TRAVIS_GO_VERSION=${TRAVIS_GO_VERSION}

docker build -t ${REPO}:${COMMIT} .
docker tag ${REPO}:${COMMIT} quay.io/${REPO}:${COMMIT}
docker tag ${REPO}:${COMMIT} quay.io/${REPO}:${TAG}
docker tag ${REPO}:${COMMIT} quay.io/${REPO}:travis-${TRAVIS_BUILD_NUMBER}

docker images

docker login -u="$QUAY_USER" -p="$QUAY_PASS" quay.io
docker push quay.io/${REPO}:${COMMIT}
docker push quay.io/${REPO}:${TAG}
docker push quay.io/${REPO}:travis-${TRAVIS_BUILD_NUMBER}

if [ "${TRAVIS_BRANCH}" == "master" ] && [ "${TRAVIS_GO_VERSION}" == "1.8" ]; then
  docker tag ${REPO}:${COMMIT} quay.io/${REPO}:latest
  docker push quay.io/${REPO}:latest
fi