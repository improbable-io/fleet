load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "api.go",
        "http.go",
        "registry.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/google.golang.org/api/googleapi:go_default_library",
        "//job:go_default_library",
        "//machine:go_default_library",
        "//registry:go_default_library",
        "//schema:go_default_library",
    ],
)
