# Javy Wazero example

This is an experiment using Javy and WASM

## Building WASM modules with Javy

Javy is a tool that creates WASM (WASI) modules from JavaScript code. It's a way to package some JavaScript code to run in WASM runtimes.

The [build.sh](build.sh) script uses Javy to compile [example.js](example.js) into [example.wasm](example.wasm).

## Running WASM modules with Wazero

Meanwhile, Wazero lets us run WASM (and WASI) modules in Go, so it's a WASM runtime for Go.

The [cmd/run/main.go](cmd/run/main.go) file uses Wazero to execute the [example.wasm](example.wasm) file. It uses the [`runjavy.NewJavyExecutor`](pkg/runjavy/run.go) from [pkg/runjavy/run.go](pkg/runjavy/run.go) to set up the WASM runtime and execute the module.

### What does the trivial example JS code do?

Thje JS code in [example.js](example.js) reads a JSON object from stdin, adds one to the `n` field, appends "!" to the `bar` field, and writes the modified JSON to stdout.


### What's next to use Javy and Wazero for something real?

Note that the convention used by Javy is that the JS code should communicate using a JSON object on stdin and stdout. I believe it expects to read one single JSON object from stdin when it runs, and it's supposed to write one single JSON object to stdout when it finishes. This may be more flexible, I'm not sure. But that can be enough for most uses, because one invocation of the WASM module can be thought of as one function call -- I think it there's the theoretical possibility of a WASM module being used as a long-lived process, but the simpler user case is just a short invocation at a time.



