load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "known_hosts.go",
        "match.go",
        "proxy.go",
        "ssh.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/golang.org/x/crypto/ssh:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/crypto/ssh/agent:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/crypto/ssh/terminal:go_default_library",
        "//log:go_default_library",
        "//pkg:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "known_hosts_test.go",
        "match_test.go",
    ],
    library = ":go_default_library",
    deps = ["//Godeps/_workspace/src/golang.org/x/crypto/ssh:go_default_library"],
)
