package matchbook

type LoginResponse struct {
	SessionToken string         `json:"session-token"`
	UserId       string         `json:"user-id"`
	Role         string         `json:"role"`
	Account      AccountDetails `json:"account"`
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phone-number"`
	Address      AddressDetails `json:"address"`
}

type AccountDetails struct {
	Id       int         `json:"id"`
	Username string      `json:"username"`
	Name     NameDetails `json:"name"`
}

type NameDetails struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

type BaseResponse struct {
	Total   int `json:"total"`
	PerPage int `json:"per-page"`
	Offset  int `json:"offset"`
}

type AddressDetails struct {
}

type SportsResult struct {
	BaseResponse
	Sports []Sport `json:"sports"`
}

type Sport struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Id   int    `json:"id"`
}

type EventsResult struct {
	BaseResponse
	Events []Event `json:"events"`
}

type Event struct {
	Id int `json:"id"`
	Name string `json:"name"`
	SportId int `json:"sport-id"`
	Start time.Date `json:"start"`
	InRunningFlag bool `json:"in-running-flag"`
	AllowLiveBetting bool `json:"allow-live-betting"`
	CategoryId []int `json:"category-id"`
	Status string `json:"status"`
	Volume float32 `json:"volume"`
	Markets []Market `json:"markets"`
	MetaTags []MetaTag `json:"meta-tags"`
}

type Market struct {
	Live book `json:"live"`
	EventId int `json:"event-id"`
	Id int `json:"id"`
	Name string `json:"name"`
	Runners []runner `json:"runners"`
	Start time.Date `json:"start"`
	InRunningFlag bool `json:"in-running-flag"`
	AllowLiveBetting bool `json:"allow-live-betting"`
	Status string `json:"status"`
	MarketType string `json:"market-type"`
	Type string `json:"type"`
	Volume float32 `json:"volume"`
	BackOverround float32 `json:"back-overround"`
	LayOverround float32 `json:"lay-overround"`
}

type Runner struct {
	Prices []PriceDetail `json:"prices"`
	EventId int `json:"event-id"`
	Id int `json:"id"`
	MarketId int `json:"market-id"`
	Name string `json:"name"`
	Status string `json:"status"`
	Volume float32 `json:"volume"`
	EventParticipantId int `json:"event-participant-id"`
}

type PriceDetail struct {
	AvailableAmount float32 `json:"available-amount"`
	Currency string `json:"currency"`
	OddsType string `json:"odds-type"`
	Odds float32 `json:"odds"`
	DecimalOdds float32 `json: decimal-odds`
	Side string `json:"side"`
	ExchangeType string `json:"exchange-type"`
}

func main{
	er := Sport{}

}
