package aggregate

import (
	"context"
	"fmt"
	"github.com/mchirico/go-firebase/pkg/gofirebase"
	"testing"
)

func TestFB_WriteMapCol2Doc2(t *testing.T) {

	credentials := "../../credentials/septapig-firebase-adminsdk.json"
	//StorageBucket := os.Getenv("FIREBASE_BUCKET")
	StorageBucket := "septapig.appspot.com"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished
	fb := &gofirebase.FB{Credentials: credentials, StorageBucket: StorageBucket}
	fb.CreateApp(ctx)

	m, err := fb.ReadCol(ctx, "Agil")
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}

	for k, _ := range m {

		fmt.Printf("k: %v\n", k)

	}

}
