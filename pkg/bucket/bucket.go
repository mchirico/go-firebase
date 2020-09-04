package bucket

import (
	"cloud.google.com/go/storage"
	"context"

	"io"
	"io/ioutil"
	"os"
	"sync"
)

type Bucket struct {
	sync.Mutex
	client        *storage.Client
	StorageBucket string
}

func FBInitBucket(client *storage.Client, bucket string) *Bucket {
	return &Bucket{client: client, StorageBucket: bucket}

}

func (b *Bucket) Upload(ctx context.Context, file string) error {
	b.Lock()
	defer b.Unlock()

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	wc := b.client.Bucket(b.StorageBucket).Object(file).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

func (b *Bucket) Download(ctx context.Context, file string) ([]byte, error) {
	b.Lock()
	defer b.Unlock()
	rc, err := b.client.Bucket(b.StorageBucket).Object(file).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (b *Bucket) DeleteFile(ctx context.Context, file string) error {
	b.Lock()
	defer b.Unlock()
	o := b.client.Bucket(b.StorageBucket).Object(file)
	if err := o.Delete(ctx); err != nil {
		return err
	}
	return nil
}
