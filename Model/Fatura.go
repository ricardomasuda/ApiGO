package Model

type Fatura struct {
	IdFatura  int    `db:"id_fatura" json:"id_fatura"`
	Valor     int    `db:"valor" json:"valor"`
	Pagou     int    `db:"pagou" json:"pagou"`
	Idusuario     int    `db:"idusuario" json:"idusuario"`
	Nome_empresa string `db:"nome_empresa" json:"nome_empresa"`
	Data_vencimento    string `db:"data_vencimento" json:"data_vencimento"`
}
