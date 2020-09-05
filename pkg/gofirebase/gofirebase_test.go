package gofirebase

import (
	"context"
	"fmt"
	util "github.com/mchirico/go-firebase/pkg/utils"

	"testing"
)

func TestReadWrite_Firebase(t *testing.T) {
	credentials := "../../credentials/septapig-firebase-adminsdk.json"

	//StorageBucket := os.Getenv("FIREBASE_BUCKET")
	StorageBucket := "septapig.appspot.com"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	number := 9
	doc := make(map[string]interface{})
	doc["application"] = "FirebaseGo"
	doc["function"] = "TestAuthenticate"
	doc["test"] = "This is example text..."
	doc["random"] = number

	fb := &FB{Credentials: credentials, StorageBucket: StorageBucket}
	fb.CreateApp(ctx)
	fb.WriteMap(ctx, doc, "testGoFirebase", "go-gofirebase-v4")
	fb.WriteMapCol2Doc2(ctx, doc, "testGoFirebase", "go-gofirebase-v4","updates","doc")

	resultFind, err := fb.Find(ctx, "testGoFirebase", "function", "==", "TestAuthenticate")

	if resultFind["test"] != "This is example text..." {
		t.Fatalf("Find not working")
	}

	dsnap, _ := fb.ReadMap(ctx, "testGoFirebase", "go-gofirebase-v4")
	result := dsnap.Data()

	fmt.Printf("Document data: %v %v\n", result["random"].(int64), number)
	if result["random"].(int64) != 9 {
		t.Fatalf("Didn't return correct value\n")
	}

	util.CreateDir(".slop")
	data := []byte("ABC€")

	util.Write(".slop/junk.txt", data, 0600)
	fb.Bucket.Upload(ctx, ".slop/junk.txt")
	util.RmDir(".slop")
	err = fb.Bucket.DeleteFile(ctx, ".slop/junk.txt")

	if err != nil {
		t.Logf("Problem with buckets")
	}

}
