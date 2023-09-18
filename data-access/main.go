package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	godotenv.Load(".env")
	connString := fmt.Sprintf("postgres://%s:%s@localhost:5432/postgres", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))

	cfg, err := pgx.ParseConfig(connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse database conn string %v\n", err)
		os.Exit(5)
	}

	conn, err := pgx.ConnectConfig(context.Background(), cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database %v\n", err)
		os.Exit(5)
	}

	albums, err := albumsByArtist(conn, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found %v\n", albums)

	albumId, err := addAlbum(conn, Album{
		ID:     0,
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Id of added album: %v\n", albumId)
}

func albumsByArtist(conn *pgx.Conn, name string) ([]Album, error) {
	var albums []Album
	rows, err := conn.Query(context.Background(), "select * from album where artist = $1;", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func addAlbum(conn *pgx.Conn, album Album) (int64, error) {
	var id int64
	err := conn.QueryRow(context.Background(), "insert into album (title, artist, price) values ($1, $2, $3) returning id", album.Title, album.Artist, album.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
