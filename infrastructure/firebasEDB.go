package infrastructure

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
)

func NewFirebaseDB() {
	ctx := context.Background()
	//configure firebase db
	conf := firebase.Config{
		DatabaseURL: "https://letschat-c3e14-default-rtdb.asia-southeast1.firebasedatabase.app/",
	}
	serviceAccountKeyFilePath, err := filepath.Abs("./FirebaseSevicekey.json")

}
