# Tiresias

`Tiresias` is a tool to help you access your servers. With the support shell completion, you can use `ping production-server` or `ssh staging-server` smoothly.

## Features

- Read from muiltple sources
  - YAML files with glob path support
  - Consul
- Write into muiltple destinations
  - hosts
  - ssh config

## Quick start

Create a config file like following:

```yaml
database: /home/xuanwo/.tiresias/db

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

## Source

Every source should be like:

```yaml
- type: consul
  options:
    address: 1.2.4.8:8500
    schema: http
    datacenter: test
    prefix: test-
  default:
    user: root
    identity_file: ~/.ssh/key
```

- `type` is the type for source.
- `options` is options for source, and different source could have different options.
- `default` is the default value for current source.

### fs

Available options:

```yaml
# A path to yaml files.
path: /home/xuanwo/.tiresias/server/*.yaml
```

Every file should be in format like:

```yaml
- name: x-blog
  address: 23.92.27.86
  user: root
  identity_file: ~/.ssh/keyA

- name: x-dev
  address: 127.0.0.1
  user: root
  port: 20000
  identity_file: ~/.ssh/keyB
```

### consul

Available options:

```yaml
# tiresias will use address to connect consul.
address: 1.2.4.8:8500
# tiresias will use the schema to build connection to consul.
schema: http
# tiresias will connect the specific datacenter in consul.
datacenter: test
# prefix is the prefix for generated server's name.
prefix: test-
```

## Destination

Every destination should be like:

```yaml
- type: hosts
  options:
    path: /etc/hosts
```

- `type` is the type for destination.
- `options` is options for destination, different destination could have different options.

### hosts

Available options:

```yaml
path: /etc/hosts
```

Example results:

```hosts
# -- Generated by tiresias at 2018-07-20 14:09:52.194975225 +0800 CST m=+3.084879404 --
192.168.10.1 testing-node-1
192.168.10.2 testing-node-2
192.168.10.3 testing-node-3
192.168.10.4 testing-node-4
```

### ssh_config

Available options:

```yaml
path: /home/xuanwo/.ssh/config
```

Example results:

```ssh_config
# -- Generated by tiresias at 2018-07-20 14:09:52.195502566 +0800 CST m=+3.085407020 --
Host testing-node-1
    HostName 192.168.10.1
    User root
    IdentityFile ~/.ssh/test_key

Host testing-node-2
    HostName 192.168.10.2
    User root
    IdentityFile ~/.ssh/test_key

Host testing-node-3
    HostName 192.168.10.3
    User root
    IdentityFile ~/.ssh/test_key

Host testing-node-4
    HostName 192.168.10.4
    User root
    IdentityFile ~/.ssh/test_key
```

## Installation

Get the latest tiresias for Linux from [releases](https://github.com/Xuanwo/tiresias/releases)

## LICENSE

The Apache License (Version 2.0, January 2004).
