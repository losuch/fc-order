package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/losuch/fc-order/util"
	"github.com/stretchr/testify/require"
)


var film Film

// createRandomFilm creates a random film for testing purposes.
// It generates random values for the film's attributes and inserts the film into the database.
// It returns the created film
func createRandomFilm(t *testing.T) Film {
	arg := CreateFilmParams{
		Name:   util.RandomString(10),
		YtLink: sql.NullString{String: util.RandomString(50), Valid: true},
		Active: "1",
		Type:   sql.NullString{String: util.RandomString(50), Valid: true},
	}

	film, err := testQueries.CreateFilm(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, film)
	require.Equal(t, arg.Name, film.Name)

	require.NotZero(t, film.FilmID)
	require.NotZero(t, film.CreatedAt)
	return film
}

// TestCreateFilm is a unit test for the createRandomFilm function.
func TestCreateFilm(t *testing.T) {
	film = createRandomFilm(t)
}

// TestGetFilm is a unit test for the GetFilm function.
func TestGetFilm(t *testing.T) {
	p, err := testQueries.GetFilm(context.Background(), film.FilmID)
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, film.FilmID, p.FilmID)
	require.Equal(t, film.Name, p.Name)
	require.Equal(t, film.YtLink, p.YtLink)
	require.Equal(t, film.Active, p.Active)
	require.Equal(t, film.Type, p.Type)
}

// TestGetFilmList is a unit test for the ListFilms function.
func TestGetFilmList(t *testing.T) {
	p, err := testQueries.GetFilmList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, p)
}

// TestUpdateFilm is a unit test for the UpdateFilm function.
func TestUpdateFilm(t *testing.T) {
	arg := UpdateFilmParams{
		FilmID: film.FilmID,
		Name:   util.RandomString(10),
		YtLink: sql.NullString{String: util.RandomString(50), Valid: true},
		Active: "1",
		Type:   sql.NullString{String: util.RandomString(50), Valid: true},
	}
	p, err := testQueries.UpdateFilm(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, arg.Name, p.Name)
	require.Equal(t, arg.YtLink, p.YtLink)
	require.Equal(t, arg.Active, p.Active)
	require.Equal(t, arg.Type, p.Type)
}

// TestDeleteFilm is a unit test for the DeleteFilm function.
func TestDeleteFilm(t *testing.T) {
	err := testQueries.DeleteFilm(context.Background(), film.FilmID)
	require.NoError(t, err)
}