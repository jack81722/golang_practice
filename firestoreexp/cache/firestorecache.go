package cache

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FireStoreCache struct {
	collection string
	client     *firestore.Client
}

func NewFireStore(projID, credFilePath, collection, document string) (c *FireStoreCache, err error) {
	config := &firebase.Config{ProjectID: projID}
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, config, option.WithCredentialsFile(credFilePath))
	if err != nil {
		return
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return
	}
	c = &FireStoreCache{
		collection: collection,
		client:     client}
	return
}

func (f *FireStoreCache) Close() {
	f.client.Close()
}

func (f *FireStoreCache) Reset() (err error) {
	all, err := f.client.Collection(f.collection).
		Documents(context.Background()).GetAll()
	for _, doc := range all {
		doc.Ref.Delete(context.Background())
	}
	return
}

func (f *FireStoreCache) Set(key, value string, second int) (err error) {
	var expire time.Time
	if second > 0 {
		ttl := time.Duration(second) * time.Second
		expire = time.Now().Add(ttl)
	} else {
		expire = time.Now().AddDate(1, 0, 0)
	}
	_, err = f.client.Collection(f.collection).
		Doc(key).
		Set(context.Background(),
			map[string]interface{}{
				"value":        value,
				"expired_time": expire},
			firestore.MergeAll)
	return
}

func (f *FireStoreCache) SetMarshal(key string, value interface{}, second int) (err error) {
	data, err := json.Marshal(value)
	if err != nil {
		return
	}
	err = f.Set(key, string(data), second)
	return
}

func (f *FireStoreCache) SetInt64(key string, value int64, second int) (err error) {
	var expire time.Time
	if second > 0 {
		ttl := time.Duration(second) * time.Second
		expire = time.Now().Add(ttl)
	} else {
		expire = time.Now().AddDate(1, 0, 0)
	}
	_, err = f.client.Collection(f.collection).
		Doc(key).
		Set(context.Background(),
			map[string]interface{}{
				"value":        value,
				"expired_time": expire},
			firestore.MergeAll)
	return
}

func (f *FireStoreCache) SetIfNotExists(key, value string, second int) bool {
	if !f.Exist(key) {
		return f.Set(key, value, second) != nil
	}
	return false
}

func (f *FireStoreCache) Get(key string) (value string) {
	value, _ = f.GetOrErr(key)
	return
}

func (f *FireStoreCache) GetUnmarshal(key string, value interface{}) (err error) {
	str := f.Get(key)
	err = json.Unmarshal([]byte(str), value)
	return
}

func (f *FireStoreCache) GetOrErr(key string) (value string, err error) {
	result, err := f.client.Collection(f.collection).
		Doc(key).
		Get(context.Background())
	if err != nil {
		return
	}
	data := result.Data()
	if data["expired_time"].(time.Time).Before(time.Now()) {
		err = errors.New("expired value")
		return
	}
	value = data["value"].(string)
	return
}

func (f *FireStoreCache) GetInt64(key string) (value int64) {
	result, err := f.client.Collection(f.collection).
		Doc(key).
		Get(context.Background())
	if err != nil {
		return
	}
	data := result.Data()
	if data["expired_time"].(time.Time).Before(time.Now()) {
		return
	}
	value = data["value"].(int64)
	return
}

func (f *FireStoreCache) Remove(key string) (err error) {
	_, err = f.client.Collection(f.collection).
		Doc(key).
		Delete(context.Background())
	return
}

func (f *FireStoreCache) RemovePrefix(prefix string) (err error) {
	refs, err := f.client.Collection(f.collection).
		Documents(context.Background()).
		GetAll()
	if err != nil {
		return
	}
	for _, ref := range refs {
		if strings.HasPrefix(ref.Ref.ID, prefix) {
			ref.Ref.Delete(context.Background())
		}
	}
	return
}

func (f *FireStoreCache) Exist(key string) bool {
	ref := f.client.Collection(f.collection).
		Doc(key)
	result, err := ref.Get(context.Background())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return false
		}
	}
	data := result.Data()
	return data["expired_time"].(time.Time).After(time.Now())
}

func (f *FireStoreCache) Incr(key string) (value int64, err error) {
	expire := time.Now().AddDate(1, 0, 0)
	_, err = f.client.Collection(f.collection).
		Doc(key).
		Update(context.Background(), []firestore.Update{
			{
				Path:  "value",
				Value: firestore.Increment(1),
			},
			{
				Path:  "expired_time",
				Value: expire,
			},
		})
	return
}

func (f *FireStoreCache) OnSnapshot(doc string, callback func(map[string]interface{})) (err error) {
	snapshot := f.client.Collection(f.collection).
		Doc(doc).
		Snapshots(context.Background())
	for {
		defer func() {
			r := recover().(error)
			if r != nil {
				log.Println(r)
			}
		}()
		snap, er := snapshot.Next()
		if status.Code(err) == codes.DeadlineExceeded {
			err = nil
			return
		}
		if er != nil {
			return
		}
		if !snap.Exists() {
			return
		}
		callback(snap.Data())
	}
}
