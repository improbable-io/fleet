load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "coreos.go",
        "fake.go",
        "machine.go",
        "state.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/docker/libcontainer/netlink:go_default_library",
        "//log:go_default_library",
        "//pkg:go_default_library",
        "//unit:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "coreos_test.go",
        "fake_test.go",
        "machine_test.go",
        "state_test.go",
    ],
    library = ":go_default_library",
    deps = ["//pkg:go_default_library"],
)
