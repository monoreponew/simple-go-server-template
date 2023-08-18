load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.20.5")

http_archive(
    name = "bazel_gazelle",
    sha256 = "29218f8e0cebe583643cbf93cae6f971be8a2484cdcfa1e45057658df8d54002",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()


http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//toolchains/docker:toolchain.bzl",
    docker_toolchain_configure = "toolchain_configure",
)

docker_toolchain_configure(
    name = "docker_config"
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

HERMETIC_CC_TOOLCHAIN_VERSION = "v2.0.0"

http_archive(
    name = "hermetic_cc_toolchain",
    sha256 = "57f03a6c29793e8add7bd64186fc8066d23b5ffd06fe9cc6b0b8c499914d3a65",
    urls = [
        "https://mirror.bazel.build/github.com/uber/hermetic_cc_toolchain/releases/download/{0}/hermetic_cc_toolchain-{0}.tar.gz".format(HERMETIC_CC_TOOLCHAIN_VERSION),
        "https://github.com/uber/hermetic_cc_toolchain/releases/download/{0}/hermetic_cc_toolchain-{0}.tar.gz".format(HERMETIC_CC_TOOLCHAIN_VERSION),
    ],
)

load("@hermetic_cc_toolchain//toolchain:defs.bzl", zig_toolchains = "toolchains")

zig_toolchains()

register_toolchains(
    "@zig_sdk//toolchain:linux_amd64_gnu.2.28",
    "@zig_sdk//toolchain:linux_arm64_gnu.2.28",
    "@zig_sdk//toolchain:darwin_amd64",
    "@zig_sdk//toolchain:darwin_arm64",
    "@zig_sdk//toolchain:windows_amd64",
    "@zig_sdk//toolchain:windows_arm64",
)

http_archive(
    name="genrules_repo",
    urls=[
        "https://github.com/genrules/repo/archive/db55ffef95c491f29e808122d3c1b297b0ca2096.zip",
    ],
    strip_prefix="repo-db55ffef95c491f29e808122d3c1b297b0ca2096",
    sha256="7c9db19e616d7646ea3ee84d2b863a203f80be11ba3d8b6ac17bca3868317e82",
)

load("@genrules_repo//:index.bzl", "repo")

repo(
    name = "genrules_steps",
    repo = "genrules/steps",
    commit = "befb6a3133bc20ae6320d7bbe88c545a637199fe",
    sha = "b28e85eca74619cf9dc88472d314e794cfadf99e4079f6a4b71571c112e5d085",
)

repo(
    name = "genrules_gcloud",
    repo = "genrules/gcloud",
    commit = "0749fdd9ccf11bb2719042bfd6b8725b5a3d4c3a",
    sha = "d4f771cb42812c812972aed024a47e08b2f101add47b5ca3c6b277e5061e2d6c",
)

repo(
    name = "genrules_crane",
    repo = "genrules/crane",
    commit = "e04e1cacba0deccfcc4128c4e6f3a90a8ebd1628",
    sha = "da553feb05a2bae6c04d8ed1a23d3c46d5f092b02fe87daa6fbfdbe493742a3d",
)

load("@genrules_gcloud//:deps.bzl", "gcloud_deps")

gcloud_deps()

load("@genrules_gcloud//:index.bzl", "gcloud_download")

gcloud_download()

load("@genrules_crane//:deps.bzl", "crane_deps")

crane_deps()

load("@genrules_crane//:index.bzl", "crane_download")

crane_download()