/*
Open a file, directory, or URI using the OS's default
application for that object type.  Optionally, you can
specify an application to use.

This is a proxy for the following commands:

	        OSX: "open"
	    Windows: "start"
	Linux/Other: "xdg-open"

This is a golang port of the node.js module: https://github.com/pwnall/node-open
*/
package open

import "context"

/*
Open a file, directory, or URI using the OS's default
application for that object type. Wait for the open
command to complete.
*/
func Run(ctx context.Context, input string) error {
	return open(ctx, input).Run()
}

/*
Open a file, directory, or URI using the OS's default
application for that object type. Don't wait for the
open command to complete.
*/
func Start(ctx context.Context, input string) error {
	return open(ctx, input).Start()
}

/*
Open a file, directory, or URI using the specified application.
Wait for the open command to complete.
*/
func RunWith(ctx context.Context, input string, appName string) error {
	return openWith(ctx, input, appName).Run()
}

/*
Open a file, directory, or URI using the specified application.
Don't wait for the open command to complete.
*/
func StartWith(ctx context.Context, input string, appName string) error {
	return openWith(ctx, input, appName).Start()
}
