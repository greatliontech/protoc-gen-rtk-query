package main

import (
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
	"github.com/thegrumpylion/protoc-gen-rtk-query/module"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	optional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.
		Init(pgs.DebugEnv("DEBUG_PGV"), pgs.SupportedFeatures(&optional)).
		RegisterModule(module.RTKQuery()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
