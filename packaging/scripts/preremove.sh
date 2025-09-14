#!/bin/sh
set -e

systemctl stop redirects.service || true
systemctl disable redirects.service || true
