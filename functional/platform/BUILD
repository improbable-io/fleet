load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cluster.go",
        "nspawn.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/code.google.com/p/go-uuid/uuid:go_default_library",
        "//Godeps/_workspace/src/github.com/coreos/go-systemd/dbus:go_default_library",
        "//functional/util:go_default_library",
    ],
)
