load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "auth.go",
        "auth_external.go",
        "auth_sha1.go",
        "call.go",
        "conn.go",
        "conn_darwin.go",
        "dbus.go",
        "decoder.go",
        "doc.go",
        "encoder.go",
        "export.go",
        "homedir.go",
        "homedir_dynamic.go",
        "message.go",
        "sig.go",
        "transport_darwin.go",
        "transport_generic.go",
        "transport_unix.go",
        "variant.go",
        "variant_lexer.go",
        "variant_parser.go",
    ],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "conn_test.go",
        "examples_test.go",
        "proto_test.go",
        "sig_test.go",
        "transport_unix_test.go",
        "variant_test.go",
    ],
    library = ":go_default_library",
)
