load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "agent_unix_test.go",
        "cert_test.go",
        "forward_unix_test.go",
        "session_test.go",
        "tcpip_test.go",
        "test_unix_test.go",
        "testdata_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//Godeps/_workspace/src/golang.org/x/crypto/ssh:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/crypto/ssh/agent:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/crypto/ssh/testdata:go_default_library",
    ],
)
