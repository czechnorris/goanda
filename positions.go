package goanda

import "encoding/json"

// Supporting OANDA docs - http://developer.oanda.com/rest-live-v20/position-ep/

type OpenPositions struct {
	LastTransactionID string     `json:"lastTransactionID"`
	Positions         []Position `json:"positions"`
}

type Position struct {
	Instrument   string       `json:"instrument"`
	Long         PositionSide `json:"long"`
	Pl           string       `json:"pl"`
	ResettablePL string       `json:"resettablePL"`
	Short        PositionSide `json:"short"`
	UnrealizedPL string       `json:"unrealizedPL"`
}

type PositionSide struct {
	AveragePrice string   `json:"averagePrice"`
	Pl           string   `json:"pl"`
	ResettablePL string   `json:"resettablePL"`
	TradeIDs     []string `json:"tradeIDs"`
	Units        string   `json:"units"`
	UnrealizedPL string   `json:"unrealizedPL"`
}

type ClosePositionPayload struct {
	LongUnits  string `json:"longUnits,omitempty"`
	ShortUnits string `json:"shortUnits,omitempty"`
}

func (c *OandaConnection) GetOpenPositions() OpenPositions {
	endpoint := "/accounts/" + c.accountID + "/openPositions"

	response := c.Request(endpoint)
	data := OpenPositions{}
	unmarshalJson(response, &data)
	return data
}

func (c *OandaConnection) ClosePosition(instrument string, body ClosePositionPayload) ModifiedTrade {
	endpoint := "/accounts/" + c.accountID + "/positions/" + instrument + "/close"
	jsonBody, err := json.Marshal(body)
	checkErr(err)
	response := c.Update(endpoint, jsonBody)
	data := ModifiedTrade{}
	unmarshalJson(response, &data)
	return data
}
