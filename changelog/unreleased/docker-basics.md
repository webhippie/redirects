Change: Use Alpine base image and define healthcheck

We just switched the base image to Alpine Linux instead of a scratch image as
this could lead to issues related to healthcheck commands defined by the Docker
CLI. Beside that we have added the healthcheck command to the Dockerfile for
having a proper handling by default.

https://github.com/webhippie/redirects/pull/67
