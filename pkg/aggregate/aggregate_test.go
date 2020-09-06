package aggregate

import (
	"context"
	"github.com/mchirico/go-firebase/pkg/gofirebase"
	"testing"
	"fmt"
)

func TestFB_WriteMapCol2Doc2(t *testing.T) {

	credentials := "../../credentials/septapig-firebase-adminsdk.json"
	//StorageBucket := os.Getenv("FIREBASE_BUCKET")
	StorageBucket := "septapig.appspot.com"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished
	fb := &gofirebase.FB{Credentials: credentials, StorageBucket: StorageBucket}
	fb.CreateApp(ctx)

	m, _ := fb.ReadCol(ctx, "Agil")
	for i,v := range m {
		fmt.Printf("%v,%v\n",i,v)
	}

}
