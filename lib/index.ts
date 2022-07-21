import type { BaseQueryFn } from '@reduxjs/toolkit/query'
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";

export const grpcBaseQuery =
  (): BaseQueryFn<
    UnaryCall,
    unknown,
    unknown
  > =>
    async (call) => {
      try {
        console.log(`### calling method "${call.method.name}"...`)

        const headers = await call.headers;
        console.log("got response headers: ", headers)

        const response = await call.response;
        console.log("got response message: ", response)

        const status = await call.status;
        console.log("got status: ", status)

        const trailers = await call.trailers;
        console.log("got trailers: ", trailers)

        console.log();

        return {
          data: response
        }
      } catch (grpcError) {
        return {
          error: {
            data: grpcError
          },
        }
      }
    }

