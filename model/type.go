package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Nasabah struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_nasabah string             `bson:"nama_nasabah,omitempty" json:"nama_nasabah,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
}

type Penagih struct {
	ID           primitive.ObjectID        `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_penagih        string             `bson:"nama_penagih,omitempty" json:"nama_penagih,omitempty"`
	Email		 		string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone_number 		string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Total_Tagihan       Tagihan            `bson:"total_tagihan,omitempty" json:"total_tagihan,omitempty"`
}

type  Tagihan struct {
	ID           		primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	Total_Tagihan   	string      		`bson:"total_tagihan,omitempty" json:"total_tagihan,omitempty"`
	Deskripsi 			string   			`bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Status				string				`bson:"status,omitempty" json:"status,omitempty"`            
	Tanggal_jatuhtempo  string              `bson:"tanggal_jatuhtempo,omitempty" json:"tanggal_jatuhtempo,omitempty"`
	Biodata      		Nasabah             `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Location            string              `bson:"location,omitempty" json:"location,omitempty"`
	Longitude           float64             `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude            float64             `bson:"latitude,omitempty" json:"latitude,omitempty"`
}

type Lokasi struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama                string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Batas               Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori            string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}

type  Bank struct {
	ID           		primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_bank           string             `bson:"nama_bank,omitempty" json:"nama_bank,omitempty"`
	Lokasi              string             `bson:"lokasi,omitempty" json:"lokasi,omitempty"`
	Daftar              Penagih            `bson:"daftar,omitempty" json:"daftar,omitempty"`
}