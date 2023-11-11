load("@rules_go//go:def.bzl", "go_binary", "go_library")
#load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@genrules//google_cloud_run:index.bzl", "deploy")

go_library(
    name = "lib",
    srcs = ["main.go"],
    importpath = "foo",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "server",
    embed = [":lib"],
    visibility = ["//visibility:public"],
)

# crane(
#     name = "go_image",
#     # command = "append -f $< -b gcr.io/distroless/static-debian11 -t go_image -o $@",
#     command = "append -f $< -b debian -t go_image -o $@",
#     compile = True,
#     deps = [":tar"],
# )

# crane(
#     name = "mutate",
#     command = "mutate -a --cmd=./server/server_/server -o $@ $<",
#     compile = True,
#     deps = [":go_image"],
# )


#go_image(
#    name = "go_image",
#    binary = ":go",
#)

deploy(
    name = "go_service",
    binary = ":server",
    port = "3333",
)
