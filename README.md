Liveplant Server [![Build Status](https://travis-ci.org/liveplant/liveplant-server.svg?branch=master)](https://travis-ci.org/liveplant/liveplant-server)
====

Liveplant server is a REST API for [liveplant.io][]

## Requirements

<table>
  <tr>
    <th>Thing</th>
    <th>Version</th>
    <th>Install With</th>
  </tr>
  <tr>
    <td>
      <a href="https://golang.org">
        go
      </a>
    </td>
    <td>1.4</td>
    <td>
      <a href="https://github.com/meatballhat/gimme#installation--usage">
        gimme
      </a>
    </td>
  </tr>
  <tr>
    <td>
      <a href="https://github.com/ddollar/foreman">
        foreman
      </a> (optional)
    </td>
    <td>stable</td>
    <td><pre>gem install foreman</pre></td>
  </tr>
</table>

## Endpoints

### GET `/current_action`

This endpoint is polled by the liveplant hardware. It exposes an enumerable
action string and a unix timestamp integer. The hardware knows when it last took action. An
is only taken action if the new action's timestamp is newer.

- [example](schema/current_action/GET/example.json)
- [JSON Schema](schema/current_action/GET/schema.json)

[liveplant.io]: https://github.com/liveplant/liveplant.io
[foreman]: https://github.com/ddollar/foreman
