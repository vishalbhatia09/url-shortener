package common

import "time"

type BrandlyResponseBody struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Slashtag    string        `json:"slashtag"`
	Destination string        `json:"destination"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Status      string        `json:"status"`
	Tags        []interface{} `json:"tags"`
	Clicks      int           `json:"clicks"`
	IsPublic    bool          `json:"isPublic"`
	ShortURL    string        `json:"shortUrl"`
	DomainID    string        `json:"domainId"`
	DomainName  string        `json:"domainName"`
	Domain      Domain        `json:"domain"`
	HTTPS       bool          `json:"https"`
	Favourite   bool          `json:"favourite"`
	Creator     Creator       `json:"creator"`
	Integrated  bool          `json:"integrated"`
}
type Protocol struct {
	Allowed []string `json:"allowed"`
	Default string   `json:"default"`
}
type Sharing struct {
	Protocol Protocol `json:"protocol"`
}
type Domain struct {
	ID       string  `json:"id"`
	Ref      string  `json:"ref"`
	FullName string  `json:"fullName"`
	Sharing  Sharing `json:"sharing"`
	Active   bool    `json:"active"`
}
type Creator struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	AvatarURL string `json:"avatarUrl"`
}

type DomainName struct {
	FullName string `json:"fullName"`
}

type BrandlyRequestBody struct {
	Destination string     `json:"destination"`
	Domain      DomainName `json:"domain"`
}

type RequestBody struct {
	OriginalUrl string `json:"original_url"`
}
