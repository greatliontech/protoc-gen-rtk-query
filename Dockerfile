FROM gcr.io/distroless/static-debian12:latest
COPY mindns /app/protoc-gen-rtk-query
ENTRYPOINT ["/app/protoc-gen-rtk-query"]
