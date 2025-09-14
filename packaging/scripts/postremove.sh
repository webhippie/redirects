#!/bin/sh
set -e

if [ ! -d /var/lib/redirects ] && [ ! -d /etc/redirects ]; then
    userdel redirects 2>/dev/null || true
    groupdel redirects 2>/dev/null || true
fi
