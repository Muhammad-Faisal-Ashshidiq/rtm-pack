package rtm_package

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Profile struct {
	Id_user      string
	Username     string
	Email        string
	Pendidikan   string
	Tanggal_lahir string
	Bio          string
}

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataProfil(db *mongo.Database, username, email, pendidikan, tanggal_lahir, bio string) (insertedID interface{}) {
	dtProf := Profile{
		Username:     username,
		Email:        email,
		Pendidikan:   pendidikan,
		Tanggal_lahir: tanggal_lahir,
		Bio:          bio,
	}
	return InsertOneDoc(db, "data_Profile", dtProf)
}

func GetDataProfil(Bio string, db *mongo.Database, col string) (data Profile) {
	user := db.Collection(col)
	filter := bson.M{"bio": Bio}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getbybio: %v\n", err)
	}
	return data
}

func GetDataUsername(username string, db *mongo.Database, col string) (data Profile) {
	user := db.Collection(col)
	filter := bson.M{"username": username}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getbyusername: %v\n", err)
	}
	return data
}

func DeleteDataProfil(Id_user string, db *mongo.Database, col string) {
	user := db.Collection(col)
	filter := bson.M{"Id_user": Id_user}
	_, err := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataProf : %v\n", err)
	}
	fmt.Println("Succes Delete data Profile")
}

func DeleteDataUsername(username string, db *mongo.Database, col string) {
	user := db.Collection(col)
	filter := bson.M{"username": username}
	_, err := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataProfile : %v\n", err)
	}
	fmt.Println("Succes Delete data")
}

//fungsi untuk meng-generate tanggal_lahir menjadi usia
func KalkulasiUsia(birthDate string) (int, error) {
	layout := "21/07/2023" // Format tanggal lahir (hari/bulan/tahun)
	birth, err := time.Parse(layout, birthDate)
	if err != nil {
		return 0, err
	}

	currentDate := time.Now()
	age := currentDate.Year() - birth.Year()

	// Pengecekan apabila tanggal tahirnya supaya bisa menyesuaikan tahun user
	if currentDate.YearDay() < birth.YearDay() {
		age--
	}

	return age, nil
}

func GetAgeFromProfile(profile Profile) (int, error) {
	return KalkulasiUsia(profile.Tanggal_lahir)
}
