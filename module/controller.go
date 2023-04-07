package module

import (
	"context"
	"fmt"
	"os"
	// "errors"
	// "time"
    model "github.com/Fedhira/be_tagihan/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "be_tagihan",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// GET ALL
func GetAllBank(db *mongo.Database, col string) (bank []model.Bank) {
	data := db.Collection(col)
	filter := bson.M{}
	cursor, err := data.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &bank)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// func GetAllBank(db *mongo.Database, col string) (tagihan []model.Bank) {
// 	data := db.Collection(col)
// 	filter := bson.M{}
// 	cursor, err := data.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetALLData :", err)
// 	}
// 	err = cursor.All(context.TODO(), &tagihan)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }



// NASABAH
func InsertNasabah(db *mongo.Database, col string, nama_nasabah string, email string, phone_number string, alamat string) (InsertedID interface{}) {
	var nasabah model.Nasabah
	nasabah.Nama_nasabah = nama_nasabah
	nasabah.Email = email
	nasabah.Phone_number = phone_number
	nasabah.Alamat = alamat
	return InsertOneDoc(db, col, nasabah)
}

func GetNasabahFromNama(nama_nasabah string, col string, db *mongo.Database) (nasabah model.Nasabah) {
	data_nasabah := db.Collection(col)
	filter := bson.M{"nama_nasabah": nama_nasabah}
	err := data_nasabah.FindOne(context.TODO(), filter).Decode(&nasabah)
	if err != nil {
		fmt.Printf("getNasabahFromNama: %v\n", err)  
	}
	return nasabah
}


// func GetNasabahAll() (nasabah []model.Nasabah, col string, db *mongo.Database) {
// 	data_nasabah := db.Collection(col)
// 	filter := bson.D{}
// 	// var results []Nasabah
// 	cur, err := data_nasabah.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Printf("GetNasabahFromNama: %v\n", err)
// 	}
// 	err = cur.All(context.TODO(), &nasabah)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

// PENAGIH

func InsertPenagih(db *mongo.Database, col string, nama_penagih string, email string, phone_number string, total_tagihan  model.Tagihan) (InsertedID interface{}) {
	var penagih model.Penagih
	penagih.Nama_penagih = nama_penagih
	penagih.Email = email
	penagih.Phone_number = phone_number
	penagih.Total_Tagihan = total_tagihan
	return InsertOneDoc(db, col, penagih)
}

func GetPenagihFromNama(nama_penagih string, col string, db *mongo.Database) (penagih model.Penagih) {
	data_penagih := db.Collection(col)
	filter := bson.M{"nama_penagih": nama_penagih}
	err := data_penagih.FindOne(context.TODO(), filter).Decode(&penagih)
	if err != nil {
		fmt.Printf("getPenagihFromNama: %v\n", err)  
	}
	return penagih
}


// func GetPenagihAll() (penagih []model.Penagih, col string, db *mongo.Database) {
// 	data_penagih := db.Collection(col)
// 	filter := bson.D{}
// 	// var results []Penagih
// 	cur, err := data_penagih.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Printf("GetPenagihFromNama: %v\n", err)
// 	}
// 	err = cur.All(context.TODO(), &penagih)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

// TAGIHAN
func InsertTagihan(db *mongo.Database, col string, total_tagihan string, deskripsi string,  status string, tanggal_jatuhtempo string, biodata  model.Nasabah, location string, longitude float64, latitude float64  ) (InsertedID interface{}) {
	var tagihan model.Tagihan
	tagihan.Total_Tagihan = total_tagihan
	tagihan.Deskripsi = deskripsi
	tagihan.Status = status
	tagihan.Tanggal_jatuhtempo = tanggal_jatuhtempo
	tagihan.Biodata = biodata
	tagihan.Location = location
	tagihan.Longitude = longitude
	tagihan.Latitude = latitude
	return InsertOneDoc(db, col, tagihan)
}

func GetTagihanFromNama_nasabah(nama_nasabah string, col string, db *mongo.Database) (tagihan model.Tagihan) {
	data_tagihan := db.Collection(col)
	filter := bson.M{"biodata.nama_nasabah": nama_nasabah}
	err := data_tagihan.FindOne(context.TODO(), filter).Decode(&tagihan)
	if err != nil {
		fmt.Printf("getTagihanFromNama_nasabah: %v\n", err)  
	}
	return tagihan
}

// func GetAllTagihanFromNama_nasabah() (tagihan []model.Tagihan, col string, db *mongo.Database) {
// 	data_tagihan := db.Collection(col)
// 	filter := bson.D{}
// 	// var results []Tagihan
// 	cur, err := data_tagihan.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Printf("GetAllTagihanFromNama_nasabah: %v\n", err)
// 	}
// 	err = cur.All(context.TODO(), &tagihan)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

func GetAllTagihanFromNama_nasabah( nama_nasabah string, db *mongo.Database, col string) (tagihan []model.Tagihan) {
	data_tagihan := db.Collection(col)
	filter := bson.M{"biodata.nama_nasabah": nama_nasabah}
	cursor, err := data_tagihan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &tagihan)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// BANK
func InsertBank(db *mongo.Database, col string, nama_bank string, lokasi string,  daftar model.Penagih) (InsertedID interface{}) {
	var bank model.Bank
	bank.Nama_bank = nama_bank
	bank.Lokasi = lokasi
	bank.Daftar = daftar
	return InsertOneDoc(db, col, bank)
}

func GetBankFromNama_bank(nama_bank string, col string, db *mongo.Database) (bank model.Bank) {
	data_bank := db.Collection(col)
	filter := bson.M{"nama_bank": nama_bank}
	err := data_bank.FindOne(context.TODO(), filter).Decode(&bank)
	if err != nil {
		fmt.Printf("GetBankFromNama_bank: %v\n", err)  
	}
	return bank
}

// func GetBankAll() (bank []model.Bank, col string, db *mongo.Database) {
// 	data_bank := db.Collection(col)
// 	filter := bson.D{}
// 	// var results []Bank
// 	cur, err := data_bank.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Printf("GetBankFromDaftar: %v\n", err)
// 	}
// 	err = cur.All(context.TODO(), &bank)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }
