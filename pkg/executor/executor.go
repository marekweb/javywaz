package executor

import (
	"bytes"
	"context"
	"os"
	"time"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

// JavyExecutor is responsible for executing a Javy-built WASM module.
type JavyExecutor struct {
	runtime wazero.Runtime
	module  wazero.CompiledModule
}

// NewJavyExecutor reads the WASM binary from disk, instantiates WASI, and compiles the module.
func NewJavyExecutor(
	ctx context.Context,
	pathToWasmBinary string,
) (*JavyExecutor, error) {
	wasmBinary, err := os.ReadFile(pathToWasmBinary)
	if err != nil {
		return nil, err
	}

	runtimeConfig := wazero.
		NewRuntimeConfigInterpreter().
		WithCloseOnContextDone(true)
	runtime := wazero.NewRuntimeWithConfig(ctx, runtimeConfig)

	_, err = wasi_snapshot_preview1.Instantiate(ctx, runtime)
	if err != nil {
		return nil, err
	}

	compiledModule, err := runtime.CompileModule(ctx, wasmBinary)
	if err != nil {
		return nil, err
	}

	return &JavyExecutor{
		runtime: runtime,
		module:  compiledModule,
	}, nil
}

// Close releases resources held by the compiled module.
func (j *JavyExecutor) Close(ctx context.Context) {
	j.module.Close(ctx)
}

// JavyExecutorResult holds the stdout and stderr output from executing the WASM module.
type JavyExecutorResult struct {
	Stdout string
	Stderr string
}

// Execute runs the Javy WASM module with the provided JSON input. The input should be
// a valid JSON string. The module reads from stdin and writes its JSON output to stdout.
func (j *JavyExecutor) Execute(ctx context.Context, jsonInput string) (*JavyExecutorResult, error) {
	executionTimeout := 5 * time.Second

	stdOutBuffer := &bytes.Buffer{}
	stdErrBuffer := &bytes.Buffer{}
	stdInBuffer := bytes.NewBufferString(jsonInput)

	// Module configuration: No additional arguments are needed since
	// the Javy module expects JSON input on stdin.
	config := wazero.NewModuleConfig().
		WithStdout(stdOutBuffer).
		WithStderr(stdErrBuffer).
		WithStdin(stdInBuffer).
		WithName("example")

	executeCtx, cancel := context.WithTimeout(ctx, executionTimeout)
	defer cancel()
	module, err := j.runtime.InstantiateModule(executeCtx, j.module, config)
	if err != nil {
		return nil, err
	}
	defer module.Close(ctx)

	result := &JavyExecutorResult{
		Stdout: stdOutBuffer.String(),
		Stderr: stdErrBuffer.String(),
	}
	return result, nil
}
