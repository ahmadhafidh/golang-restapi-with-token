package services

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"spc-api-v2/configs"
	"spc-api-v2/models"
	"spc-api-v2/utils"
	"strings"
	"time"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	var rc string
	var ResClient models.ResClient
	var merchantObj models.Merchant
	var response models.ResponseToken

	username, password, ok := r.BasicAuth()
	if !ok {
		rc = configs.Incorrectcredential
		ResClient = utils.ResponseBuilder(rc)
		utils.OutputJSON(w, ResClient)
		return
	}

	merchantDB := GetMerchant(username, password)

	if merchantDB.Id == "" {
		rc = configs.Authenticationfailed
		ResClient = utils.ResponseBuilder(rc)
		utils.OutputJSON(w, ResClient)
		return
	}

	Checkip, ip := utils.GetIP(r, merchantDB.IpWhitelist)
	fmt.Println("IP Address : " + ip)

	if !Checkip {
		rc = configs.Invalidipaddress
		ResClient = utils.ResponseBuilder(rc)
		utils.OutputJSON(w, ResClient)
		return
	} else {
		dt := time.Now()

		layout := "2006-01-02 15:04:05"
		parsedate, err := time.Parse(layout, merchantDB.AccessTokenExpired)

		if err != nil {
			fmt.Println(err)
		}

		//check token null or already expired
		if len(merchantDB.AccessToken) == 0 || dt.After(parsedate) {
			merchantObj.Id = merchantDB.Id
			merchantObj.AccessTokenExpired = dt.Add(time.Second * time.Duration(3500)).Format("2006-01-02 15:04:05")
			hash := sha256.Sum256([]byte(merchantObj.MerchantName + utils.Uniqid("spc") + merchantObj.MerchantSecretKey))
			merchantObj.AccessToken = fmt.Sprintf("%x", hash)
			validate := UpdateMerchantToken(merchantObj, merchantObj.Id)
			if !validate {
				log.Println("Failed Update Merchant")
				return
			} else {
				log.Println("Success Update Merchant")
			}
		} else {
			layout := "2006-01-02 15:04:05"
			parsedate, err := time.Parse(layout, merchantDB.AccessTokenExpired)
			if err != nil {
				fmt.Println(err)
			}
			merchantObj.AccessTokenExpired = parsedate.Format("2006-01-02 15:04:05")
			merchantObj.AccessToken = merchantDB.AccessToken
			merchantObj.Id = merchantDB.Id
		}

		response.ResponseCode = configs.Success
		response.ResponseMessage = "Success"
		response.Data.AccessToken = merchantObj.AccessToken
		response.Data.ExpiredDatetime = merchantObj.AccessTokenExpired
		response.Data.TokenType = "Bearer"

		utils.OutputJSON(w, response)
	}
}

func GetTokenFromBNI(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	URL := configs.BNIUrl
	v := url.Values{}
	v.Set("grant_type", configs.BNIGrantType)
	// v.Set("username", configs.BNIUser)
	// v.Set("password", configs.BNIPass)

	//pass the values to the request's body
	req, err := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", configs.BNIContenTypeForm)
	req.SetBasicAuth(configs.BNIClientId, configs.BNIClientSecret)

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var iot models.TokenBNI
	err = json.Unmarshal(bodyText, &iot)
	if err != nil {
		panic(err)
	}

	// Get Access Token
	fmt.Println("Access Token:", string(iot.AccessToken))

	// Marshal back to json (as original)
	out, _ := json.Marshal(&iot)
	fmt.Println(string(out))

	utils.OutputJSON(w, iot)

}
