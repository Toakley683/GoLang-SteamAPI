package steamapi

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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
