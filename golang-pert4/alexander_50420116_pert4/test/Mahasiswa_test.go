package test

import (
	"alexander_50420116_pert4/model" // Ganti dengan nama folder kalian masing￾masing
	"testing"
)

func TestMahasiswa(t *testing.T) {
	var dataInsertMhs = []model.Mahasiswa{
		model.Mahasiswa{
			NPM:   "12345678",
			Nama:  "Budi Doremi",
			Kelas: "3IA20",
		},
		model.Mahasiswa{
			NPM:   "50420116",
			Nama:  "Alexander",
			Kelas: "3I16",
		},
		model.Mahasiswa{
			NPM:   "44444444",
			Nama:  "DoBud",
			Kelas: "4IA21",
		},
	}
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Testing insert mahasiswa", func(t *testing.T) {
		for _, dataInsert := range dataInsertMhs {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("Testing update mahasiswa", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"nama": "Abdi Teh reza"}
		data := dataInsertMhs[0]
		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing Get mahasiswa", func(t *testing.T) {
		_, err := model.GetMahasiswa(db, "44444444")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing Get mahasiswa", func(t *testing.T) {
		_, err := model.GetAllMahasiswa(db)
		if err != nil {
			t.Fatal(err)
		}
	})
	// t.Run("Testing delete mahasiswa", func(t *testing.T) {
	// data := dataInsertMhs[0]
	// if err := data.Delete(db); err != nil {
	// t.Fatal(err)
	// }
	// })
	defer db.Close()
}
