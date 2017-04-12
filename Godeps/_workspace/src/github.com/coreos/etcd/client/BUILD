load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "auth_role.go",
        "auth_user.go",
        "client.go",
        "curl.go",
        "discover.go",
        "doc.go",
        "keys.go",
        "members.go",
        "srv.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/etcd/pkg/types:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/net/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "client_test.go",
        "keys_test.go",
        "members_test.go",
        "srv_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//Godeps/_workspace/src/github.com/coreos/etcd/pkg/types:go_default_library",
        "//Godeps/_workspace/src/golang.org/x/net/context:go_default_library",
        "@com_github_coreos_etcd//pkg/testutil:go_default_library",
    ],
)
