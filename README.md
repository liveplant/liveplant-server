Liveplant Server [![Build Status](https://travis-ci.org/liveplant/liveplant-server.svg?branch=master)](https://travis-ci.org/liveplant/liveplant-server)
====

Liveplant server is a REST API for [liveplant.io][]

## Get started

### Development Environment

1. [Install Vagrant][]
2. run `vagrant up`
3. `vagrant ssh` to access your brand new, fully configured development
   environment.

### Development Commands

- `make` will compile and install liveplant-server inside your GOPATH.
- `make run` will compile and run your code. Your server will be available at
  [localhost:5000](http://localhost:5000).
- `make fmt` will format your code with gofmt.

## Usage

For usage run `liveplant-server -h`

```
-debug=false: Whether or not to enable debug logger.
```

## Environment Variables

- `LIVEPLANTDEBUG` set to `1` to enable debug level logging.

## Endpoints

### GET `/current_action`

This endpoint is polled by the liveplant hardware. It exposes an enumerable
action string and a unix timestamp integer. The hardware knows when it last took action. An
is only taken action if the new action's timestamp is newer.

- [example](schema/current_action/GET/example.json)
- [JSON Schema](schema/current_action/GET/schema.json)

[liveplant.io]: https://github.com/liveplant/liveplant.io
[foreman]: https://github.com/ddollar/foreman
[Install Vagrant]: https://docs.vagrantup.com/v2/installation/index.html
