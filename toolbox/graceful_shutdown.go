package toolbox

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulShutdown handles the graceful shutdown of the server and database.
func GracefulShutdown(server *http.Server, db *sql.DB, timeout time.Duration) error {
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulStop

	// Context with timeout for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error shutting down server: %w", err)
	}
	fmt.Println("Server gracefully stopped")

	// Close the database connection
	if err := db.Close(); err != nil {
		return fmt.Errorf("error closing database connection: %w", err)
	}
	fmt.Println("Database connection gracefully closed")

	return nil
}

// Rest of your main function (database connection, HTTP handlers, etc.)
