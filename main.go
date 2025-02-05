package main

import (
	"os"

	_ "golang.org/x/oauth2"
)

// DefaultBaseImage for non-scratch builds.
// renovate: datasource=docker versioning=docker depName=alpine
const DefaultBaseImage = "docker.io/alpine:3.21"

// DefaultDirMode is UNIX file mode for mkdir.
const DefaultDirMode os.FileMode = 0o755

// DefaultPath is default value for PATH environment variable.
const DefaultPath = "/bin:/usr/bin:/sbin:/usr/sbin"

// PkgYaml is the filename of 'pkg.yaml'.
const PkgYaml = "pkg.yaml"

// VarsYaml is the filename of 'vars.yaml'.
const VarsYaml = "vars.yaml"

// Pkgfile is the filename of 'Pkgfile'.
const Pkgfile = "Pkgfile"

// StageXBusyboxImage is the image name for busybox from stageX.
// renovate: datasource=docker versioning=docker depName=ghcr.io/siderolabs/stagex/core-busybox
const StageXBusyboxImage = "ghcr.io/siderolabs/stagex/core-busybox@sha256:c0b551b47d8f1ac2fd5f4712eafddb8717e6e563a47203e02f94f944f64c18b2"

func main() {
}
