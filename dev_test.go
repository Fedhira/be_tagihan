package namapackage

import (
	"fmt"
	"testing"

	model "github.com/Fedhira/be_tagihan/model"
	"github.com/Fedhira/be_tagihan/module"
)

//NASABAH
func TestInsertNasabah(t *testing.T) {
	nama_nasabah := "Farel"
	email:= "farel@gmail.com"
	phone_number := "08360086721"
	alamat :="Bandung"
		
	hasil:= module.InsertNasabah(module.MongoConn, "nasabah", nama_nasabah, email, phone_number, alamat)
	fmt.Println(hasil)
}

func TestGetNasabahFromNama(t *testing.T) {
	nama_nasabah := "Auliyah Safana"
	biodata:= module.GetNasabahFromNama(nama_nasabah, "nasabah", module.MongoConn)
	fmt.Println(biodata)
}

func TestGetAllTagihanFromNama_nasabah(t *testing.T) {
	nama_nasabah := "Auliyah Safana"
	data1 := module.GetAllTagihanFromNama_nasabah(nama_nasabah, module.MongoConn, "tagihan")
	fmt.Println(data1)
}

//PENAGIH
func TestInsertPenagih(t *testing.T) {
	nama_penagih := "Dimas"
	email := "dimas@gmail.com"
	phone_number := "0879013425"
	total_tagihan := model.Tagihan{
		Total_Tagihan : "10000000",
		Deskripsi : "Kartu Kredit",
		Status : "Belum Lunas",
		Tanggal_jatuhtempo : "16 Desember 2021",
		Biodata : model.Nasabah{
			Nama_nasabah :"Farel",
			Email : "farel@gmail.com",
			Phone_number : "08360086721",
			Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 90.675675,
		Latitude : 200.126126,
	}
	
	hasil:= module.InsertPenagih(module.MongoConn, "penagih", nama_penagih, email, phone_number, total_tagihan)
	fmt.Println(hasil)
}

func TestGetPenagihFromNama(t *testing.T) {
	nama_penagih := "Rizkyria"
	biodata:= module.GetPenagihFromNama(nama_penagih,  "penagih", module.MongoConn)
	fmt.Println(biodata)
}

// func TestGetPenagihAll(t *testing.T) {
// 	biodata := module.GetPenagihAll()
// 	fmt.Println(biodata)
// }

//TAGIHAN
func TestInsertTagihan(t *testing.T) {
	total_tagihan := "10000000"
	deskripsi := "Kartu Kredit"
	status := "Belum Lunas"
	tanggal_jatuhtempo := "16 Desember 2021"
	biodata := model.Nasabah{
		Nama_nasabah :"Farel",
		Email : "farel@gmail.com",
		Phone_number : "08360086721",
		Alamat : "Bandung",
		}
	location := "Bandung"
	longitude := 90.675675
	latitude := 200.126126
	hasil:= module.InsertTagihan(module.MongoConn, "tagihan", total_tagihan, deskripsi, status, tanggal_jatuhtempo, biodata, location, longitude, latitude )
	fmt.Println(hasil)
}

func TestGetTagihanFromNama_nasabah(t *testing.T) {
	nama_nasabah := "Auliyah Safana"
	biodata:= module.GetTagihanFromNama_nasabah( nama_nasabah, "tagihan", module.MongoConn)
	fmt.Println(biodata)
}

// func TestGetTagihanAll(t *testing.T) {
// 	biodata := module.GetTagihanAll()
// 	fmt.Println(biodata)
// }

func TestInsertBank(t *testing.T) {
	nama_bank := "bank abc"
	lokasi := "Bandung"
	daftar := model.Penagih{
		Nama_penagih :  "Dimas",
		Email : "dimas@gmail.com",
		Phone_number : "0879013425",
		Total_Tagihan : model.Tagihan{
			Total_Tagihan : "10000000",
		Deskripsi : "Kartu Kredit",
		Status : "Belum Lunas",
		Tanggal_jatuhtempo : "16 Desember 2021",
		Biodata : model.Nasabah{
			Nama_nasabah :"Farel",
			Email : "farel@gmail.com",
			Phone_number : "08360086721",
			Alamat : "Bandung",
		},
		Location : "Bandung",
		Longitude : 90.675675,
		Latitude : 200.126126,
		},
	}
	
	hasil:= module.InsertBank(module.MongoConn, "bank", nama_bank, lokasi, daftar)
	fmt.Println(hasil)
}

func TestGetBankFromNama_bank(t *testing.T) {
	nama_bank := "bank abc"
	biodata:= module.GetBankFromNama_bank(nama_bank, "bank", module.MongoConn)
	fmt.Println(biodata)
}

// func TestGetAllTagihanF(t *testing.T) {
// 	biodata := module.GetAllTagihan(module.MongoConn, "bank")
// 	fmt.Println(biodata)
// }
