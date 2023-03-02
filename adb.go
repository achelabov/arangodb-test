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

func getVerticesCountInNextDepth(ctx context.Context, db adb.Database, rootUser string) int {
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

func getNextVertex(ctx context.Context, db adb.Database, prevUser string, usersInSameDepth []string) (user string) {
	log.Println("head", prevUser, "not in", usersInSameDepth)
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

	user = metadata.Key
	return
}

func getHeadVertex(ctx context.Context, db adb.Database, curr string) (user string) {
	querystring := "FOR v IN 2..2 INBOUND @coll GRAPH 'minions' RETURN v"
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

	user = metadata.Key
	return
}

func treversalWithoutCompression(ctx context.Context, db adb.Database, user string, lo *int) {
	querystring := "FOR v,e,p IN 1..100 OUTBOUND @coll GRAPH 'minions_test' RETURN v"
	bindVars := map[string]interface{}{
		"coll": "partners_test/" + user,
	}

	cursor, err := db.Query(ctx, querystring, bindVars)
	if err != nil {
		log.Fatalf("query failed: %v", err)
	}
	defer cursor.Close()

	for {
		var doc User
		//var metadata adb.DocumentMeta
		_, err := cursor.ReadDocument(ctx, &doc)

		if adb.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Fatalf("Doc returned: %v", err)
		}
		//log.Println("treverse ", metadata.Key)
		if doc.Lo != 0 {
			*lo += doc.Lo
		}
	}
}

func compressionTreversal(ctx context.Context, db adb.Database, user string, lo *int) {
	querystring := "FOR v,e,p IN 1..1 OUTBOUND @coll GRAPH 'minions_test' RETURN v"
	bindVars := map[string]interface{}{
		"coll": "partners_test/" + user,
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
		//log.Println("treverse ", metadata.Key)
		if doc.Lo != 0 {
			*lo += doc.Lo
			queue <- metadata.Key
			//			log.Println("user:", metadata.Key, "added", doc.Lo, "lo,", "total lo =", *lo)
		} else {
			compressionTreversal(ctx, db, metadata.Key, lo)
		}
	}
}

var queue = make(chan string, 100000)

func GetPersonalVolumes(ctx context.Context, adb adb.Database) int {
	var lo int
	var l *int = &lo
	compressionTreversal(ctx, adb, "user1", l)
	for len(queue) != 0 {
		compressionTreversal(ctx, adb, <-queue, l)
	}
	return lo
}

func GetPersonalVolumesWithoutCompression(ctx context.Context, adb adb.Database) int {
	var lo int
	var l *int = &lo
	treversalWithoutCompression(ctx, adb, "user1", l)

	return lo
}

func main() {
	db, err := dbConn()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	/* 1..2 lvl partners
	compressionTreversal(ctx, db, "user1", l)
	usersCount := len(queue)
	for usersCount > 0 {
		compressionTreversal(ctx, db, <-queue, l)
		usersCount--
	}*/

	/* 1..inf lvl partners */
	lo := GetPersonalVolumes(ctx, db)

	log.Println("personal volume", lo)

}
