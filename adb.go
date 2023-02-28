package main

import (
	"context"
	"log"
	"time"

	adb "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func dbConn() (adb.Database, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		return nil, err
	}

	client, err := adb.NewClient(adb.ClientConfig{
		Connection:     conn,
		Authentication: adb.BasicAuthentication("root", "rootpassword"),
	})
	if err != nil {
		return nil, err
	}

	db, err := client.Database(nil, "_system")
	if err != nil {
		return nil, err
	}

	return db, nil
}

type User struct {
	Lo int `json:"lo"`
	Go int `json:"go"`
}

func getVerticesCountInSameDepth(ctx context.Context, db adb.Database, rootUser string) int {
	querystring := "FOR v IN 1..1 OUTBOUND @coll GRAPH 'minions' OPTIONS { bfs: true } RETURN v"
	bindVars := map[string]interface{}{
		"coll": "partners/" + rootUser,
	}

	cursor, err := db.Query(ctx, querystring, bindVars)
	if err != nil {
		log.Fatalf("next vertex query failed: %v", err)
	}
	defer cursor.Close()

	var counter int
	for {
		var doc User
		_, err := cursor.ReadDocument(ctx, &doc)

		if adb.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Fatalf("Doc returned: %v", err)
		}
		counter++
	}

	return counter
}

func getNextVertex(ctx context.Context, db adb.Database, prevUser string, usersInSameDepth []string) (user string, userLo int) {
	//	log.Println("head", prevUser, "not in", usersInSameDepth)
	querystring := "FOR v IN 1..1 OUTBOUND @coll GRAPH 'minions' OPTIONS { bfs: true } FILTER v._key NOT IN @users RETURN v"
	bindVars := map[string]interface{}{
		"coll":  "partners/" + prevUser,
		"users": usersInSameDepth,
	}

	cursor, err := db.Query(ctx, querystring, bindVars)
	if err != nil {
		log.Fatalf("next vertex query failed: %v", err)
	}
	defer cursor.Close()

	var doc User
	var metadata adb.DocumentMeta
	metadata, err = cursor.ReadDocument(ctx, &doc)

	if adb.IsNoMoreDocuments(err) {
		return
	} else if err != nil {
		log.Fatalf("Doc returned: %v", err)
	}

	user, userLo = metadata.Key, doc.Lo
	return
}

func getHeadVertex(ctx context.Context, db adb.Database, curr string) string {
	querystring := "FOR v IN 1..1 INBOUND @coll GRAPH 'minions' RETURN v"
	bindVars := map[string]interface{}{
		"coll": "partners/" + curr,
	}

	cursor, err := db.Query(ctx, querystring, bindVars)
	if err != nil {
		log.Fatalf("next vertex query failed: %v", err)
	}
	defer cursor.Close()

	var doc User
	var metadata adb.DocumentMeta
	metadata, err = cursor.ReadDocument(ctx, &doc)

	if adb.IsNoMoreDocuments(err) {
		return ""
	} else if err != nil {
		log.Fatalf("Doc returned: %v", err)
	}

	return metadata.Key
}

func compressionTreversal(ctx context.Context, db adb.Database, user string, lo *int) {
	querystring := "FOR v,e,p IN 1..1 OUTBOUND @coll GRAPH 'minions' RETURN v"
	bindVars := map[string]interface{}{
		"coll": "partners/" + user,
	}

	cursor, err := db.Query(ctx, querystring, bindVars)
	if err != nil {
		log.Fatalf("compression query failed: %v", err)
	}
	defer cursor.Close()

	for {
		var doc User
		var metadata adb.DocumentMeta
		metadata, err = cursor.ReadDocument(ctx, &doc)
		if adb.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Fatalf("Doc returned: %v", err)
		}
		if doc.Lo != 0 {
			log.Println("user:", metadata.Key, "personal volume =", *lo, "+", doc.Lo)
			*lo += doc.Lo
			var nextUser string
			var nextUserLo int
			var usersInSameDepth = []string{metadata.Key}
			headVertex := getHeadVertex(ctx, db, metadata.Key)
			virticesCount := getVerticesCountInSameDepth(ctx, db, headVertex)
			cntr := 1
			//			log.Println("vertices count", virticesCount)
			for nextUser, nextUserLo = getNextVertex(ctx, db, headVertex, usersInSameDepth); ; {
				if nextUserLo == 0 || cntr == virticesCount {
					break
				} else if nextUserLo != 0 {
					log.Println("user:", nextUser, "personal volume =", *lo, "+", nextUserLo)
					*lo += nextUserLo
				}
				cntr++

				usersInSameDepth = append(usersInSameDepth, nextUser)
			}

			compressionTreversal(ctx, db, nextUser, lo)
			return
		}
	}
}

func main() {
	db, err := dbConn()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var lo int
	var l *int = &lo

	compressionTreversal(ctx, db, "user1", l)
	log.Println("personal volume", lo)

}
