package api

import (
    "bytes" 
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

const baseURL string = "https://www.myservice.com/v1"
 
type Client struct {
	Username string
	Password string
}
 
func NewBasicAuthC lient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

//type AutoGenerated struct {
	Incident struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ObjectID          string `json:"ObjectId"`
		ProcessTypeID     string `json:"ProcessTypeId"`
		StatusID          string `json:"StatusId"`
		PriorityID        string `json:"PriorityId"`
		Title             string `json:"Title"`
		ResponsibleID     string `json:"ResponsibleId"`
		AssignmentGroupID string `json:"AssignmentGroupId"`
		AssignmentGroup   struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AssignmentGroup"`
                ProcessType   struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"ProcessType"`
		Status struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"Status"`
		Priority struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"Priority"`
		Responsible struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"Responsible"`
	} `json:"Incident"`
//}
func (s *Client) GetIncident(id int) (*Todo, error) {
	
	//url := fmt.Sprintf(baseURL+"/%s/todos/%d", s.Username, id)
	url := fmt.Sprintf(baseURL+"Incidents('20026563')?$format=json")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Incident
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

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
	// https://www.scaledrone.com/blog/creating-an-api-client-in-go/
	client := NewBasicAuthClient("username", "password")
	t, _ := client.GetIncident(1)
	
/*	
type FamilyMember struct {
    Name    string
    Age     int
    Parents []string
}

    var m FamilyMember
    err := json.Unmarshal(b, &m)
  
	
  fmt.Println("Starting the application...")
    response, err := http.Get("")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
    jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://yourdomaingoeshere.com/crm", "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
    fmt.Println("Terminating the application...")
    */
}
