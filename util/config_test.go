package util

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
    // Set up some environment variables
    os.Setenv("DB_DRIVER", "test_driver")
    os.Setenv("DB_SOURCE", "test_source")
    os.Setenv("HTTP_SERVER_ADDRESS", "test_http_address")
    os.Setenv("GRPC_SERVER_ADDRESS", "test_grpc_address")

    // Call LoadConfig
    config, err := LoadConfig("../")
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    // Check that the config values match the environment variables
    if config.DBDriver != "test_driver" {
        t.Errorf("expected DB_DRIVER to be 'test_driver', got '%s'", config.DBDriver)
    }
    if config.DBSource != "test_source" {
        t.Errorf("expected DB_SOURCE to be 'test_source', got '%s'", config.DBSource)
    }
    if config.HTTPServerAddress != "test_http_address" {
        t.Errorf("expected SERVER_ADDRESS to be 'test_address', got '%s'", config.HTTPServerAddress)
    }
    if config.GRPCServerAddress != "test_grpc_address" {
        t.Errorf("expected GRPC_SERVER_ADDRESS to be 'test_grpc_address', got '%s'", config.GRPCServerAddress)
    }

    // Clean up
    os.Unsetenv("DB_DRIVER")
    os.Unsetenv("DB_SOURCE")
    os.Unsetenv("HTTP_SERVER_ADDRESS")
    os.Unsetenv("GRPC_SERVER_ADDRESS")
    viper.Reset()
}