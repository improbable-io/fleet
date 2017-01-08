load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "call.go",
        "introspect.go",
        "introspectable.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//Godeps/_workspace/src/github.com/godbus/dbus:go_default_library"],
)
