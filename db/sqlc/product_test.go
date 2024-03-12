package db

import (
	"context"
	"testing"
)

func TestGetProducts(t *testing.T) {
    ctx := context.Background()

    // // Assuming that your test database is initially empty
    products, err := testQueries.GetProducts(ctx)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(products) != 0 {
        t.Fatalf("expected no products, got %d", len(products))
    }

    // TODO: Add more tests. For example, insert a product into the test database
    // and then check that GetProducts returns that product.
}