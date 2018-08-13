package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func main() {
    fmt.Println("Starting the application...")
    response, err := http.Get("https://crmconnection-a2be7e59b.dispatcher.hana.ondemand.com/ZUHD_INCIDENT_HANDLER_SRV/Incidents")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
    jsonValue, _ := json.Marshal(jsonData)
    response, err = http.Post("https://crmconnection-a2be7e59b.dispatcher.hana.ondemand.com/ZUHD_INCIDENT_HANDLER_SRV/Incidents('20026563')?$format=json", "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    fmt.Println("Terminating the application...")
}
