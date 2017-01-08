load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["job.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg:go_default_library",
        "//unit:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["job_test.go"],
    library = ":go_default_library",
    deps = [
        "//pkg:go_default_library",
        "//unit:go_default_library",
    ],
)
