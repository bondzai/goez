package toolbox

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulShutdown handles the graceful shutdown of the server and any io.Closer (like a database).
func GracefulShutdown(server *http.Server, closer io.Closer, timeout time.Duration) error {
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulStop
	log.Println("Shutdown signal received")

	// Context with timeout for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down server: %w", err)
	}
	log.Println("Server gracefully stopped")

	// Close the provided io.Closer (e.g., database connection)
	if err := closer.Close(); err != nil {
		return fmt.Errorf("error closing resource: %w", err)
	}
	log.Println("Resource gracefully closed")

	return nil
}

// Rest of your main function (database connection, HTTP handlers, etc.)
