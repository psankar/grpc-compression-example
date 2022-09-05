# grpc-compression-example

## Dependencies:
* Go compiler - for compiling the gRPC server
* Typescript compiler - for compiling the gRPC client
* Node JS - for launching the gRPC client

# launch server
```
$ cd server
server $ go run main.go ;# Should launch the gRPC server in 8080
```

# launch client
```
$ cd client
client $ npm install
client $ tsc app.ts && node app.js
```
