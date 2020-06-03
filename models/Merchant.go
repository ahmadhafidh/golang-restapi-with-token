package models

type Merchant struct {
	Id                 string `form:"id" json:"id"`
	MerchantName       string `form:"merchant_name" json:"merchant_name"`
	MerchantId         string `form:"merchant_id" json:"merchant_id"`
	MerchantSecretKey  string `form:"merchant_secret_key" json:"merchant_secret_key"`
	EncryptionKey      string `form:"encryption_key" json:"encryption_key"`
	AccountNumber      string `form:"account_number" json:"account_number"`
	PicEmail           string `form:"pic_email" json:"pic_email"`
	PicAddress         string `form:"pic_address" json:"pic_address"`
	PicPhoneNumber     string `form:"pic_phone_number" json:"pic_phone_number"`
	Status             string `form:"status" json:"status"`
	CreatedDate        string `form:"created_date" json:"created_date"`
	CreatedBy          string `form:"created_by" json:"created_by"`
	UpdatedDate        string `form:"updated_date" json:"updated_date"`
	UpdatedBy          string `form:"updated_by" json:"updated_by"`
	FeeTransaction     string `form:"fee_transaction" json:"fee_transaction"`
	FeeSpe             string `form:"fee_spe" json:"fee_spe"`
	FeeBni             string `form:"fee_bni" json:"fee_bni"`
	FeeTotal           string `form:"fee_total" json:"fee_total"`
	IsFeeInclude       string `form:"is_fee_include" json:"is_fee_include"`
	AccessToken        string `form:"access_token" json:"access_token"`
	AccessTokenExpired string `form:"access_token_expired" json:"access_token_expired"`
	IpWhitelist        string `form:"ip_whitelist" json:"ip_whitelist"`
	CallbackUrlDev     string `form:"callback_url_dev" json:"callback_url_dev"`
	CallbackUrlProd    string `form:"callback_url_prod" json:"callback_url_prod"`
}

func (Merchant) TableName() string {
	return "merchant"
}
