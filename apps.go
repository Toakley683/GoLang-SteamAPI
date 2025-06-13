package steamapi

import (
	"encoding/json"
	"strconv"
)

/* --[[ Request Latest [AppID] Game News ]]-- */

type NewsResponse struct {
	Data struct {
		AppID     int `json:"appid"`
		Newsitems []struct {
			Gid           string `json:"gid"`
			Title         string `json:"title"`
			URL           string `json:"url"`
			IsExternalURL bool   `json:"is_external_url"`
			Author        string `json:"author"`
			Contents      string `json:"contents"`
			Feedlabel     string `json:"feedlabel"`
			Date          int    `json:"date"`
			Feedname      string `json:"feedname"`
			FeedType      int    `json:"feed_type"`
			AppID         int    `json:"appid"`
		} `json:"newsitems"`
		Count int `json:"count"`
	} `json:"appnews"`
}

func (App *AppInformation) GetLatestNewsForApp(newsCount, maxLength uint32) (*NewsResponse, error) {

	appID := strconv.Itoa(int(App.AppID))
	newsC := strconv.Itoa(int(newsCount))
	maxL := strconv.Itoa(int(maxLength))

	Data, err := requestAPI("http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002/?appid=" + appID + "&count=" + newsC + "&maxlength=" + maxL + "&format=json")
	if err != nil {
		return nil, err
	}

	Response := NewsResponse{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	return &Response, nil

}

/* --[[ Request Latest [AppID] Global Achievements ]]-- */

type Achievement struct {
	Name    string `json:"name"`
	Percent string `json:"percent"`
}

type GlobalAchievementResponse struct {
	Data struct {
		Achievements []Achievement `json:"achievements"`
	} `json:"achievementpercentages"`
}

func (App *AppInformation) GetGlobalAchievementsForApp() (*GlobalAchievementResponse, error) {

	appID := strconv.Itoa(int(App.AppID))

	Data, err := requestAPI("https://api.steampowered.com/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002/?gameid=" + appID + "&format=json")
	if err != nil {
		return nil, err
	}

	Response := GlobalAchievementResponse{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	return &Response, nil

}

/* --[[ Request AppID's asset prices ]]-- */

type Prices struct {
	USD int `json:"USD"`
	GBP int `json:"GBP"`
	EUR int `json:"EUR"`
	RUB int `json:"RUB"`
	BRL int `json:"BRL"`
	JPY int `json:"JPY"`
	NOK int `json:"NOK"`
	IDR int `json:"IDR"`
	MYR int `json:"MYR"`
	PHP int `json:"PHP"`
	SGD int `json:"SGD"`
	THB int `json:"THB"`
	VND int `json:"VND"`
	KRW int `json:"KRW"`
	UAH int `json:"UAH"`
	MXN int `json:"MXN"`
	CAD int `json:"CAD"`
	AUD int `json:"AUD"`
	NZD int `json:"NZD"`
	PLN int `json:"PLN"`
	CHF int `json:"CHF"`
	CNY int `json:"CNY"`
	TWD int `json:"TWD"`
	HKD int `json:"HKD"`
	INR int `json:"INR"`
	AED int `json:"AED"`
	SAR int `json:"SAR"`
	ZAR int `json:"ZAR"`
	COP int `json:"COP"`
	PEN int `json:"PEN"`
	CLP int `json:"CLP"`
	CRC int `json:"CRC"`
	ILS int `json:"ILS"`
	KWD int `json:"KWD"`
	QAR int `json:"QAR"`
	UYU int `json:"UYU"`
	KZT int `json:"KZT"`
}

type Tags struct {
	Cosmetics    string `json:"Cosmetics"`
	Tools        string `json:"Tools"`
	Weapons      string `json:"Weapons"`
	Taunts       string `json:"Taunts"`
	ClassBundles string `json:"Class_Bundles"`
	Bundles      string `json:"Bundles"`
	Limited      string `json:"Limited"`
	Highlighted  string `json:"Highlighted"`
	Maps         string `json:"Maps"`
	New          string `json:"New"`
	Halloween    string `json:"Halloween"`
}

type TagIDs struct {
	Num0  int64 `json:"0"`
	Num1  int64 `json:"1"`
	Num2  int   `json:"2"`
	Num3  int   `json:"3"`
	Num4  int   `json:"4"`
	Num5  int64 `json:"5"`
	Num6  int64 `json:"6"`
	Num7  int   `json:"7"`
	Num8  int   `json:"8"`
	Num9  int   `json:"9"`
	Num10 int   `json:"10"`
}

type Asset struct {
	Prices Prices `json:"prices"`
	Name   string `json:"name"`
	Date   string `json:"date"`
	Class  []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"class"`
	Classid        string   `json:"classid"`
	Tags           []string `json:"tags"`
	TagIds         []int64  `json:"tag_ids"`
	OriginalPrices Prices   `json:"original_prices,omitempty"`
}

type AssetPrices struct {
	Result struct {
		Success bool    `json:"success"`
		Assets  []Asset `json:"assets"`
		Tags    Tags    `json:"tags"`
		TagIds  TagIDs  `json:"tag_ids"`
	} `json:"result"`
}

func (App *AppInformation) GetAssetPrices() (*AssetPrices, error) {

	appID := strconv.Itoa(int(App.AppID))

	Data, err := requestAPI("https://api.steampowered.com/ISteamEconomy/GetAssetPrices/v1/?key=" + globalContext.APIKey.Key + "&appid=" + appID + "&format=json")
	if err != nil {
		return nil, err
	}

	Response := AssetPrices{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	return &Response, nil

}
