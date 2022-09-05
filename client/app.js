"use strict";
exports.__esModule = true;
var exsvc = require("../libs/js/example_pb_service");
var expb = require("../libs/js/example_pb");
var grpc = require("grpc");
var grpcCli = new exsvc.ExampleServiceClient("localhost:8080", grpc.credentials.createInsecure(), { 'grpc.default_compression_algorithm': 3, 'grpc.default_compression_level': 3 });
console.log('grpc client is created');
var req = new expb.Request();
var call = grpcCli.query(req);
console.log('Network call is made');
call.on('data', function (response) {
    console.debug("############## Response size is: ".concat(response.serializeBinary().byteLength));
});
