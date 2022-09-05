// package: example
// file: example.proto

var example_pb = require("./example_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ExampleService = (function () {
  function ExampleService() {}
  ExampleService.serviceName = "example.ExampleService";
  return ExampleService;
}());

ExampleService.Query = {
  methodName: "Query",
  service: ExampleService,
  requestStream: false,
  responseStream: true,
  requestType: example_pb.Request,
  responseType: example_pb.Result
};

exports.ExampleService = ExampleService;

function ExampleServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ExampleServiceClient.prototype.query = function query(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(ExampleService.Query, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.ExampleServiceClient = ExampleServiceClient;

