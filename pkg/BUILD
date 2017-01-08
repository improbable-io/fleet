load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "args.go",
        "backoff.go",
        "filepath.go",
        "filesystem.go",
        "flag.go",
        "http.go",
        "reconcile.go",
        "set.go",
        "tls.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/jonboulle/clockwork:go_default_library",
        "//log:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "args_test.go",
        "backoff_test.go",
        "filepath_test.go",
        "filesystem_test.go",
        "flag_test.go",
        "reconcile_test.go",
        "set_test.go",
        "tls_test.go",
    ],
    library = ":go_default_library",
    deps = ["//Godeps/_workspace/src/github.com/jonboulle/clockwork:go_default_library"],
)
