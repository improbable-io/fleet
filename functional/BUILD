load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "client_test.go",
        "cluster_test.go",
        "fleetctl_test.go",
        "node_test.go",
        "scheduling_test.go",
        "systemd_test.go",
        "unit_action_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//functional/platform:go_default_library",
        "//functional/util:go_default_library",
        "//job:go_default_library",
        "//systemd:go_default_library",
        "//unit:go_default_library",
        "//version:go_default_library",
    ],
)
