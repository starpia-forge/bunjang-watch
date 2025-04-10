package bunjang

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	QueryUrl = "https://api.bunjang.co.kr/api/1/find_v2.json"
)

type Request struct {
	Query                string `json:"q"`                      // Search Text
	Order                string `json:"order"`                  // score
	Page                 int    `json:"page"`                   // 0
	RequestID            string `json:"request_id"`             // 2025409003205 (2025-04-09_00:32:05 Fired Request)
	StatDevice           string `json:"stat_device"`            // w
	N                    int    `json:"n"`                      // 100
	StatCategoryRequired int    `json:"stat_category_required"` // 1
	ReqRef               string `json:"req_ref"`                // search
	Version              int    `json:"version"`                // 5
}

type Response struct {
	Result          string        `json:"result"`
	NoResult        bool          `json:"no_result"`
	NoResultType    interface{}   `json:"no_result_type"`
	NoResultMessage interface{}   `json:"no_result_message"`
	AdMoreInfo      bool          `json:"ad_more_info"`
	AdProducts      []interface{} `json:"ad_products"`
	List            []Product     `json:"list"`
}

const (
	ProductStatusOnSale   = 0
	ProductStatusReserved = 1
	ProductStatusSoldOut  = 3
)

const (
	ProductUsedLightly = 1
	ProductUsedNew     = 2
)

type Product struct {
	Pid              string      `json:"pid"`
	Name             string      `json:"name"`
	Price            string      `json:"price"`
	ProductImage     string      `json:"product_image"`
	Status           string      `json:"status"`
	Ad               bool        `json:"ad"`
	Inspection       string      `json:"inspection"`
	Care             bool        `json:"care"`
	Location         string      `json:"location"`
	Badges           interface{} `json:"badges"`
	NamePrefix       interface{} `json:"name_prefix"`
	Bizseller        bool        `json:"bizseller"`
	Checkout         bool        `json:"checkout"`
	ContactHope      bool        `json:"contact_hope"`
	FreeShipping     bool        `json:"free_shipping"`
	IsAdult          bool        `json:"is_adult"`
	IsSuperUpShop    interface{} `json:"is_super_up_shop"`
	MaxCpc           interface{} `json:"max_cpc"`
	NumComment       string      `json:"num_comment"`
	NumFaved         string      `json:"num_faved"`
	OnlyNeighborhood bool        `json:"only_neighborhood"`
	OutlinkUrl       string      `json:"outlink_url"`
	PuId             interface{} `json:"pu_id"`
	Style            string      `json:"style"`
	SuperUp          interface{} `json:"super_up"`
	Tag              string      `json:"tag"`
	Uid              string      `json:"uid"`
	UpdateTime       int         `json:"update_time"`
	Used             int         `json:"used"`
	Proshop          bool        `json:"proshop"`
	CategoryId       string      `json:"category_id"`
	RefCampaign      interface{} `json:"ref_campaign"`
	RefCode          interface{} `json:"ref_code"`
	RefMedium        interface{} `json:"ref_medium"`
	RefContent       string      `json:"ref_content"`
	RefSource        string      `json:"ref_source"`
	RefTest          interface{} `json:"ref_test"`
	Tracking         struct {
		ImpId string `json:"imp_id"`
	} `json:"tracking"`
	AdRef  string      `json:"ad_ref"`
	Faved  bool        `json:"faved"`
	Type   string      `json:"type"`
	AppUrl interface{} `json:"app_url"`
	ImpId  string      `json:"imp_id"`
}

func Query(ctx context.Context) ([]Product, error) {
	u, err := url.Parse(QueryUrl)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := Response{}
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.List, nil
}
