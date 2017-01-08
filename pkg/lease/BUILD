load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "etcd.go",
        "interface.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/etcd/client:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/net/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["etcd_test.go"],
    library = ":go_default_library",
    deps = ["//Godeps/_workspace/src/github.com/coreos/etcd/client:go_default_library"],
)
