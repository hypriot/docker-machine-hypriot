# **WORK IN PROGRESS**

You may find our still working docker-machine-hypriot up to version 0.4.1 binary at http://blog.hypriot.com/downloads/
That is a complete `docker-machine` binary with the hypriot driver and provisioner compiled into it. Do not use this binary as a plugin for docker-machine 0.5.x.
At the moment this plugin is rather useless as there is no provisioner plugin concept in docker-machine. So the current plan is to build a special docker-machine binary 0.5.x again with a hypriot provisioner compiled into it.

**See discussion in these issues for more details:**
  * [hypriot/docker-machine-hypriot#1](https://github.com/hypriot/docker-machine-hypriot/issues/1#issuecomment-169363905)
  * [docker/machine#2822](https://github.com/docker/machine/issues/2822)

If you want to contribute or build the binary from source read the [Contributing Guide](https://github.com/hypriot/docker-machine-hypriot/blob/master/CONTRIBUTING.md)

# Docker Machine Hypriot Driver

This is a plugin for [Docker Machine](https://docs.docker.com/machine/) allowing
to connect to remote Raspberry Pis running [HypriotOS](http://blog.hypriot.com).

## Requirements
* [Docker Machine](https://docs.docker.com/machine/) 0.5.0+

## Installation

To install this plugin manually, download the binary `docker-machine-driver-hypriot`
and  make it available by `$PATH`, for example by putting it to `/usr/local/bin/`:

```bash
$ curl -L  -o /usr/local/bin/docker-machine-driver-hypriot https://github.com/hypriot/docker-machine-hypriot/releases/download/v0.0.1/docker-machine-driver-hypriot
$ chmod +x /usr/local/bin/docker-machine-driver-hypriot
```

The latest version of `docker-machine-driver-hypriot` binary is available on
the ["Releases"](https://github.com/hypriot/docker-machine-hypriot/releases) page.


## Usage

Official documentation for Docker Machine [is available here](https://docs.docker.com/machine/).

Create machines using an existing Raspberry Pi with installed [Hypriot](http://blog.hypriot.com/) image (based on Raspbian).

For more details on installing the basic system (including a preinstalled Docker daemon) please
see the detailled tutorial [Getting started with Docker on your Raspberry Pi](http://blog.hypriot.com/getting-started-with-docker-on-your-arm-device/).

Before using the `docker-machine create` command you have to copy your SSH keys to the device:

```
$ ssh-copy-id root@$HYPRIOT_IP_ADDRESS
```

Now you can easily provision your Docker engine with:

```
$ docker-machine create --driver hypriot --hypriot-ip-address=$HYPRIOT_IP_ADDRESS black-pearl
```

In order to upgrade to the latest available Docker version, just run:

```
$ docker-machine upgrade black-pearl
```

Options:

 - `--hypriot-ip-address`: **required** IP Address of host.
 - `--hypriot-ssh-user`: SSH username used to connect.
 - `--hypriot-ssh-key`: Path to the SSH user private key.
 - `--hypriot-ssh-port`: Port to use for SSH.

> Note: currently only a numerical IP address is supported.

Environment variables and default values:

| CLI option                 | Environment variable | Default             |
|----------------------------|----------------------|---------------------|
| **`--hypriot-ip-address`** | -                    | -                   |
| `--hypriot-ssh-user`       | -                    | `root`              |
| `--hypriot-ssh-key`        | -                    | `$HOME/.ssh/id_rsa` |
| `--hypriot-ssh-port`       | -                    | `22`                |

Once you have created your Docker machine run this command to configure your shell:

```
$ eval "$(docker-machine env black-pearl)"
```

Now you can use the Docker machine as usual:

```
$ docker version
Client:
 Version:      1.9.0
 API version:  1.21
 Go version:   go1.5.1
 Git commit:   76d6bc9
 Built:        Tue Nov  3 21:18:00 UTC 2015
 OS/Arch:      darwin/amd64

Server:
 Version:      1.9.0
 API version:  1.21
 Go version:   go1.4.3
 Git commit:   76d6bc9
 Built:        Tue Nov  3 19:20:09 UTC 2015
 OS/Arch:      linux/arm
```
