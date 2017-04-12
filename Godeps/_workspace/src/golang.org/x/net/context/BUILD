load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["context.go"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["context_test.go"],
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = ["withtimeout_test.go"],
    deps = ["//Godeps/_workspace/src/golang.org/x/net/context:go_default_library"],
)
