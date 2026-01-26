# protoc-gen-rtk-query

> **Note**: This project is archived. See [Connect-Query](https://connectrpc.com/docs/web/query/getting-started/) for an actively maintained alternative.

A protoc plugin that generates [RTK Query](https://redux-toolkit.js.org/rtk-query/overview) (Redux Toolkit Query) API endpoints from gRPC service definitions.

## What it did

- Generated TypeScript RTK Query API slices from `.proto` service definitions
- Worked with [protobuf-ts](https://github.com/timostamm/protobuf-ts) generated types
- Enabled type-safe data fetching with Redux Toolkit from gRPC backends
- Post-processed output with Prettier for consistent formatting

## Why archived

The [Connect](https://connectrpc.com/) ecosystem now provides [Connect-Query](https://connectrpc.com/docs/web/query/getting-started/), which offers:
- First-party support for TanStack Query (React Query)
- Better maintained and more actively developed
- Part of the broader Connect/buf ecosystem

## License

Apache License 2.0
