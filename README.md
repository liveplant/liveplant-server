Liveplant Server
====

Liveplant server is a REST api for [liveplant.io][]

## Endpoints

### GET `/current_action`

This endpoint is polled by the liveplant hardware. It exposes an enumerable
action string and a unix timestamp integer. The hardware knows when it last took action. An
is only taken action if the new action's timestamp is newer.

- [example](schema/current_action/GET/example.json)
- [JSON Schema](schema/current_action/GET/schema.json)

[liveplant.io]: https://github.com/liveplant/liveplant.io
