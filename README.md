# Tiresias

`Tiresias` is a tool to help you access your servers. With the support shell completion, you can use `ping production-server` or `ssh staging-server` smoothly.

## Features

- Generate ssh config
- Generate hosts files

## Quick start

Create a config file like following:

```yaml
src:
  - type: fs
    path: /path/to/source/file

dst:
  - type: ssh_config
    path: /home/xuanwo/.ssh/config
  - type: hosts
    path: /etc/hosts
```

And run:

```ssh
sudo tiresias --config ~/.tiresias/config.yaml
```

Your will see all hosts and ssh config generated.

## Installation

Get the latest tiresias for Linux from [releases](https://github.com/Xuanwo/tiresias/releases)

## LICENSE

The Apache License (Version 2.0, January 2004).
