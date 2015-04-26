package goblocktrail

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
    "bytes"
)

var NetworkList = []string{"tbtc", "btc"}

type API struct {
	apiKey        string
    network       string
    version       string
	client        *http.Client
}

func NewAPI(apiKey string) *API {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
    client := &http.Client{Transport: tr}

    return &API{apiKey, "BTC", "v1", client}
}

func (this *API) call(action, httpMethod string, params map[string]string) ([]byte, error) {

	var err error
	var res *http.Response

	if httpMethod=="POST" {

        jsonString, jsonError := json.Marshal(params)
        if jsonError!=nil {
            return nil, jsonError
        }

        req, _ := http.NewRequest("POST", "https://api.blocktrail.com/"+this.version+"/"+this.network+"/"+action+"?api_key="+this.apiKey, bytes.NewBuffer([]byte(jsonString)))
        res, err = this.client.Do(req)

	} else {
        valuesStr := "api_key="+this.apiKey
        for key, val := range params {
            valuesStr += "&"+key+"="+val
        }
		res, err = this.client.Get("https://api.blocktrail.com/"+this.version+"/"+this.network+"/"+action+"?"+valuesStr)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return body, err
}

func (this *API) SetNetwork(network string) bool{
    network = strings.ToLower(network)
    for _, n := range NetworkList {
        if n == network {
            this.network = network
            return true
        }
    }
    return false
}

func (this *API) SetVersion(version string) {
    this.version = version
}

func (this *API) Address(address string) (interface{}, error) {
    dataStream, err := this.call("address/"+address, "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) Transactions(address string, page, limit int, sort string) (interface{}, error) {
    if limit==0 {
        limit = 20
    }
    if sort != "asc"{
        sort = "desc"
    }
    var params = map[string]string{
        "page": strconv.Itoa(page),
        "limit": strconv.Itoa(limit),
        "sort_dir": sort,
    }
    dataStream, err := this.call("address/"+address+"/transactions", "GET", params)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) UnconfirmedTransactions(address string, page, limit int, sort string) (interface{}, error) {
    if limit==0 {
        limit = 20
    }
    if sort != "asc"{
        sort = "desc"
    }
    var params = map[string]string{
        "page": strconv.Itoa(page),
        "limit": strconv.Itoa(limit),
        "sort_dir": sort,
    }
    dataStream, err := this.call("address/"+address+"/unconfirmed-transactions", "GET", params)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) UnspentOutputs(address string, page, limit int, sort string) (interface{}, error) {
    if limit==0 {
        limit = 20
    }
    if sort != "asc"{
        sort = "desc"
    }
    var params = map[string]string{
        "page": strconv.Itoa(page),
        "limit": strconv.Itoa(limit),
        "sort_dir": sort,
    }
    dataStream, err := this.call("address/"+address+"/unspent-outputs", "GET", params)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) Block(block string) (interface{}, error) {
    dataStream, err := this.call("block/"+block, "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) BlockByHeight(height int) (interface{}, error) {
    return this.Block(strconv.Itoa(height))
}

func (this *API) BlockTransactions(block string, page, limit int, sort string) (interface{}, error) {
    if limit==0 {
        limit = 20
    }
    if sort != "asc"{
        sort = "desc"
    }
    var params = map[string]string{
        "page": strconv.Itoa(page),
        "limit": strconv.Itoa(limit),
        "sort_dir": sort,
    }
    dataStream, err := this.call("block/"+block+"/transactions", "GET", params)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) BlockTransactionsByHeight(height, page, limit int, sort string) (interface{}, error) {
    return this.BlockTransactions(strconv.Itoa(height), page, limit, sort)
}

func (this *API) AllBlocks( page, limit int, sort string) (interface{}, error) {
    if limit==0 {
        limit = 20
    }
    if sort != "asc"{
        sort = "desc"
    }
    var params = map[string]string{
        "page": strconv.Itoa(page),
        "limit": strconv.Itoa(limit),
        "sort_dir": sort,
    }
    dataStream, err := this.call("all-blocks", "GET", params)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) LatestBlock() (interface{}, error) {
    dataStream, err := this.call("block/latest", "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) Transaction(hash string) (interface{}, error) {
    dataStream, err := this.call("transaction/"+hash, "GET", nil)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    return data, err
}

func (this *API) VerifyMessage( message, address, signature string) (bool, error) {
    var params = map[string]string{
        "message": message,
        "address": address,
        "signature": signature,
    }
    dataStream, err := this.call("verify_message", "POST", params)
    data := map[string]interface{}{}
    json.Unmarshal(dataStream, &data)
    i, ok := data["result"]
    if ok && i==true {
        return true, err
    }
    return false, err
}

