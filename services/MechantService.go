package services

import (
	"spc-api-v2/configs"
	"spc-api-v2/models"

	_ "github.com/go-sql-driver/mysql"
)

var validate bool

func GetMerchant(merchantId string, merchantKey string) models.Merchant {

	var merchant models.Merchant

	db, _ := configs.Connect()
	defer db.Close()

	db.Table(merchant.TableName()).Where("merchant_id=? and merchant_secret_key=?", merchantId, merchantKey).Scan(&merchant)
	return merchant
}

func UpdateMerchantToken(req models.Merchant, merchantid string) bool {
	var Merchant models.Merchant
	db, _ := configs.Connect()
	db.Table(Merchant.TableName()).Where("id = ?", merchantid).Update(&req)
	if db.Error != nil {
		validate = false
	} else {
		validate = true
	}
	defer db.Close()
	return validate
}
