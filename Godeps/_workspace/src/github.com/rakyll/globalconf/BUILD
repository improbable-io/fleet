load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["globalconf.go"],
    visibility = ["//visibility:public"],
    deps = ["//Godeps/_workspace/src/github.com/rakyll/goini:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["globalconf_test.go"],
    library = ":go_default_library",
)
