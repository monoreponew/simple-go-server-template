load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

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

sh_binary(
    name = "deploy",
    srcs = ["deploy.sh"],
    data = ["@gcloud", "@crane", ":go_image.tar"],
)