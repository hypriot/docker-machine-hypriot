# Building

The requirements to build Machine are:

1.  A running instance of Docker or a Golang 1.5.2 development environment
2.  The `bash` shell
3.  [Make](https://www.gnu.org/software/make/)

## Build using Docker containers

To build the `docker-machine-hypriot` binary using containers, simply run:

    $ make build

## Local Go development environment

Make sure the source code directory is under a correct directory structure to use Go 1.5 vendoring;
example of cloning and preparing the correct environment `GOPATH`:

        mkdir docker-machine-hypriot
        cd docker-machine-hypriot
        export GOPATH="$PWD"
        go get github.com/hypriot/docker-machine-hypriot
        cd src/github.com/hypriot/docker-machine-hypriot

At this point, simply run:

    $ make build-local
