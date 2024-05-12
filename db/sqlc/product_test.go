package db

import (
	"context"
	"database/sql"
	"strconv"
	"testing"

	"github.com/losuch/fc-order/util"
	"github.com/stretchr/testify/require"
)

var product Product

// createRandomProduct creates a random product for testing purposes.
// It generates random values for the product's attributes and inserts the product into the database.
// It returns the created product
func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Name:        util.RandomString(10),
		Description: sql.NullString{String: util.RandomString(50), Valid: true},
		ImagesUrl:   sql.NullString{String: util.RandomString(50), Valid: true},
		Active:      "1",
		TypeID:      int32(util.RandomInt64(0, 2)),
		Price:       strconv.FormatInt(util.RandomInt64(1000, 10000), 10),
	}

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

// TestGetProduct is a unit test for the GetProduct function.
func TestGetProduct(t *testing.T) {
	p, err := testQueries.GetProduct(context.Background(), product.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, product.ProductID, p.ProductID)
	require.Equal(t, product.Name, p.Name)
	require.Equal(t, product.Description, p.Description)
	require.Equal(t, product.ImagesUrl, p.ImagesUrl)
	require.Equal(t, product.Price, p.Price)
	require.Equal(t, product.Active, p.Active)
	require.Equal(t, product.TypeID, p.TypeID)
}

// TestGetProductList is a unit test for the ListProducts function.
func TestGetProductList(t *testing.T) {

	products, err := testQueries.GetProductList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, products)

	// Check if the created product is in the list
	found := false
	for _, p := range products {
		if p.ProductID == product.ProductID {
			found = true
			break
		}
	}
	require.True(t, found)
}

// TestUpdateProduct is a unit test for the UpdateProduct function.
func TestUpdateProduct(t *testing.T) {
	arg := UpdateProductParams{
		ProductID:   product.ProductID,
		Name:        util.RandomString(10),
		Description: sql.NullString{String: util.RandomString(50), Valid: true},
		ImagesUrl:   sql.NullString{String: util.RandomString(50), Valid: true},
		Active:      "1",
		TypeID:      int32(util.RandomInt64(0, 2)),
		Price:       strconv.FormatInt(util.RandomInt64(1000, 10000), 10),
	}

	p, err := testQueries.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, arg.Name, p.Name)
	require.Equal(t, arg.Description, p.Description)
	require.Equal(t, arg.ImagesUrl, p.ImagesUrl)
	require.Equal(t, arg.Price, p.Price)
	require.Equal(t, arg.Active, p.Active)
	require.Equal(t, arg.TypeID, p.TypeID)
}

// TestDeleteProduct is a unit test for the DeleteProduct function.
func TestDeleteProduct(t *testing.T) {
	err := testQueries.DeleteProduct(context.Background(), product.ProductID)
	require.NoError(t, err)

	// Check if the product is deleted
	p, err := testQueries.GetProduct(context.Background(), product.ProductID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, p)
}