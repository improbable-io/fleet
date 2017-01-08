load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "engine.go",
        "reconciler.go",
        "scheduler.go",
        "state.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//agent:go_default_library",
        "//job:go_default_library",
        "//log:go_default_library",
        "//machine:go_default_library",
        "//metrics:go_default_library",
        "//pkg:go_default_library",
        "//pkg/lease:go_default_library",
        "//registry:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "engine_test.go",
        "reconciler_test.go",
        "scheduler_test.go",
        "state_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//agent:go_default_library",
        "//job:go_default_library",
        "//machine:go_default_library",
        "//registry:go_default_library",
        "//unit:go_default_library",
    ],
)
