package steamapi

import (
	"encoding/json"
	"errors"
	"strconv"
)

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
