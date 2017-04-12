load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "forward.go",
        "keyring.go",
        "server.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//Godeps/_workspace/src/golang.org/x/crypto/ssh:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "client_test.go",
        "server_test.go",
        "testdata_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//Godeps/_workspace/src/golang.org/x/crypto/ssh:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/crypto/ssh/testdata:go_default_library",
    ],
)
