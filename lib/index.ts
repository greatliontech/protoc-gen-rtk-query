import type { BaseQueryFn } from '@reduxjs/toolkit/query'
import type { RpcMetadata, RpcStatus, UnaryCall } from "@protobuf-ts/runtime-rpc";

export type GrpcBaseQueryMeta = {
  status: RpcStatus
  headers: RpcMetadata
  trailers: RpcMetadata
}

export type WithMetadata<T> = T & {
  grpcQueryMetadata: GrpcBaseQueryMeta;
}

type PublicOf<T> = { [K in keyof T]: T[K] };

export const grpcBaseQuery =
  (): BaseQueryFn<
    PublicOf<UnaryCall>,
    unknown,
    unknown,
    unknown,
    GrpcBaseQueryMeta
  > =>
    async (call) => {
      try {
        console.log(`### calling method "${call.method.name}"...`)

        const headers = await call.headers;
        const response = await call.response;
        const status = await call.status;
        const trailers = await call.trailers;

        return {
          data: response,
          meta: {
            status,
            headers,
            trailers,
          },
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
