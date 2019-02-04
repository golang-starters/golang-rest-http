# Golang REST HEEP
A booster demonstrating rest http and recovery using the kubernetes/openshift liveliness and readiness probes.

## How to run
### Using souce-to-image
The project uses openshift s2i to build docker image. Run the following command to build the docker image

` s2i build .  centos/go-toolset-7-centos7:latest golang-rest-http-app `

To start the web service, run

` docker run -p 8080:8080 golang-rest-http-app `

The web service should now be accessible on ` http://localhost:8080 `

### Using make
#### Build
```
$ make build
```
This will perform following actions: fetch dependencies, generate Goa files from design folder, compile. When in doubt just make help.

#### Run Server Locally
```
$ make dev
```
