package steamapi

import (
	"encoding/json"
	"strconv"
)

/* --[[ Get Global Trade History ]]-- */

type GlobalTrades struct {
	Response struct {
		More   bool `json:"more"`
		Trades []struct {
			Tradeid        string `json:"tradeid"`
			SteamidOther   string `json:"steamid_other"`
			TimeInit       int    `json:"time_init"`
			Status         int    `json:"status"`
			AssetsReceived []struct {
				Appid        int    `json:"appid"`
				Contextid    string `json:"contextid"`
				Assetid      string `json:"assetid"`
				Amount       string `json:"amount"`
				Classid      string `json:"classid"`
				Instanceid   string `json:"instanceid"`
				NewAssetid   string `json:"new_assetid"`
				NewContextid string `json:"new_contextid"`
			} `json:"assets_received,omitempty"`
			TimeMod     int `json:"time_mod,omitempty"`
			AssetsGiven []struct {
				Appid        int    `json:"appid"`
				Contextid    string `json:"contextid"`
				Assetid      string `json:"assetid"`
				Amount       string `json:"amount"`
				Classid      string `json:"classid"`
				Instanceid   string `json:"instanceid"`
				NewAssetid   string `json:"new_assetid"`
				NewContextid string `json:"new_contextid"`
			} `json:"assets_given,omitempty"`
			TimeEscrowEnd int `json:"time_escrow_end,omitempty"`
		} `json:"trades"`
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
