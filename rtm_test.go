package rtm_package

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "Profile",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestUjiInsertData(t *testing.T) {
	Username := "John Doe"
	Email := "johm@yahoo.co.id"
	Pendidikan := "Universitas Logistik dan Bisnis Internasional"
	Tanggal_lahir := "06/04/2000"
	Bio := "Hallo"
	hasil := InsertDataProfil(MongoConn, Username, Email, Pendidikan, Tanggal_lahir, Bio)
	fmt.Println(hasil)
}

func TestUjiGetDatauser(t *testing.T) {
	id := "15648852"
	hasil := GetDataProfil(id, MongoConn, "data_user")
	fmt.Println(hasil)
}

func TestUjiDeleteData(t *testing.T) {
	Id_user := "Universitas Logistik dan Bisnis Internasional"
	DeleteDataProfil(Id_user, MongoConn, "data_user")
}
