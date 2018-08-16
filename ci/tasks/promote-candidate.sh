#!/usr/bin/env bash

set -e -x

source bosh-cpi-release/ci/tasks/utils.sh

check_param S3_ACCESS_KEY_ID
check_param S3_SECRET_ACCESS_KEY

source /etc/profile.d/chruby.sh
chruby 2.4.4

gem install bosh_cli --no-ri --no-rdoc

integer_version=$( cat version-semver/number | cut -d. -f1 )
echo $integer_version > promoted/integer_version

cp -r bosh-cpi-release promoted/repo

dev_release=$(echo $PWD/bosh-cpi-dev-artifacts/*.tgz)

pushd promoted/repo
  set +x
  echo creating config/private.yml with blobstore secrets
  cat > config/private.yml << EOF
---
blobstore:
  s3:
    access_key_id: $S3_ACCESS_KEY_ID
    secret_access_key: $S3_SECRET_ACCESS_KEY
EOF
  set -x

  echo "using bosh CLI version..."
  bosh version

  echo "finalizing CPI release..."
  bosh finalize release ${dev_release} --version $integer_version

  rm config/private.yml

  git diff | cat
  git add .

  git config --global user.email zhanggbj@cn.ibm.com
  git config --global user.name zhanggbj
  git commit -m "New final release v $integer_version"
popd


