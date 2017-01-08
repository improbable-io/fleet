load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "discovery.go",
        "machines.go",
        "mux.go",
        "pagination.go",
        "path.go",
        "serialization.go",
        "server.go",
        "state.go",
        "units.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//client:go_default_library",
        "//job:go_default_library",
        "//log:go_default_library",
        "//machine:go_default_library",
        "//pkg:go_default_library",
        "//registry:go_default_library",
        "//schema:go_default_library",
        "//unit:go_default_library",
        "//version:go_default_library",
        "@com_github_prometheus_client_golang//prometheus:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "deserialization_test.go",
        "discovery_test.go",
        "machines_test.go",
        "mux_test.go",
        "pagination_test.go",
        "path_test.go",
        "serialization_test.go",
        "state_test.go",
        "units_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//client:go_default_library",
        "//job:go_default_library",
        "//machine:go_default_library",
        "//registry:go_default_library",
        "//schema:go_default_library",
        "//unit:go_default_library",
        "//version:go_default_library",
    ],
)
