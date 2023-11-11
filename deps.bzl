load("@genrules//crane:index.bzl", "crane_download")
load("@genrules//gcloud:index.bzl", "gcloud_download")

def _deps(_ctx):
    gcloud_download()
    crane_download()


deps = module_extension(
    implementation = _deps,
)