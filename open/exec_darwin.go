//go:build darwin
// +build darwin

package open

import (
	"context"
	"os/exec"
)

func open(ctx context.Context, input string) *exec.Cmd {
	return exec.CommandContext(ctx, "open", input)
}

func openWith(ctx context.Context, input string, appName string) *exec.Cmd {
	return exec.CommandContext(ctx, "open", "-a", appName, input)
}
