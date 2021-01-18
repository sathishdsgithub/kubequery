# Copyright (c) 2020-present, The kubequery authors
#
# This source code is licensed as defined by the LICENSE file found in the
# root directory of this source tree.
#
# SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)

FROM ubuntu:20.04

ARG KUBEQUERY_VERSION="1.0"

LABEL \
  com.uptycs.description="Kubequery powered by Osquery" \
  com.uptycs.name="kubequery" \
  com.uptycs.version="$OSQUERY_VERSION" \
  com.uptycs.schema-version="1.0" \
  com.uptycs.url="https://github.com/Uptycs/kubequery" \
  com.uptycs.vendor="Uptycs Inc"

ADD https://pkg.osquery.io/deb/osquery_4.6.0-1.linux_amd64.deb /tmp/osquery.deb
ADD kubequery /usr/bin/kubequery

RUN set -ex; \
    DEBIAN_FRONTEND=noninteractive apt-get update -y && \
    DEBIAN_FRONTEND=noninteractive apt-get upgrade -y && \
    chmod 700 /usr/bin/kubequery && \
    dpkg -i /tmp/osquery.deb && \
    /etc/init.d/osqueryd stop && \
    rm -rf /var/osquery/* /var/log/osquery/* /var/lib/apt/lists/* /var/cache/apt/* /tmp/*

# TODO
ENTRYPOINT ["sleep", "3600"]
