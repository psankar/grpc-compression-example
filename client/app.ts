import * as exsvc from 'grpc-compression-example/example_pb_service';
import * as expb from 'grpc-compression-example/example_pb';
import * as grpc from 'grpc';

const grpcCli = new exsvc.ExampleServiceClient("localhost:8080",
    grpc.credentials.createInsecure(),
    // { 'grpc.default_compression_algorithm': 3, 'grpc.default_compression_level': 3 },
);

console.log('grpc client is created');

let req = new expb.Request();
const call = grpcCli.query(req);
console.log('Network call is made');

call.on('data', (response: expb.Result) => {
    console.debug(`############## Response size is: ${response.serializeBinary().byteLength}`);
});

