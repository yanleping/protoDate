package util

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	SIZE_320x50     = 101
	SIZE_300x250    = 102
	SIZE_480x320    = 103
	SIZE_320x480    = 104
	SIZE_300x300    = 105
	SIZE_1200x627   = 106
	JS_TAG_320x50   = 107
	JS_TAG_300x250  = 108
	JS_TAG_480x320  = 109
	JS_TAG_320x480  = 110
	JS_TAG_300x300  = 111
	JS_TAG_1200x627 = 112
	VIDEO           = 201
	JS_TAG          = 202
	BRAND           = 301
	APP_NAME        = 401
	APP_DESC        = 402
	APP_RATE        = 403
	CTA_BUTTON      = 404
	ICON            = 405
	COMMENT         = 406
	SIZE_288x150    = 11001
	SIZE_428x238    = 11004
	SIZE_690x388    = 11009
	SIZE_710x396    = 11024
	ENDCARD         = 50001
	PLAYABLE_URL    = 61001
	PLAYABLE_ZIP    = 61002
)

var CreativeTypes = []int{
	SIZE_320x50,
	SIZE_300x250,
	SIZE_480x320,
	SIZE_320x480,
	SIZE_300x300,
	SIZE_1200x627,
	JS_TAG_320x50,
	JS_TAG_300x250,
	JS_TAG_480x320,
	JS_TAG_320x480,
	JS_TAG_300x300,
	JS_TAG_1200x627,
	VIDEO,
	JS_TAG,
	BRAND,
	APP_NAME,
	APP_DESC,
	APP_RATE,
	CTA_BUTTON,
	ICON,
	COMMENT,
	SIZE_288x150,
	SIZE_428x238,
	SIZE_690x388,
	SIZE_710x396,
	ENDCARD,
	PLAYABLE_URL,
	PLAYABLE_ZIP,
}

// CreativeInfo struct.
type CreativeInfo struct {
	ID               bson.ObjectId          `bson:"_id,omitempty" json:"_id,omitempty"`
	CampaignID       *int64                 `bson:"campaignId,omitempty" json:"campaignId"`
	Status           int32                  `bson:"status,omitempty" json:"status"`
	Source           *int32                 `bson:"source,omitempty" json:"source"`
	PackageName      *string                `bson:"packageName,omitempty" json:"packageName"`
	CountryCode      *string                `bson:"countryCode,omitempty" json:"countryCode"`
	SourceCreativeID *int32                 `bson:"sourceCreativeId,omitempty" json:"sourceCreativeId"`
	Updated          int64                  `bson:"updated,omitempty" json:"updated,omitempty"`
	UpdatedDate      *string                `bson:"updatedDate,omitempty" json:"updatedDate"`
	Content          map[string]interface{} `bson:"content,omitempty" json:"content"`
}

type Content struct {
	Url             string `bson:"url,omitempty" json:"url,omitempty"`
	VideoLength     int32  `bson:"videoLength,omitempty" json:"videoLength,omitempty"`
	VideoSize       int32  `bson:"videoSize,omitempty" json:"videoSize,omitempty"`
	VideoResolution string `bson:"videoResolution,omitempty" json:"videoResolution,omitempty"`
	Width           int32  `bson:"width,omitempty" json:"width,omitempty"`
	Height          int32  `bson:"height,omitempty" json:"height,omitempty"`
	//VideoTruncation int32  `bson:"videoTruncation,omitempty" json:"videoTruncation,omitempty"`
	WatchMile int32 `bson:"watchMile,omitempty" json:"watchMile,omitempty"`
	BitRate   int32 `bson:"bitrate,omitempty" json:"bitrate,omitempty"`
	//ScreenShot      string `bson:"screenShot,omitempty" json:"screenShot,omitempty"`
	Resolution string      `bson:"resolution,omitempty" json:"resolution,omitempty"`
	Value      interface{} `bson:"value,omitempty" json:"value,omitempty"`
	Mime       string      `bson:"mime,omitempty" json:"mime,omitempty"`
	//Attribute     *string `bson:"attribute,omitempty" json:"attribute,omitempty"`
	AdvCreativeId string `bson:"advCreativeId,omitempty" json:"advCreativeId,omitempty"`
	CreativeId    *int64 `bson:"creativeId,omitempty" json:"creativeId,omitempty"`
	Source        *int32 `bson:"source,omitempty" json:"source,omitempty"`
	FMd5          string `bson:"fMd5,omitempty" json:"fMd5,omitempty`
}
