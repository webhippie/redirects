#!/bin/sh
set -e

chown -R redirects:redirects /etc/redirects
chown -R redirects:redirects /var/lib/redirects
chmod 750 /var/lib/redirects

if [ -d /run/systemd/system ]; then
    systemctl daemon-reload

    if systemctl is-enabled --quiet redirects.service; then
        systemctl restart redirects.service
    fi
fi
