//go:build !windows && !darwin
// +build !windows,!darwin

package open

import (
	"context"
	"os/exec"
)

// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-open/
// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-mime/

func open(ctx context.Context, input string) *exec.Cmd {
	return exec.CommandContext(ctx, "xdg-open", input)
}

func openWith(ctx context.Context, input string, appName string) *exec.Cmd {
	return exec.CommandContext(ctx, appName, input)
}
