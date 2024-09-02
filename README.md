# Instructions

Download from my repository
```
git clone https://github.com/canyouhearthemusic/sqlc-prac.git
cd sqlc-prac
```

Build docker image of the project
```
make build
```

Run it
```
make up
```

To Stop
```
make stop
```

To Start
```
make start
```

To Down Container
```
make down
```



# Endpoints

### `POST /api/v1/reservation`

Request Body:
```json
{
  "room_id": "15A",
  "start_time": "2024-01-24 15:06:00",
  "end_time": "2024-01-24 15:07:00"
}
```

Response Body:
`201 Created`
```json
{
  "message": "reservation created",
  "data": {
    "id": 1,
    "room_id": "15A",
    "start_time": "2024-01-24T15:06:00Z",
    "end_time": "2024-01-24T15:07:00Z"
  }
}
```

Trying to send that request again:
`409 Conflict`
```json
{
  "message": "reservation conflicts with an existing reservation",
  "data": {
    "room_id": "15A",
    "start_time": "2024-01-24T15:06:00Z",
    "end_time": "2024-01-24T15:07:00Z"
  }
}
```

### `GET "/api/v1/reservation/{room_id}"`

Response Body:
` 200 OK`
```json
{
  "data": [
    {
      "id": 1,
      "room_id": "15A",
      "start_time": "2024-01-24T15:06:00Z",
      "end_time": "2024-01-24T15:07:00Z"
    }
  ]
}
```
