example/app/src/gen/todo.api.ts example/service/gen/todo.pb.go: example/proto/todo.proto install lib/dist/index.d.ts
	buf generate --template buf.gen.example.yaml example/proto

install: ~/.local/bin/protoc-gen-rtk-query

~/.local/bin/protoc-gen-rtk-query: module/rtkquery.go proto/rtkquery/rtkquery.pb.go
	go build -o $@ .

proto/rtkquery/rtkquery.pb.go: proto/rtkquery/rtkquery.proto
	buf generate proto

lib/dist/index.d.ts: lib/index.ts
	cd lib && tsc

clean:
	rm -rf example/app/src/gen example/service/gen
