package gocap

import (
	"bytes"
	"encoding/json"
	"log"
	"io/ioutil"
	"net/http"
)




func (x *structs) GetTaskID(APIKey string) interface{} {
	payload := map[string]interface{}{
        "clientKey": APIKey,
        "task": map[string]interface{}{
			"type": "HCaptchaTaskProxyless",
			"userAgent": x.UserAgent,
			"websiteKey": x.WebKey,
			"websiteURL": "https://" + x.WebURL,
		},
	}
	xp,_ := json.Marshal(payload)
	cap := capresp{}
	req,_ := http.NewRequest("POST", "https://api.capmonster.cloud/createTask", bytes.NewBuffer(xp))
	resp, err := Client.Do(req)
	call_err(err, false)
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	call_err(err, false)


	if resp.StatusCode == 200 {
		cap.TaskID = data["taskId"]
	}
	return cap.TaskID
}


func (x *structs) SetPayload(Key string, Url string, Agent string) error {
	x.WebKey = Key
	x.WebURL = Url
	x.UserAgent = Agent

	return nil
}


func (x *capresp) Solve(apikey string) interface{} {
	xpayload := map[string]interface{}{
		"clientKey": apikey,
		"taskId": x.TaskID,
	}
	cap := capresp{}
	xpy,_ := json.Marshal(xpayload)
	for true {
		req, err := http.NewRequest("POST", "https://api.capmonster.cloud/getTaskResult", bytes.NewBuffer(xpy))
		call_err(err, true)

		resp, err := Client.Do(req)		
		call_err(err, true)
		
		defer resp.Body.Close()
		responseBody := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&responseBody)
		status := responseBody["status"]

		if status == "ready" {	
			cap.Captcha = responseBody["solution"].(map[string]interface{})["gRecaptchaResponse"].(string)
			break
		} else if status == "processing" {
			continue
		} else {
			log.Print("[ERR] | ", responseBody)
		}
	
	}
	return cap.Captcha
}
	


