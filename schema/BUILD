load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "mapper.go",
        "v1-gen.go",
        "v1-json.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/go-systemd/unit:go_default_library",
        "//Godeps/_workspace/src/google.golang.org/api/googleapi:go_default_library",
        "//job:go_default_library",
        "//machine:go_default_library",
        "//unit:go_default_library",
    ],
)
