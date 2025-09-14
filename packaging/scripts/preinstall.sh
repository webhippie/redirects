#!/bin/sh
set -e

if ! getent group redirects >/dev/null 2>&1; then
    groupadd --system redirects
fi

if ! getent passwd redirects >/dev/null 2>&1; then
    useradd --system --create-home --home-dir /var/lib/redirects --shell /bin/bash -g redirects redirects
fi
