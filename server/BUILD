load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@genrules_gcloud//:index.bzl", "gcloud_run_deploy", "gcloud_services_enable")
load("@genrules_steps//:index.bzl", "steps")
load("@genrules_crane//:index.bzl", "crane_auth_login", "crane_push")

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

gcloud_services_enable(
    name = "enable_cloud_run",
    service = "run.googleapis.com",
    project = "$GCP_PROJECT",
)

gcloud_run_deploy(
    name = "deploy_image",
    service = "test-$RANDOM",
    image = "gcr.io/$GCP_PROJECT/go_image:latest",
    port = "3333",
    region = "us-west1",
    project = "$GCP_PROJECT",
    allow_unauthenticated = True,
)

crane_auth_login(
    name = "gcr_login",
    registry = "gcr.io",
    user = "oauth2accesstoken",
    password = "$CLOUDSDK_AUTH_ACCESS_TOKEN",
)

crane_push(
    name = "push_image",
    target = ":go_image.tar",
    image = "gcr.io/$GCP_PROJECT/go_image:latest"
)

steps(
    name="deploy",
    steps = [
        ":gcr_login",
        ":push_image",
        ":enable_cloud_run",
        ":deploy_image",
    ],
)
