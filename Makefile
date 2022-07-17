install: ~/.local/bin/protoc-gen-rtk-query

~/.local/bin/protoc-gen-rtk-query: module/rtkquery.go
	buf generate proto
	go build -o $@ .

gen: test/proto/test.proto
	buf generate --template buf.gen.test.yaml test/proto

clean:
	rm -rf gen
