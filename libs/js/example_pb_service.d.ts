// package: example
// file: example.proto

import * as example_pb from "./example_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ExampleServiceQuery = {
  readonly methodName: string;
  readonly service: typeof ExampleService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof example_pb.Request;
  readonly responseType: typeof example_pb.Result;
};

export class ExampleService {
  static readonly serviceName: string;
  static readonly Query: ExampleServiceQuery;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ExampleServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  query(requestMessage: example_pb.Request, metadata?: grpc.Metadata): ResponseStream<example_pb.Result>;
}

