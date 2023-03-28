package main

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/greatliontech/protoc-gen-rtk-query/module"
	pgs "github.com/lyft/protoc-gen-star"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	optional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.
		Init(pgs.DebugEnv("DEBUG_PGV"), pgs.SupportedFeatures(&optional)).
		RegisterModule(module.RTKQuery()).
		RegisterPostProcessor(prettierFmt{}).
		Render()
}

type prettierFmt struct{}

func (p prettierFmt) Match(a pgs.Artifact) bool {
	var n string

	switch a := a.(type) {
	case pgs.GeneratorFile:
		n = a.Name
	case pgs.GeneratorTemplateFile:
		n = a.Name
	case pgs.CustomFile:
		n = a.Name
	case pgs.CustomTemplateFile:
		n = a.Name
	default:
		return false
	}

	return strings.HasSuffix(n, ".ts")
}

func (p prettierFmt) Process(in []byte) ([]byte, error) {
	_, err := exec.LookPath("prettier")
	if err != nil {
		// Prettier is not found, return input as is
		return in, nil
	}

	cmd := exec.Command("prettier", "--parser", "typescript")
	cmd.Stdin = bytes.NewReader(in)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

var _ pgs.PostProcessor = prettierFmt{}
