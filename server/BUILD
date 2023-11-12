load("@rules_go//go:def.bzl", "go_library", "go_binary", "go_cross_binary")
load("@genrules//google_cloud_run:index.bzl", "deploy")

go_library(
    name = "lib",
    srcs = ["main.go"],
    importpath = "server",
)

go_binary(
    name = "server",
    embed = [":lib"],
)

go_cross_binary(
    name = "server_linux_amd64",
    platform = "@rules_go//go/toolchain:linux_amd64",
    target = ":server",
)

deploy(
    name = "go_service",
    binary = ":server_linux_amd64",
    port = "3333",
)
