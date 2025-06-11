package main

import (
	"fmt"

	SteamAPI "github.com/Toakley683/GoLang-SteamAPI"
)

func main() {

	Context := SteamAPI.SetSteamAPIContext(
		&SteamAPI.APIKey{
			Key: "API KEY HERE",
		},
	)

	TF2 := SteamAPI.AppInformation{
		AppID: 440,
	}

	Toakley682 := SteamAPI.ClientInformation{
		SteamID64: "76561198170087194",
	}

	fmt.Println("Game:", TF2)
	fmt.Println("User:", Toakley682)
	fmt.Println("Context:", Context)

	TF2.GetLatestNewsForApp(3, 300)
	TF2.GetGlobalAchievementsForApp()
	TF2.GetAssetPrices()

	Toakley682.GetProfileSummary()
	Toakley682.GetFriendsList(SteamAPI.RELATIONSHIP_FRIEND)
	Toakley682.GetAppAchievements(TF2)
	Toakley682.GetGameStats(TF2)
	Toakley682.GetOwnedGames()
	Toakley682.GetRecentGames()

	Context.GetAppList()
	Context.GetServersAtAddress(SteamAPI.Address{
		Address: "192.168.0.32",
		Port:    27075,
	})
	Context.GetGlobalTradeHistory(SteamAPI.GlobalTradeContext{
		MaxTrades:          30,
		TradesAfterTime:    0,    // Unix Time
		TradesAfterTradeID: 0,    // TradeID
		NavigateBack:       true, // Show last page
		GetDescriptions:    true,
		IncludeTotal:       true,
	})

}
