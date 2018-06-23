# Golang workspace environment

This directory contains a Golang application auto build versioning example.

This example is based on [goworkspace](https://github.com/harobed/goworkspace) build environment.

```
$ ./bin/darwin_amd64/myproject version
Version: master
UTC Build Time: 2018-06-23_04:18:27PM
Git Commit Hash: 5e980301d3e3e401473e7b8d7eb4d295d34111e2
Git Commit Date: 2018-06-23_17:54:06_+0200
```


## Prerequisite

* Docker and docker-compose

Prerequisites installation with [Brew](https://brew.sh/index_fr):

```
$ brew cask install docker
```


## Usage

Build the project:

```
$ make up-and-build
```

Version output example:

```
$ ./bin/darwin_amd64/myproject version
Version: master
UTC Build Time: 2018-06-23_04:18:27PM
Git Commit Hash: 5e980301d3e3e401473e7b8d7eb4d295d34111e2
Git Commit Date: 2018-06-23_17:54:06_+0200
```
