package(default_visibility = ["//visibility:public"])

load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/johnskopis/repro
# gazelle:go_proto_compilers @io_bazel_rules_go//proto:gogoslick_proto
# gazelle:go_grpc_compilers @io_bazel_rules_go//proto:gogoslick_grpc
gazelle(
    name = "gazelle",
    external = "vendored",
)

# gazelle:exclude dist
