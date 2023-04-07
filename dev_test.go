package namapackage

import (
	"fmt"
	"testing"

	model "github.com/Fedhira/Tagihan/model"
	"github.com/Fedhira/Tagihan/module"
)

//NASABAH
func TestInsertNasabah(t *testing.T) {
		nama_nasabah := "Auliyah Safana"
		email:= "aol@gmail.com"
		phone_number :=  "08946737892"
		alamat := "Depok"
		
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
	nama_penagih := "Rizkyria"
	email := "12@gmail.com"
	phone_number := "0813456789"
	total_tagihan := model.Tagihan{
		Total_Tagihan : "1000000",
		Deskripsi : "Kartu Kredit",
		Status : "Belum Lunas",
		Tanggal_jatuhtempo : "23 Maret 2021",
		Biodata : model.Nasabah{
			Nama_nasabah : "Auliyah Safana",
			Email : "aol@gmail.com",
			Phone_number : "08946737892",
			Alamat : "Depok",
		},
		Location : "Depok",
		Longitude : 89.567899,
		Latitude : 124.781781,
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
	total_tagihan := "1000000"
	deskripsi := "Kartu Kredit"
	status := "Belum Lunas"
	tanggal_jatuhtempo := "23 Maret 2021"
	biodata := model.Nasabah{
		Nama_nasabah : "Auliyah Safana",
		Email : "aol@gmail.com",
		Phone_number : "08946737892",
		Alamat : "Depok",
		}
	location := "Depok"
	longitude := 89.567899
	latitude := 124.781781

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
	total_tagihan:= model.Tagihan{
			Total_Tagihan : "1000000",
			Deskripsi : "Kartu Kredit",
			Status : "Belum Lunas",
			Tanggal_jatuhtempo : "23 Maret 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Auliyah Safana",
				Email : "aol@gmail.com",
				Phone_number : "08946737892",
				Alamat : "Depok",
		},
			Location : "Depok",
			Longitude : 89.567899,
			Latitude : 124.781781,
	}

	daftar := model.Penagih{
		Nama_penagih : "Rizkyria",
		Email : "12@gmail.com",
		Phone_number : "0813456789",
		Total_Tagihan : model.Tagihan{
			Total_Tagihan : "1000000",
			Deskripsi : "Kartu Kredit",
			Status : "Belum Lunas",
			Tanggal_jatuhtempo : "23 Maret 2021",
			Biodata : model.Nasabah{
				Nama_nasabah : "Auliyah Safana",
				Email : "aol@gmail.com",
				Phone_number : "08946737892",
				Alamat : "Depok",
		},
		Location : "Depok",
		Longitude : 89.567899,
		Latitude : 124.781781,
		},	
	}

	biodata := model.Nasabah{
			Nama_nasabah : "Auliyah Safana",
			Email : "aol@gmail.com",
			Phone_number : "08946737892",
			Alamat : "Depok",
	}
	
	hasil:= module.InsertBank(module.MongoConn, "bank", nama_bank, lokasi, total_tagihan, daftar, biodata)
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
