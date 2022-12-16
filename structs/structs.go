package structs

type Premi struct {
	Id         int64  `json:"id"`
	Kelas      string `json:"kelas"`
	Premi      int64  `json:"premi"`
	Created_at string `json:"created"`
	Updated_at string `json:"update"`
}

type Master struct {
	NIK        string `json:"nik"`
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	TglLahir   string `json:"tgl_lahir"`
	NoHp       string `json:"no_hp"`
	Alamat     string `json:"alamat"`
	Updated_at string `json:"update"`
	Created_at string `json:"create"`
}

type DataKesehatan struct {
	NIK        string `json:"nik"`
	Kelas      string `json:"kelas"`
	Faskes     string `json:"faskes"`
	NoBPJS     string `json:"no_bpjs"`
	TotalPremi string `json:"total_premi"`
	Updated_at string `json:"update"`
	Created_at string `json:"created"`
}

type DataPembayaran struct {
	Id         int64  `json:"id"`
	NIK        string `json:"nik"`
	Periode    string `json:"periode"`
	Premi      string `json:"premi"`
	Updated_at string `json:"update"`
}