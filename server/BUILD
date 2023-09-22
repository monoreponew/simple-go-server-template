load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@genrules//google_cloud_run:index.bzl", "deploy")

go_library(
    name = "lib",
    srcs = ["main.go"],
    importpath = "foo",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "go",
    embed = [":lib"],
    visibility = ["//visibility:public"],
)

go_image(
    name = "go_image",
    binary = ":go",
)

deploy(
    name="go_service",
    target=":server_image",
)
