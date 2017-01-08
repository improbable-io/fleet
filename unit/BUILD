load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "fake.go",
        "generator.go",
        "manager.go",
        "unit.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/go-systemd/unit:go_default_library",
        "//log:go_default_library",
        "//pkg:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "fake_test.go",
        "generator_test.go",
        "unit_test.go",
    ],
    library = ":go_default_library",
    deps = ["//pkg:go_default_library"],
)
