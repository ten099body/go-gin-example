package models

type Merchant struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Ctime    int    `json:"ctime"`
}

// func MerchantModel() *Merchant {
// 	return &Merchant{}
// }

// func b() {
// 	// a := Merchant{}
// 	a := MerchantModel()
// 	a.GetByAccount("aa")
// 	fmt.Println(a)
// }

func (merchant *Merchant) GetByAccount(account string) (Merchant, error) {
	var m Merchant
	err := db.Where(Merchant{Account: account}).First(&m).Error
	return m, err
}
