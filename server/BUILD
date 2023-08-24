load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@genrules//gcloud:index.bzl", "gcloud_run_deploy", "gcloud_services_enable", "gcloud_auth", "gcloud_auth_print_access_token")
load("@genrules//run_all:index.bzl", "run_all")
load("@genrules//run_if:index.bzl", "run_if")
load("@genrules//crane:index.bzl", "crane_auth_login", "crane_push")

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

run_all(
    name="deploy",
    steps = [
        ":auth_if_needed",
        ":access_token",
        ":gcr_login",
        ":push_image",
        ":enable_cloud_run",
        ":deploy_image",
    ],
)

gcloud_auth_print_access_token(
    name = "access_token",
    out = "~/access_token",
)

crane_auth_login(
    name = "gcr_login",
    registry = "gcr.io",
    user = "oauth2accesstoken",
    password = "$(cat ~/access_token)",
)

crane_push(
    name = "push_image",
    target = ":go_image.tar",
    image = "gcr.io/$GCP_PROJECT/go_image:latest"
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

run_if(
    name = "auth_if_needed",
    not_empty = "$CLOUDSDK_AUTH_REFRESH_TOKEN",
    then = ":auth",
)

gcloud_auth(
    name = "auth"
)
