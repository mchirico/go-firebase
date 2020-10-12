package gofirebase

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/iterator"

	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/mchirico/go-firebase/pkg/bucket"

	"google.golang.org/api/option"
	"log"
	"sync"
)

// Firebase struct
type FB struct {
	sync.Mutex
	Credentials   string
	App           *firebase.App
	StorageBucket string

	Bucket *bucket.Bucket
	// Private
	bucketHandle *storage.BucketHandle
	err          error
}

func (fb *FB) WriteMap(ctx context.Context, doc map[string]interface{}, collection string, Doc string) {
	fb.Lock()
	defer fb.Unlock()
	client, err := fb.App.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.Collection(collection).Doc(Doc).Set(ctx, doc)

	if err != nil {
		log.Fatalf("Failed adding record: %v", err)
	}
	defer client.Close()

}

func (fb *FB) WriteMapCol2Doc2(ctx context.Context, doc map[string]interface{},
	collection string, Doc string, collection2, Doc2 string) {
	fb.Lock()
	defer fb.Unlock()
	client, err := fb.App.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.Collection(collection).Doc(Doc).Collection(collection2).Doc(Doc2).Set(ctx, doc)

	if err != nil {
		log.Fatalf("Failed adding record: %v", err)
	}
	defer client.Close()

}

func (fb *FB) ReadMap(ctx context.Context, path string, Doc string) (*firestore.DocumentSnapshot,
	error) {
	fb.Lock()
	defer fb.Unlock()
	client, err := fb.App.Firestore(ctx)
	defer client.Close()

	dsnap, err := client.Collection(path).Doc(Doc).Get(ctx)
	if err != nil {
		return dsnap, err
	}
	return dsnap, err
}

func (fb *FB) ReadCol(ctx context.Context, path string) ([]map[string]interface{}, error) {
	fb.Lock()
	defer fb.Unlock()
	client, err := fb.App.Firestore(ctx)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer client.Close()

	iter := client.Collection(path).Documents(ctx)
	resultFind := []map[string]interface{}{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return resultFind, err
		}
		m := map[string]interface{}{}
		for k, v := range doc.Data() {
			fmt.Printf("K: %v\n", k)
			m[k] = v
		}
		resultFind = append(resultFind, m)

	}
	return resultFind, nil
}

func (fb *FB) ReadMapCol2Doc2(ctx context.Context, path, Doc, col2, Doc2 string) (*firestore.DocumentSnapshot,
	error) {
	fb.Lock()
	defer fb.Unlock()
	client, err := fb.App.Firestore(ctx)
	defer client.Close()

	dsnap, err := client.Collection(path).Doc(Doc).Collection(col2).Doc(Doc2).Get(ctx)
	if err != nil {
		return dsnap, err
	}
	return dsnap, err
}

func (fb *FB) Find(ctx context.Context, collection, path, op string, value interface{}) (map[string]interface{}, error) {
	fb.Lock()
	defer fb.Unlock()
	client, err := fb.App.Firestore(ctx)
	// You need to close
	defer client.Close()
	if err != nil {
		return map[string]interface{}{}, err
	}

	// query := client.Collection(collection).Where("state", "==", "CA")
	iter := client.Collection(collection).Where(path, op, value).Documents(ctx)
	resultFind := map[string]interface{}{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return map[string]interface{}{}, err
		}

		for k, v := range doc.Data() {
			resultFind[k] = v
		}

	}

	return resultFind, nil

}

func (fb *FB) FindCol2Doc2(ctx context.Context, collection, Doc, col2, path, op string, value interface{}) (map[string]interface{}, error) {
	fb.Lock()
	defer fb.Unlock()
	client, err := fb.App.Firestore(ctx)
	// You need to close
	defer client.Close()
	if err != nil {
		return map[string]interface{}{}, err
	}

	// query := client.Collection(collection).Where("state", "==", "CA")
	iter := client.Collection(collection).Doc(Doc).Collection(col2).Where(path, op, value).Documents(ctx)
	resultFind := map[string]interface{}{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return map[string]interface{}{}, err
		}
		fmt.Println(doc.Data())
		for k, v := range doc.Data() {
			resultFind[k] = v
		}

	}

	return resultFind, nil

}

func (fb *FB) CreateApp(ctx context.Context) (*firebase.App, error) {
	fb.Lock()
	defer fb.Unlock()
	opt := option.WithCredentialsFile(fb.Credentials)
	storageClient, err := storage.NewClient(ctx, opt)

	fb.Bucket = bucket.FBInitBucket(storageClient, fb.StorageBucket)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	fb.App = app
	return app, nil
}
