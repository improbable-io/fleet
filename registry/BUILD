load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "etcd.go",
        "event.go",
        "fake.go",
        "interface.go",
        "job.go",
        "job_state.go",
        "machine.go",
        "registry.go",
        "unit.go",
        "unit_state.go",
        "version.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/etcd/client:go_default_library",
        "//Godeps/_workspace/src/github.com/coreos/go-semver/semver:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/net/context:go_default_library",
        "//job:go_default_library",
        "//log:go_default_library",
        "//machine:go_default_library",
        "//metrics:go_default_library",
        "//pkg:go_default_library",
        "//pkg/lease:go_default_library",
        "//unit:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "event_test.go",
        "fake_test.go",
        "unit_state_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/etcd/client:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/net/context:go_default_library",
        "//job:go_default_library",
        "//machine:go_default_library",
        "//pkg:go_default_library",
        "//unit:go_default_library",
    ],
)
