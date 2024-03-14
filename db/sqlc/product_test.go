package db

import (
	"context"
	"log"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/losuch/fc-order/util"
	"github.com/stretchr/testify/require"
)


var product Product

// createRandomProduct creates arandom product for testing purposes.
// It generates random values for the product's attributes and inserts the product into the database.
// It returns the created product
func createRandomProduct(t *testing.T) Product {
    arg := CreateProductParams{
        Name:        util.RandomProductName(),
        Description: pgtype.Text{String: util.RandomString(30), Valid: true},
        ImagesUrl:   pgtype.Text{String: util.RandomString(10), Valid: true},
        Price:       pgtype.Numeric{Int: util.RandomBigInt(1, 100), Valid: true},
        Active:      pgtype.Numeric{Int: big.NewInt(1), Valid: true},
        TypeID:      util.RandomInt64(1, 3),
    }

    log.Printf("arg: %v", arg)

    product, err := testQueries.CreateProduct(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, product)
    require.Equal(t, arg.Name, product.Name)

    require.NotZero(t, product.ProductID)
    require.NotZero(t, product.CreatedAt)
    return product
}

// TestCreateProduct is a unit test for the createRandomProduct function.
func TestCreateProduct(t *testing.T) {
    product = createRandomProduct(t)
}

// TestGetProducts is a unit test for the GetActiveProducts function.
func TestGetProducts(t *testing.T) {
    ctx := context.Background()

    // Assuming that your test database is initially empty
    products, err := testQueries.GetActiveProducts(ctx)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    
    //find product with the productID
    found := false
    for _, p := range products {
        if p.ProductID == product.ProductID {
            found = true
            break
        }
    }
    require.True(t, found)
}