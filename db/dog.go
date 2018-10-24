package db

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mojicaj/dog_rescue/model"
)

// Collection is the database collection set from main
var Collection *mgo.Collection

// CreateDog creates a new dog in the database with the specified data
func CreateDog(dog *model.Dog) error {
	if err := Collection.Insert(dog); err != nil {
		return fmt.Errorf("could not create a record for %s: %s", dog.Name, err)
	}
	return nil
}

// GetDog retrieves the specified dog from the database
func GetDog(name string) (dog model.Dog, err error) {
	if err := Collection.Find(bson.M{"name": name}).One(&dog); err != nil {
		return dog, fmt.Errorf("%s not found: %s", name, err)
	}
	return dog, nil
}

// GetAllDogs retrieves all dogs stored in the database
func GetAllDogs() (dogs []model.Dog, err error) {
	if err := Collection.Find(bson.M{}).All(&dogs); err != nil {
		return nil, fmt.Errorf("no dogs found: %s", err)
	}
	return dogs, nil
}

// UpdateDog updates the specified dog's information stored in the database
func UpdateDog(dog *model.Dog) error {
	if err := Collection.Update(bson.M{"name": dog.Name}, &dog); err != nil {
		return fmt.Errorf("could not update %s: %s", dog.Name, err)
	}
	return nil
}

// RemoveDog deletes the specified dog from the database
func RemoveDog(name string) error {
	if err := Collection.Remove(bson.M{"name": name}); err != nil {
		return fmt.Errorf("could not delete %s: %s", name, err)
	}
	return nil
}
