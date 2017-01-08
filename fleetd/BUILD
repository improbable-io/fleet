load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["fleetd.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//Godeps/_workspace/src/github.com/rakyll/globalconf:go_default_library",
        "//agent:go_default_library",
        "//config:go_default_library",
        "//log:go_default_library",
        "//pkg:go_default_library",
        "//registry:go_default_library",
        "//server:go_default_library",
        "//version:go_default_library",
    ],
)

go_binary(
    name = "fleetd",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
