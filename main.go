package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func main() {
	//URL := "localhost:8080/bk/api/actions.json?distinct_id=22397_1495093020959_3130&timestamp=1495093225&event=update_cart_source&params%5Bcart_token%5D=4edf5bb6c3bd13844a95614e7cfecb10&params%5Bsources%5D%5B%5D=mobilewebboost_conversion_driver&params%5Bsources%5D%5B%5D=mobilewebboost_app&params%5Bsources%5D%5B%5D=mobilewebboost_web&params%5BclientDevice%5D=mobile&params%5BclientIsMobile%5D=true&params%5BclientIsSmallScreen%5D=true&api_key=8eb697157e35c4b20e158328de64aac7"
	URL := "localhost:8080/bk/api/actions.json?distinct_id=763_1495187301909_3495&timestamp=1495187523&event=product_add_cart&params%5BproductRefId%5D=8284563078&params%5Bapps%5D%5B%5D=precommend&params%5Bapps%5D%5B%5D=bsales&params%5Bsource%5D=item&params%5Boptions%5D%5Bsegment%5D=cart_recommendation&params%5Boptions%5D%5Btype%5D=up_sell&params%5BtimeExpire%5D=1495187599642&params%5Brecommend_system_product_source%5D=item&params%5Bproduct_id%5D=8284563078&params%5Bvariant_id%5D=27661944134&params%5Bsku%5D=00483332%20(black)&params%5Bsources%5D%5B%5D=product_recommendation&params%5Bcart_token%5D=dc2c336a009edf2762128e65806dfb1d&params%5Bquantity%5D=1&params%5Bnew_popup_upsell_mobile%5D=false&params%5BclientDevice%5D=desktop&params%5BclientIsMobile%5D=false&params%5BclientIsSmallScreen%5D=false&params%5Bnew_popup_crossell_mobile%5D=false&api_key=14c5b7dacea9157029265b174491d340"

	// Parse URL string
	u, _ := url.Parse(URL)

	// Parse get query
	query, _ := url.ParseQuery(u.RawQuery)

	content := make(map[string]interface{})
	params := make(map[string]interface{})

	for k, v := range query {
		detectKey(k, v)
		if strings.Contains(k, "params") {
			key := removeString(k, []string{"params", "[", "]"})
			value := getValue(v)
			params[key] = value
		} else {
			// we have only key - value
			key := removeString(k, []string{"params", "[", "]"})
			content[key] = v[0]
		}
	}
	content["params"] = params
	//contenJSON, _ := json.Marshal(content)
	//fmt.Println(string(contenJSON))
}

func Parse(request *http.Request) string {

	query := request.URL.Query()

	content := make(map[string]interface{})
	params := make(map[string]interface{})

	for k, v := range query {
		//detectKey(k, v)
		if strings.Contains(k, "params") {
			key := removeString(k, []string{"params", "[", "]"})
			value := getValue(v)
			params[key] = value
		} else {
			// we have only key - value
			key := removeString(k, []string{"params", "[", "]"})
			content[key] = v[0]
		}
	}
	content["params"] = params
	contenJSON, _ := json.Marshal(content)
	return string(contenJSON)
}

func getValue(data interface{}) interface{} {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Slice {
		if v.Len() == 1 {
			first, _ := v.Interface().([]string)
			return first[0]
		}
		return data
	}
	return ""
}

// Get key
func removeString(src string, args []string) string {
	temp := src
	for _, v := range args {
		temp = strings.Replace(temp, v, "", -1) // -1 to replace all position if it's exists
	}
	return temp
}

func detectKey(key string, value interface{}) interface{} {
	fmt.Println(key)
	// rgx := regexp.MustCompile(`\[(.*?)\]`)
	// rs := rgx.FindAllStringSubmatch(key, -1)
	// for _, v := range rs {
	// 	fmt.Println(v[1])
	// }
	return ""
}
