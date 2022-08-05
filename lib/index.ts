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

// https://redux-toolkit.js.org/rtk-query/usage/automated-refetching#abstracting-common-providesinvalidates-usage
export function providesList<R extends { id: string | number }[], T extends string>(
  resultsWithIds: R | undefined,
  tagType: T
) {
  return resultsWithIds
    ? [
      { type: tagType, id: 'LIST' },
      ...resultsWithIds.map(({ id }) => ({ type: tagType, id })),
    ]
    : [{ type: tagType, id: 'LIST' }]
}

