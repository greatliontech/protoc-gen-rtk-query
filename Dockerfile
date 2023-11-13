FROM gcr.io/distroless/static-debian12:latest
COPY protoc-gen-rtk-query /app/protoc-gen-rtk-query
ENTRYPOINT ["/app/protoc-gen-rtk-query"]
