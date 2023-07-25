example/app/src/gen/todo.api.ts example/service/gen/todo.pb.go: example/proto/todo.proto install lib/dist/index.d.ts
	buf generate --template buf.gen.example.yaml example/proto

install: ~/.local/bin/protoc-gen-rtk-query

~/.local/bin/protoc-gen-rtk-query: main.go module/funcs.go module/imports.go module/params.go module/rtkquery.go module/templates.go proto/rtkquery/rtkquery.pb.go
	go build -o $@ .

proto/rtkquery/rtkquery.pb.go: proto/rtkquery/rtkquery.proto
	buf generate proto

lib/dist/index.d.ts: lib/index.ts
	cd lib && npm i && tsc

clean:
	rm -rf example/app/src/gen example/service/gen
