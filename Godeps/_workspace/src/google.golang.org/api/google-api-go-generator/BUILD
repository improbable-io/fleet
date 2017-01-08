load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["gen.go"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "google-api-go-generator",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["gen_test.go"],
    library = ":go_default_library",
)
