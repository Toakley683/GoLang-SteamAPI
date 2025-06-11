## This is my GoLang Steam WebAPI library

This currently supports all API Endpoints in https://partner.steamgames.com/doc/webapi which use the api.steampowered.com endpoint URL

Installation:<br>
```go get github.com/Toakley683/GoLang-SteamAPI```

<br>

Example:<br>
```https://github.com/Toakley683/GoLang-SteamAPI/blob/main/example/example.go```

<br>

# API Endpoints Supported:

## App/Game Endpoints
<br>

`AppInformation.GetLatestNewsForApp(NewsCount, MaxLength)` - Gets Latest News of AppID <br><br>
`AppInformation.GetGlobalAchievementsForApp()` - Gets Global Achievement Information of AppID <br><br>
`AppInformation.GetAssetPrices()` - Gets all assets sold by that game, and the prices of those assets <br><br>

## User Endpoints
<br>

`ClientInformation.GetProfileSummary()` - Gets User Profile Summaries <br><br>
`ClientInformation.GetFriendsList( RelationshipType )` - Gets User's Friend List [Must be PUBLIC to see!] <br><br>
`ClientInformation.GetAppAchievements( AppInformation )` - Gets User achievements in requested AppID <br><br>
`ClientInformation.GetGameStats( AppInformation )` - Gets User game stats in requested AppID <br><br>
`ClientInformation.GetOwnedGames()` - Gets User's Owned Games [Must be PUBLIC to see!] <br><br>
`ClientInformation.GetRecentGames()` - Gets User's Recently Played Games <br><br>

## Global Endpoints
<br>

`GlobalContext.GetAppList()` - Gets all games which the steam store has <br><br>
`GlobalContext.GetServersAtAddress( Address )` - Gets game servers running at X address <br><br>
`GlobalContext.GetGlobalTradeHistory( GlobalTradeContext )` - Gets global trade history between players <br><br>
