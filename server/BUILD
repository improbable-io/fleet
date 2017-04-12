load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["server.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/etcd/client:go_default_library",
        "//Godeps/_workspace/src/github.com/coreos/go-systemd/activation:go_default_library",
        "//agent:go_default_library",
        "//api:go_default_library",
        "//config:go_default_library",
        "//engine:go_default_library",
        "//heart:go_default_library",
        "//log:go_default_library",
        "//machine:go_default_library",
        "//pkg:go_default_library",
        "//pkg/lease:go_default_library",
        "//registry:go_default_library",
        "//systemd:go_default_library",
        "//unit:go_default_library",
        "//version:go_default_library",
    ],
)
