# liveplant Server

## Structs

### User

#### Example

```js
{
  "id": "UUID",
  "email": "austin@austinpray.com",
  "secret": "big long string"
}
```

### Poll

#### Example

```js
{
  "id": "UUID",
  "action": "water",
  "displayname": "Water The Plant",
  "deadline": "2015-01-1T00:00:00Z",
  "yee": [
    // User UUID
  ],
  "orNah": [
    // User UUID
  ],
  "yeeCount": 1,
  "orNahCount": 1
}
```

### Plant

#### Example

```js
{
  "name": "big-john",
  "displayName": "Big John",
  "currentPolls": [
    // []Poll
  ]
}
```

## Endpoints

All endpoints are prefixed with `/api/v1/`

### GET `plants`

Return all the plants the server knows about.

### GET `plants/:name`

Return a specific plant. 

### POST `login`

#### Request

```js
{
  "email": "austin@austinpray.com",
  "secret": "big long string"
}
```

#### Response

```js
{
  "success": true,
  "email": "austin@austinpray.com",
  "secret": "big long string"
}
```

### POST `vote/:id`

```js
{
  "email": "austin@austinpray.com",
  "secret": "big long string",
  "yee": true,
  "orNah": false
}
```

### GET `whatdo`

Return what the bot client should be doing right now. Grabs freshly expired
deadlines (5 minutes).
