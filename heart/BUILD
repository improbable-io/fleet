load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "heart.go",
        "monitor.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//log:go_default_library",
        "//machine:go_default_library",
        "//registry:go_default_library",
    ],
)
