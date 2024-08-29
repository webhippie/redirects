# Changelog for unreleased

The following sections list the changes for unreleased.

## Summary

 * Chg #67: Use Alpine base image and define healthcheck
 * Chg #76: Replace libkv by valkeyrie

## Details

 * Change #67: Use Alpine base image and define healthcheck

   We just switched the base image to Alpine Linux instead of a scratch image as
   this could lead to issues related to healthcheck commands defined by the Docker
   CLI. Beside that we have added the healthcheck command to the Dockerfile for
   having a proper handling by default.

   https://github.com/webhippie/redirects/pull/67

 * Change #76: Replace libkv by valkeyrie

   We switched the library powering the support for etcd, zookeper and consul by a
   better maintained alternative named valkeyrie.

   https://github.com/webhippie/redirects/issues/76


# Changelog for 1.0.1

The following sections list the changes for 1.0.1.

## Summary

 * Fix #59: Bind flags correctly to variables

## Details

 * Bugfix #59: Bind flags correctly to variables

   We fixed the binding of flags to variables as this had been bound to the root
   command instead of the server command where it belongs to.

   https://github.com/webhippie/redirects/issues/59


# Changelog for 1.0.0

The following sections list the changes for 1.0.0.

## Summary

 * Chg #25: Initial release of basic version

## Details

 * Change #25: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/webhippie/redirects/issues/25


