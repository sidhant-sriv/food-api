Create a new record:

POST /record

```
{
    "orderID": "string"
}
```
Response:
```
{
    "status": "string",
    "message": "string"
}
```

Get a record:
GET /order/{orderID}

Response:
```
{
    "orderID": "string",
    "orderStatus": "string",
}
```

POST /prepare/{orderID}

Response:
```
{
    "orderID": "string",
    "orderStatus": "string"
}
```

POST /dispatch/{orderID}

Response:
```
{
    "orderID": "string",
    "orderStatus": "string"
}
```