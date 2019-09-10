package videos

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"../../cmd"
)

// Video model.
type Video struct {
	VideoURL string `firestore:"videoUrl"`
}

// GetVideos take all videos list.
func GetVideos(w http.ResponseWriter, r *http.Request) {

	var ctx = context.Background()
	app := cmd.FBInstance()
	db, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return
	}

	docsnap, err := db.Collection("videos").Documents(ctx).Next()
	if err != nil {
	}

	var v Video
	if err := docsnap.DataTo(&v); err != nil {
		println(err)
	}

	// fmt.Println(v)

	json.NewEncoder(w).Encode(v)

	// println(result.)
	// db, err := app.Database(ctx)
	// if err != nil {
	// 	log.Fatalf("error getting Auth client: %v\n", err)
	// 	return
	// }

	// ref := db.NewRef("test/testingo")
	// results, err := ref.OrderByKey().GetOrdered(ctx)

	// books := []Book{}
	// if err != nil {
	// 	log.Fatalln("Error querying database:", err)
	// }
	// for _, r := range results {
	// 	// data := r.Unmarshal(Book)
	// 	// books = append(books)
	// 	var book Book

	// 	if err := r.Unmarshal(&book); err != nil {
	// 		print(err)
	// 	}
	// 	books = append(books, book)

	// }
}
