package domain

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type URL struct {
	// Unique alias for redirection
	Alias string `json:"alias" bson:"_id,omitempty" maxLength:"8" example:"qwerty"`
	// Original URL
	Original string `json:"original" bson:"original" format:"valid URL" example:"https://google.com/"`
	// Time of creation
	CreatedAt time.Time `json:"createdAt" bson:"createdAt" format:"yyyy-MM-ddThh:mm:ss.ZZZ" example:"2021-05-09T09:29:18.169Z"`
	// Expiration time
	ExpiredAt time.Time `json:"expiredAt" bson:"expiredAt" format:"yyyy-MM-ddThh:mm:ss.ZZZ" example:"2021-06-09T09:29:18.169Z"`
	// Id of owner
	Owner primitive.ObjectID `json:"owner" bson:"owner" format:"hexadecimal string" example:"6095872d75ff40c9238bdb29"`
} // @name URL

type URLCreate struct {
	// Original URL
	Original string `json:"original" binding:"required,url" format:"valid URL" example:"https://google.com/"`
	// Duration of life of URL in seconds
	Duration int                `json:"duration" binding:"gte=0" example:"3600"`
	Owner    primitive.ObjectID `swaggerignore:"true"`
} // @name URLCreate

type URLProlong struct {
	// Duration of life of URL in seconds
	Duration int `json:"duration" binding:"gte=0" example:"3600"`
} // @name URLProlong

// NewURL create new URL from URLCreate and alias
func NewURL(toCreate URLCreate, alias string) URL {
	return URL{
		Alias:     alias,
		Original:  toCreate.Original,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Duration(toCreate.Duration) * time.Second),
		Owner:     toCreate.Owner,
	}
}

// MarshalBinary implement encoding.BinaryMarshaler for redis scanning
func (url URL) MarshalBinary() ([]byte, error) {
	return json.Marshal(url)
}

// UnmarshalBinary implement encoding.BinaryUnmarshaler for redis scanning
func (url *URL) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, url)
}

// Expired whether url is expired
func (url URL) Expired() bool {
	return url.ExpiredAt.Before(time.Now())
}
