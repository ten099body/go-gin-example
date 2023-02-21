package models

type Merchant struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Account  string `json:"username"`
	Password string `json:"password"`
	Ctime    int    `json:"ctime"`
}

func (merchant *Merchant) GetByAccount(account string) (Merchant, error) {
	var m Merchant
	err := db.Where(Merchant{Account: account}).First(&m).Error
	return m, err
}
