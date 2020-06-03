package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"spc-api-v2/configs"
	"spc-api-v2/models"
	"strings"
	"time"
)

//func to show print error on terminal
func Logger(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

//func to show all json
func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		Logger(w.Write([]byte(err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	Logger(w.Write(res))
}

//compare ip
func InArray(val string, array []string) (exists bool) {
	exists = false
	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})
	for _, dbprefix := range array {
		fmt.Println(dbprefix)
		exists = strings.HasPrefix(val, dbprefix)
		if !exists {
			continue
		} else {
			exists = true
			return
		}
	}
	return
}

//function get ip client
func GetIP(r *http.Request, IPWhitelist string) (bool, string) {

	var ipdetect string
	ipdetect = r.Header.Get("X-FORWARDED-FOR")

	if ipdetect == "" {
		ipstrings := strings.Contains(r.RemoteAddr, "[::1]")
		if ipstrings {
			ipdetect = "127.0.0.1"
		} else {
			ipremote := strings.Split(r.RemoteAddr, ":")
			ipdetect = ipremote[0]
		}
	}
	IPWhitelist = strings.Replace(IPWhitelist, " ", "", -1)

	arrIP := strings.Split(IPWhitelist, ",")
	exist := InArray(ipdetect, arrIP)

	return exist, ipdetect
}

//function all response to client
func ResponseBuilder(rc string) models.ResClient {
	var ResClient models.ResClient
	ResClient.Status = rc
	ResClient.Status, ResClient.Message = GetMessage(rc)
	return ResClient
}

//get message
func GetMessage(rc string) (string, string) {
	var message string
	var status string
	status, message = GetMessageFromDB(rc)
	return status, message
}

//get message from Db
func GetMessageFromDB(rc string) (string, string) {
	var message string
	var RC models.RefCode
	db, _ := configs.Connect()
	// select * from table where "000"
	db.Table(RC.TableName()).Where("code = ?", rc).Scan(&RC)
	message = RC.Message
	return rc, message
}

//get uniq id for generate token
func Uniqid(prefix string) string {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000
	return fmt.Sprintf("%s%08x%05x", prefix, sec, usec)
}
