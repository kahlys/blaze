# Setup Golang

## Basic setup

Add the following lines to your `MODULE.bazel` file:

```starlark
bazel_dep(name = "rules_go", version = "0.39.1")
bazel_dep(name = "gazelle", version = "0.31.0")
```

Check https://registry.bazel.build/ for the latest versions.

## Go sdk

If you want to use a specific Go sdk version, you can add the following lines to your `MODULE.bazel` file:

```starlark
go_sdk = use_extension("@rules_go//go:extensions.bzl", "go_sdk")

# Download an SDK for the host OS & architecture as well as common remote execution platforms.
go_sdk.download(version = "1.23.1")

# Alternately, download an SDK for a fixed OS/architecture.
go_sdk.download(
    version = "1.23.1",
    goarch = "amd64",
    goos = "linux",
)

# Register the Go SDK installed on the host.
go_sdk.host()
```

## Generate BUILD files

Add the following to your top-level BUILD file:

```starlark
load("@gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    mode = "fix",
)
```

The `go.mod` file should be in the same directory as the `BUILD` file. If not, add the following to your `MODULE.bazel` file:

```starlark
# gazelle:prefix github.com/example/project
```

Generate the BUILD files:

```bash
bazel run //:gazelle
```

If needed you can create the go.mod file

```bash
bazel run @rules_go//go mod init github.com/example/project
```

Register external dependencies by adding these lines to your `MODULE.bazel` file:

```starlark
go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")
```

Then run:

```bash
bazel mod tidy
```

## Sources

- https://github.com/bazelbuild/rules_go/blob/master/docs/go/core/bzlmod.md