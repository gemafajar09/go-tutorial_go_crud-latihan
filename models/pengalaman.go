package models

type Pengalaman struct {
	Id           uint   `json:"id" gorm:"primary_key"`
	TanggalAwal  string `json:"tanggal_awal"`
	TanggalAkhir string `json:"tanggal_akhir"`
	Perusahaan   string `json:"perusahaan"`
	Jobdesk      string `json:"jobdesk"`
}
