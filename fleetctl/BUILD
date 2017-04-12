load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cat.go",
        "destroy.go",
        "fdforward.go",
        "fleetctl.go",
        "help.go",
        "journal.go",
        "list_machines.go",
        "list_unit_files.go",
        "list_units.go",
        "load.go",
        "ssh.go",
        "start.go",
        "status.go",
        "stop.go",
        "submit.go",
        "unload.go",
        "verify.go",
        "version.go",
    ],
    visibility = ["//visibility:private"],
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/etcd/client:go_default_library",
        "//api:go_default_library",
        "//client:go_default_library",
        "//job:go_default_library",
        "//log:go_default_library",
        "//machine:go_default_library",
        "//pkg:go_default_library",
        "//registry:go_default_library",
        "//schema:go_default_library",
        "//ssh:go_default_library",
        "//unit:go_default_library",
        "//version:go_default_library",
    ],
)

go_binary(
    name = "fleetctl",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "fleetctl_test.go",
        "list_machines_test.go",
        "list_unit_files_test.go",
        "list_units_test.go",
        "ssh_test.go",
        "start_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/go-semver/semver:go_default_library",
        "//client:go_default_library",
        "//job:go_default_library",
        "//machine:go_default_library",
        "//registry:go_default_library",
        "//schema:go_default_library",
        "//unit:go_default_library",
        "//version:go_default_library",
    ],
)
