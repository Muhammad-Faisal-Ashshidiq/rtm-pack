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

func InsertDataProfil(db *mongo.Database, username, email, pendidikan string, tanggal_lahir string, bio string) (InsertedID interface{}) {
	var dtProf Profil
	dtProf.Username = username
	dtProf.Email = email
	dtProf.Pendidikan = pendidikan
	dtProf.Tanggal_lahir = bio
	dtProf.Bio = bio
	return InsertOneDoc(db, "data_user", dtProf)
}

func GetDataProfil(pendidikan string, db *mongo.Database, col string) (data Profile) {
	user := db.Collection(col)
	filter := bson.M{"pendidikan": pendidikan}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getbypendidikan: %v\n", err)
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
func KalkulasiUsia(Tanggal_lahir string) (int, error) {
	layout := "02/01/1999" // Format tanggal lahir (hari/bulan/tahun)
	tglLahir, err := time.Parse(layout, Tanggal_lahir)
	if err != nil {
		return 0, err
	}

	currentDate := time.Now()
	usia := currentDate.Year() - tglLahir.Year()

	// Pengecekan apabila tanggal tahirnya supaya bisa menyesuaikan tahun user
	if currentDate.YearDay() < tglLahir.YearDay() {
		usia--
	}

	return usia, nil
}

func GetAgeFromProfile(profile Profile) (int, error) {
	return KalkulasiUsia(profile.Tanggal_lahir)
}
