package steamapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

/* --[[ Errors ]]-- */

const (
	NO_VALID_STEAM_CONTEXT = " The steam context provided was nil"
	NO_VALID_RESPONSE      = " No valid SteamAPI response"
	API_TIMEOUT            = " Too many SteamAPI Requests"
)

/* --[[ Type Structs ]]-- */

type Language struct {
	data string
}

type FriendListRelationship struct {
	data string
}

type GlobalTradeContext struct {
	MaxTrades          uint32
	TradesAfterTime    uint32
	TradesAfterTradeID uint64
	NavigateBack       bool
	GetDescriptions    bool
	Language           Language
	IncludeTotal       bool
}

type Address struct {
	Address string
	Port    uint32
}

/* --[[ Types ]]-- */

var (
	RELATIONSHIP_FRIEND = FriendListRelationship{data: "friend"}
	RELATIONSHIP_ALL    = FriendListRelationship{data: "all"}
)

/* --[[ Structs ]]-- */

type APIKey struct {
	Key string
}

type AppInformation struct {
	AppID uint32
}

type ClientInformation struct {
	SteamID64 string
}

type GlobalContext struct {
	APIKey APIKey
}

var (
	globalContext *GlobalContext
)

/* --[[ Context Setup ]]-- */

func SetSteamAPIContext(APIKey *APIKey) GlobalContext {

	globalContext = &GlobalContext{
		APIKey: *APIKey,
	}

	return *globalContext

}

/* --[[ API Requester ]]-- */

func requestAPI(url string) ([]byte, error) {

	if globalContext == nil {
		log.Println(NO_VALID_STEAM_CONTEXT)
		return nil, nil
	}

	fmt.Println("Requested:", "'"+url+"'")

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 429 {
		return nil, errors.New(API_TIMEOUT)
	}

	if response.StatusCode != 200 {
		return nil, errors.New(NO_VALID_RESPONSE)
	}

	Body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return Body, nil

}

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

type GlobalAchievementResponse struct {
	Data struct {
		Achievements []struct {
			Name    string `json:"name"`
			Percent string `json:"percent"`
		} `json:"achievements"`
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

type AssetPrices struct {
	Result struct {
		Success bool `json:"success"`
		Assets  []struct {
			Prices struct {
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
			} `json:"prices"`
			Name  string `json:"name"`
			Date  string `json:"date"`
			Class []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"class"`
			Classid        string   `json:"classid"`
			Tags           []string `json:"tags"`
			TagIds         []int64  `json:"tag_ids"`
			OriginalPrices struct {
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
			} `json:"original_prices,omitempty"`
		} `json:"assets"`
		Tags struct {
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
		} `json:"tags"`
		TagIds struct {
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
		} `json:"tag_ids"`
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

/* --[[ Request Player Summaries ]]-- */

type PlayerProfileSummary struct {
	Data struct {
		Players []struct {
			Steamid                  string `json:"steamid"`
			Communityvisibilitystate int    `json:"communityvisibilitystate"`
			Profilestate             int    `json:"profilestate"`
			Personaname              string `json:"personaname"`
			Commentpermission        int    `json:"commentpermission"`
			Profileurl               string `json:"profileurl"`
			Avatar                   string `json:"avatar"`
			Avatarmedium             string `json:"avatarmedium"`
			Avatarfull               string `json:"avatarfull"`
			Avatarhash               string `json:"avatarhash"`
			Lastlogoff               int    `json:"lastlogoff"`
			Personastate             int    `json:"personastate"`
			Primaryclanid            string `json:"primaryclanid"`
			Timecreated              int    `json:"timecreated"`
			Personastateflags        int    `json:"personastateflags"`
			Loccountrycode           string `json:"loccountrycode"`
		} `json:"players"`
	} `json:"response"`
}

func (Client *ClientInformation) GetProfileSummary() (*PlayerProfileSummary, error) {

	SteamID64 := Client.SteamID64

	Data, err := requestAPI("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=" + globalContext.APIKey.Key + "&steamids=" + SteamID64 + "&format=json")
	if err != nil {
		return nil, err
	}

	Response := PlayerProfileSummary{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	if len(Response.Data.Players) <= 0 {
		return nil, errors.New(NO_VALID_RESPONSE)
	}

	return &Response, nil

}

/* --[[ Request Player Friend's List ]]-- */

type PlayerFriends struct {
	Data struct {
		Friends []struct {
			Steamid      string `json:"steamid"`
			Relationship string `json:"relationship"`
			FriendSince  int    `json:"friend_since"`
		} `json:"friends"`
	} `json:"friendslist"`
}

func (Client *ClientInformation) GetFriendsList(relationship FriendListRelationship) (*PlayerFriends, error) {

	SteamID64 := Client.SteamID64

	Data, err := requestAPI("https://api.steampowered.com/ISteamUser/GetFriendList/v0001/?key=" + globalContext.APIKey.Key + "&steamid=" + SteamID64 + "&relationship=" + relationship.data + "&format=json")
	if err != nil {
		return nil, err
	}

	Response := PlayerFriends{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	if len(Response.Data.Friends) <= 0 {
		return nil, errors.New(NO_VALID_RESPONSE)
	}

	return &Response, nil

}

/* --[[ Request Player Achievement List for AppID ]]-- */

type PlayerAchievements struct {
	Data struct {
		SteamID      string `json:"steamID"`
		GameName     string `json:"gameName"`
		Achievements []struct {
			AchievementName string `json:"apiname"`
			Achieved        int    `json:"achieved"`
			Unlocktime      int    `json:"unlocktime"`
		} `json:"achievements"`
		Success bool `json:"success"`
	} `json:"playerstats"`
}

func (Client *ClientInformation) GetAppAchievements(app AppInformation) (*PlayerAchievements, error) {

	appID := strconv.Itoa(int(app.AppID))
	SteamID64 := Client.SteamID64

	Data, err := requestAPI("http://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v0001/?appid=" + appID + "&key=" + globalContext.APIKey.Key + "&steamid=" + SteamID64 + "&format=json")
	if err != nil {
		return nil, err
	}

	Response := PlayerAchievements{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	if !Response.Data.Success {
		return nil, errors.New(NO_VALID_RESPONSE)
	}

	return &Response, nil

}

/* --[[ Request Player Stats for AppID ]]-- */

type PlayerStats struct {
	Data struct {
		SteamID  string `json:"steamID"`
		GameName string `json:"gameName"`
		Stats    []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"stats"`
	} `json:"playerstats"`
}

func (Client *ClientInformation) GetGameStats(app AppInformation) (*PlayerStats, error) {

	appID := strconv.Itoa(int(app.AppID))
	SteamID64 := Client.SteamID64

	Data, err := requestAPI("https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?appid=" + appID + "&key=" + globalContext.APIKey.Key + "&steamid=" + SteamID64 + "&format=json")
	if err != nil {
		return nil, err
	}

	Response := PlayerStats{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	return &Response, nil

}

/* --[[ Request Owned Games for User ]]-- */

type PlayerGame struct {
	AppID                int `json:"appid"`
	Playtime             int `json:"playtime_forever"`
	PlaytimeOnWindows    int `json:"playtime_windows_forever"`
	PlaytimeOnMac        int `json:"playtime_mac_forever"`
	PlaytimeOnLinux      int `json:"playtime_linux_forever"`
	PlaytimeOnSteamDeck  int `json:"playtime_deck_forever"`
	TimeLastPlayed       int `json:"rtime_last_played"`
	PlaytimePast2Weeks   int `json:"playtime_2weeks,omitempty"`
	PlaytimeDisconnected int `json:"playtime_disconnected"`
}

type PlayerGamesList struct {
	Data struct {
		GameCount int          `json:"game_count"`
		Games     []PlayerGame `json:"games"`
	} `json:"response"`
}

func (Client *ClientInformation) GetOwnedGames() (*PlayerGamesList, error) {

	SteamID64 := Client.SteamID64

	Data, err := requestAPI("https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=" + globalContext.APIKey.Key + "&steamid=" + SteamID64 + "&format=json")
	if err != nil {
		return nil, err
	}

	Response := PlayerGamesList{}

	err = json.Unmarshal(Data, &Response)
	if err != nil {
		return nil, err
	}

	return &Response, nil

}

/* --[[ Request Recently Played for User ]]-- */

type PlayerRecentGames struct {
	Data struct {
		TotalCount int `json:"total_count"`
		Games      []struct {
			AppID               int    `json:"appid"`
			Name                string `json:"name"`
			ImgIconURL          string `json:"img_icon_url"`
			Playtime            int    `json:"playtime_forever"`
			PlaytimeOnWindows   int    `json:"playtime_windows_forever"`
			PlaytimeOnMac       int    `json:"playtime_mac_forever"`
			PlaytimeOnLinux     int    `json:"playtime_linux_forever"`
			PlaytimeOnSteamDeck int    `json:"playtime_deck_forever"`
			PlaytimePast2Weeks  int    `json:"playtime_2weeks,omitempty"`
		} `json:"games"`
	} `json:"response"`
}

func (Client *ClientInformation) GetRecentGames() (*PlayerRecentGames, error) {

	SteamID64 := Client.SteamID64

	Data, err := requestAPI("https://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v0001/?key=" + globalContext.APIKey.Key + "&steamid=" + SteamID64 + "&format=json")
	if err != nil {
		return nil, err
	}

	Response := PlayerRecentGames{}

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

/* --[[ Get Server at Address ]]-- */

type GameServer struct {
	Data struct {
		Success bool `json:"success"`
		Servers []struct {
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
		} `json:"servers"`
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
