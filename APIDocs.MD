# Development Usage (Running in terminal)

- Run `go run cmd/main.go -cmd cli` to have access to the `REPL` to type commands


# To make an insert

- `insert <row_id> <username> <email>`

# To make a select
- `select` returns all rows
- `select <row_id>` returns row with selected ID


# Development Usage (over REST endpoints)

- Set up port number in `.env` in root directory, as shown in `env-local`

- Run `go run cmd/main.go -http` to run server

- API Endpoints

## createUser

**Request**

+ `POST /createUser`

+ `Content-Type:` application/json

+ `JSON body:` 
```
{
"id": "34",
"username": "kb"
"email": "k@mail.com"
}
```

**Response**
```
{
    "data": {
        "ID": "34",
        "Username": "kb",
        "Email": "k@mail.com"
    },
    "message": "record has successfully been created"
}
```
# readUser

**Request**
+ `GET /readUser/<userID>`

**Response**
```
{
    "data": "ID: 1, Name: d, Email: d",
    "message": "record has successfully been read"
}
```
# readUsers

**Request**
+ `GET /readUsers`

**Response**
```
{
    "data": [
        "ID: 2, Name: d, Email: d",
        "ID: 1, Name: d, Email: d",
        "ID: 5, Name: kwakubiney, Email: k@mail.com"
    ],
    "message": "records have successfully been read"
}
```


