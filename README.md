## Running the server

### With hot reloading on file changes

```
fresh
```

Note: assumes you have [fresh](https://github.com/gravityblast/fresh) installed

### Without hot reloading

```
go run albums.go
```

## Testing the endpoints

Note: all these commands are assuming the server is running on Port 8888.

### Get a list of albums (GET)

```
curl http://localhost:8888/albums \
--header "Content-Type: application/json" \
--request "GET"
```

### Add a new album (POST)

```
curl http://localhost:8888/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

### View a single album (GET by ID)

```
curl http://localhost:8888/albums/2
```
