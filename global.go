package steamapi

import (
	"encoding/json"
	"strconv"
)

/* --[[ Get Global Trade History ]]-- */

type AssetsReceived struct {
	AppID        int    `json:"appid"`
	Contextid    string `json:"contextid"`
	Assetid      string `json:"assetid"`
	Amount       string `json:"amount"`
	Classid      string `json:"classid"`
	Instanceid   string `json:"instanceid"`
	NewAssetid   string `json:"new_assetid"`
	NewContextid string `json:"new_contextid"`
}

type AssetsGiven struct {
	Appid        int    `json:"appid"`
	Contextid    string `json:"contextid"`
	Assetid      string `json:"assetid"`
	Amount       string `json:"amount"`
	Classid      string `json:"classid"`
	Instanceid   string `json:"instanceid"`
	NewAssetid   string `json:"new_assetid"`
	NewContextid string `json:"new_contextid"`
}

type Trade struct {
	Tradeid        string           `json:"tradeid"`
	SteamidOther   string           `json:"steamid_other"`
	TimeInit       int              `json:"time_init"`
	Status         int              `json:"status"`
	AssetsReceived []AssetsReceived `json:"assets_received,omitempty"`
	TimeMod        int              `json:"time_mod,omitempty"`
	AssetsGiven    []AssetsGiven    `json:"assets_given,omitempty"`
	TimeEscrowEnd  int              `json:"time_escrow_end,omitempty"`
}

type GlobalTrades struct {
	Response struct {
		More   bool    `json:"more"`
		Trades []Trade `json:"trades"`
	} `json:"response"`
}

func (Client *GlobalContext) GetGlobalTradeHistory(trade GlobalTradeContext) (*GlobalTrades, error) {

	maxTrades := strconv.Itoa(int(trade.MaxTrades))
	tradesAfterXTime := strconv.Itoa(int(trade.TradesAfterTime))
	tradesAfterXTradeID := strconv.Itoa(int(trade.TradesAfterTradeID))
	navigateBack := strconv.FormatBool(trade.NavigateBack)
	getDescriptions := strconv.FormatBool(trade.GetDescriptions)
	language := trade.Language
	includeTotal := strconv.FormatBool(trade.IncludeTotal)

	ReqURL := "https://api.steampowered.com/IEconService/GetTradeHistory/v1/"

	ReqURL = ReqURL + "?key=" + globalContext.APIKey.Key
	ReqURL = ReqURL + "&max_trades=" + maxTrades
	ReqURL = ReqURL + "&start_after_time=" + tradesAfterXTime
	ReqURL = ReqURL + "&start_after_tradeid=" + tradesAfterXTradeID
	ReqURL = ReqURL + "&navigating_back=" + navigateBack
	ReqURL = ReqURL + "&get_descriptions=" + getDescriptions
	ReqURL = ReqURL + "&language=" + language.data
	ReqURL = ReqURL + "&include_total=" + includeTotal

	Data, err := requestAPI(ReqURL)
	if err != nil {
		return nil, err
	}

	Response := GlobalTrades{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	return &Response, nil

}

/* --[[ Get Server at Address ]]-- */

type Server struct {
	Addr          string `json:"addr"`
	Gmsindex      int    `json:"gmsindex"`
	ServerSteamID string `json:"steamid"`
	AppID         int    `json:"appid"`
	GameDir       string `json:"gamedir"`
	Region        int    `json:"region"`
	Secure        bool   `json:"secure"`
	Lan           bool   `json:"lan"`
	GamePort      int    `json:"gameport"`
	SpecPort      int    `json:"specport"`
}
type GameServer struct {
	Data struct {
		Success bool     `json:"success"`
		Servers []Server `json:"servers"`
	} `json:"response"`
}

func (Client *GlobalContext) GetServersAtAddress(address Address) (*GameServer, error) {

	Addr := address.Address
	Port := strconv.Itoa(int(address.Port))

	Data, err := requestAPI("https://api.steampowered.com/ISteamApps/GetServersAtAddress/v1/?addr=" + Addr + ":" + Port)
	if err != nil {
		return nil, err
	}

	Response := GameServer{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	return &Response, nil

}

/* --[[ Get Global App List ]]-- */

type GlobalApps struct {
	Data struct {
		Apps []struct {
			AppID int    `json:"appid"`
			Name  string `json:"name"`
		} `json:"apps"`
	} `json:"applist"`
}

func (Client *GlobalContext) GetAppList() (*GlobalApps, error) {

	Data, err := requestAPI("https://api.steampowered.com/ISteamApps/GetAppList/v2/?format=json")
	if err != nil {
		return nil, err
	}

	Response := GlobalApps{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	return &Response, nil

}
