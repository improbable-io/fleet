load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "agent.go",
        "cache.go",
        "reconcile.go",
        "state.go",
        "task.go",
        "unit_state.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/jonboulle/clockwork:go_default_library",
        "//job:go_default_library",
        "//log:go_default_library",
        "//machine:go_default_library",
        "//pkg:go_default_library",
        "//registry:go_default_library",
        "//unit:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "agent_test.go",
        "reconcile_test.go",
        "state_test.go",
        "task_test.go",
        "unit_state_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//Godeps/_workspace/src/github.com/jonboulle/clockwork:go_default_library",
        "//job:go_default_library",
        "//machine:go_default_library",
        "//registry:go_default_library",
        "//unit:go_default_library",
    ],
)
