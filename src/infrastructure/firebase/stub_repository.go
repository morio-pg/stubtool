package firebase

import (
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/morio-pg/stubtool/src/domain/model"
	"github.com/morio-pg/stubtool/src/domain/repository"
	"google.golang.org/api/iterator"
)

type stubRepository struct{}

// NewStubRepository return repository.StubRepository
func NewStubRepository() repository.StubRepository {
	return &stubRepository{}
}

func (r *stubRepository) Get(c *gin.Context, stubID string) (stub *model.Stub, err error) {
	client, err := r.firestore(c)
	if err != nil {
		return nil, fmt.Errorf(model.ErrInitializingApp, err)
	}
	defer client.Close()

	dsnap, err := client.Collection("stubs").Doc(stubID).Get(c)
	if err != nil {
		return nil, fmt.Errorf(model.ErrGetData, err)
	}
	dsnap.DataTo(&stub)

	return stub, nil
}

func (r *stubRepository) GetAll(c *gin.Context, uid string) (stubs []model.Stub, err error) {
	client, err := r.firestore(c)
	if err != nil {
		return nil, fmt.Errorf(model.ErrInitializingApp, err)
	}
	defer client.Close()

	iter := client.Collection("stubs").
		Where("uid", "==", uid).
		OrderBy("created_at", firestore.Desc).
		Documents(c)

	stubs = []model.Stub{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf(model.ErrGetData, err)
		}

		var stub model.Stub
		if err := doc.DataTo(&stub); err != nil {
			return nil, fmt.Errorf(model.ErrConvertData, err)
		}
		stubs = append(stubs, stub)
	}

	return stubs, nil
}

func (r *stubRepository) Save(c *gin.Context, stub *model.Stub) (err error) {
	client, err := r.firestore(c)
	if err != nil {
		return fmt.Errorf(model.ErrInitializingApp, err)
	}
	defer client.Close()

	_, err = client.Collection("stubs").Doc(stub.StubID).Set(c, stub)
	if err != nil {
		return fmt.Errorf(model.ErrSaveData, err)
	}

	return nil
}

func (r *stubRepository) Delete(c *gin.Context, stubID string) (err error) {
	client, err := r.firestore(c)
	if err != nil {
		return fmt.Errorf(model.ErrInitializingApp, err)
	}
	defer client.Close()

	_, err = client.Collection("stubs").Doc(stubID).Delete(c)
	if err != nil {
		return fmt.Errorf(model.ErrDeleteData, err)
	}

	return nil
}

func (r *stubRepository) DeleteAll(c *gin.Context, uid string) (err error) {
	client, err := r.firestore(c)
	if err != nil {
		return fmt.Errorf(model.ErrInitializingApp, err)
	}
	defer client.Close()

	iter := client.Collection("stubs").
		Where("uid", "==", uid).
		Documents(c)

	batch := client.Batch()
	i := 0
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf(model.ErrGetData, err)
		}

		batch.Delete(doc.Ref)
		i++
	}

	if i > 0 {
		if _, err := batch.Commit(c); err != nil {
			return fmt.Errorf(model.ErrDeletingUser, err)
		}
	}

	return nil
}

func (r *stubRepository) firestore(c *gin.Context) (*firestore.Client, error) {
	app, err := firebase.NewApp(c, nil, opt)
	if err != nil {
		return nil, err
	}

	return app.Firestore(c)
}
