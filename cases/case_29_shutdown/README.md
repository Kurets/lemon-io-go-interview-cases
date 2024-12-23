# Graceful Shutdown Server

## Task Description

You must implement a server that:
1. **Listens** for incoming connections or requests (e.g., HTTP server).
2. **Gracefully shuts down** upon receiving `SIGINT` or `SIGTERM` signals.
3. **Prints a message** (e.g. "Shutting down gracefully...") before exiting.

```go
package main

func runServer() {
	// Your code here:
	// 1. Start a server (e.g., an http.Server on some port).
	// 2. Listen for OS signals SIGINT/SIGTERM in a separate goroutine or use signal.Notify.
	// 3. When signal is received, shut down the server gracefully (e.g. server.Shutdown(ctx)).
	// 4. Print "Shutting down gracefully" or any final message.
}
```
