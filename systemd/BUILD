load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["manager.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/go-systemd/dbus:go_default_library",
        "//log:go_default_library",
        "//pkg:go_default_library",
        "//unit:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["manager_test.go"],
    library = ":go_default_library",
)
