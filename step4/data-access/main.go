package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// export DBPASSWORD=password
// export DBUSER=user
var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", os.Getenv("DBPASSWORD"), "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "localhost", "the database server")
	user          = flag.String("user", os.Getenv("DBUSER"), "the database user")
)

func main() {

	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;Database=recordings", *server, *user, *password, *port)

	if *debug {
		fmt.Printf(" connectionString:%s \n", connString)
	}

	db, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open Connection Failed", err.Error())
	}

	fmt.Println(db.Stats())

	// defer conn.Close()

	// stmt, err := conn.Prepare("Select getdate()")

	// if err != nil {
	// 	log.Fatal("Prepare Failed", err.Error())
	// }

	// defer stmt.Close()

	// row := stmt.QueryRow()
	// var dateTime time.Time
	// err = row.Scan(&dateTime)
	// if err != nil {
	// 	log.Fatal("Scan failed:", err.Error())
	// }
	// fmt.Printf("dateTime:%d\n, %v", dateTime.Format(time.RFC3339), time.Now())

	albums, err := albumsByArtist("John Coltrane", db)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums Found %v \n", albums)

	album, err := albumById(2, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album Found: %v \n", album)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	fmt.Printf("bye\n")

}

func albumsByArtist(name string, db *sql.DB) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	fmt.Println(db.Stats())

	var albums []Album

	rows, err := db.Query("SELECT * FROM dbo.album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func albumById(id int64, db *sql.DB) (Album, error) {

	var alb Album
	row := db.QueryRow("SELECT * FROM album where id = ?", id)

	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumById %d: %v", id, err)
	}
	return alb, nil

}

func addAlbum(album Album, db *sql.DB) (int64, error) {

	var lastInsertId2 int64
	err := db.QueryRow("INSERT INTO album(title, artist, price) OUTPUT inserted.id values (?, ?, ?)", album.Title, album.Artist, album.Price).Scan(&lastInsertId2)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return lastInsertId2, nil
}
