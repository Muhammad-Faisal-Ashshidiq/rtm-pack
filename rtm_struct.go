package rtm_package

import "go.mongodb.org/mongo-driver/bson/primitive"

type Profil struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Id_user      string             `bson:"id_user,omitempty" json:"id_user,omitempty"`
	Username     string             `bson:"username,omitempty" json:"username,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Pendidikan   string             `bson:"pendidikan,omitempty" json:"pendidikan,omitempty"`
	Tanggal_lahir string             `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty"`
	Bio          string             `bson:"bio,omitempty" json:"bio,omitempty"`
}
