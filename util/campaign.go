package util

type CampaignInfo struct {
	CampaignId         int64                        `bson:"campaignId,omitempty" json:"campaignId,omitempty"`
	AdvertiserId       *int32                       `bson:"advertiserId,omitempty" json:"advertiserId,omitempty"`
	TrackingUrl        *string                      `bson:"trackingUrl,omitempty" json:"trackingUrl,omitempty"`
	DirectUrl          *string                      `bson:"directUrl,omitempty" json:"directUrl,omitempty"`
	Price              *float64                     `bson:"price,omitempty" json:"price,omitempty"`
	OriPrice           *float64                     `bson:"oriPrice,omitempty" json:"oriPrice,omitempty"`
	//CityCode           interface{}                  `bson:"cityCode,omitempty" json:"cityCode,omitempty"`
	Status             int32                        `bson:"status,omitempty" json:"status,omitempty"`
	Network            *int32                       `bson:"network,omitempty" json:"network,omitempty"`
	PreviewUrl         *string                      `bson:"previewUrl,omitempty" json:"previewUrl,omitempty"`
	PackageName        *string                      `bson:"packageName,omitempty" json:"packageName,omitempty"`
	Updated            int64               `bson:"updated,omitempty" json:"updated,omitempty"`
	CampaignType       *int32                       `bson:"campaignType,omitempty" json:"campaignType,omitempty"`
	Ctype              *int32                       `bson:"ctype,omitempty" json:"ctype,omitempty"`
	AppSize            *string                      `bson:"appSize,omitempty" json:"appSize,omitempty"`
	Tag                *int32                       `bson:"tag,omitempty" json:"tag,omitempty"`
	AdSourceId         *int32                       `bson:"adSourceId,omitempty" json:"adSourceId,omitempty"`
	PublisherId        *int64                       `bson:"publisherId,omitempty" json:"publisherId,omitempty"`
	PreClickCacheTime  *int32                       `bson:"preClickCacheTime,omitempty" json:"preClickCacheTime,omitempty"`
	FrequencyCap       *int32                       `bson:"frequencyCap,omitempty" json:"frequencyCap,omitempty"`
	DirectPackageName  *string                      `bson:"directPackageName,omitempty" json:"directPackageName,omitempty"`
	SdkPackageName     *string                      `bson:"sdkPackageName,omitempty" json:"sdkPackageName,omitempty"`
	AdvImp             *[]AdvImp                    `bson:"advImp,omitempty" json:"advImp,omitempty"`
	AdUrlList          *[]string                    `bson:"adUrlList,omitempty" json:"adUrlList,omitempty"`
	JumpType           *int32                       `bson:"jumpType,omitempty" json:"jumpType,omitempty"`
	VbaConnecting      *int32                       `bson:"vbaConnecting,omitempty" json:"vbaConnecting,omitempty"`
	VbaTrackingLink    *string                      `bson:"vbaTrackingLink,omitempty" json:"vbaTrackingLink,omitempty"`
	RetargetingDevice  *int32                       `bson:"retargetingDevice,omitempty" json:"retargetingDevice,omitempty"`
	SendDeviceidRate   *int32                       `bson:"sendDeviceidRate,omitempty" json:"sendDeviceidRate,omitempty"`
	Endcard            *map[string]EndCard          `bson:"endcard,omitempty" json:"endcard,omitempty"`
	Loopback           *LoopBack                    `bson:"loopback,omitempty" json:"loopback,omitempty"`
	BelongType         *int32                       `bson:"belongType,omitempty" json:"belongType,omitempty"`
	ConfigVBA          *ConfigVBA                   `bson:"configVBA,omitempty" json:"configVBA,omitempty"`
	AppPostList        *AppPostList                 `bson:"appPostList,omitempty" json:"appPostList,omitempty"`
	BlackSubidListV2   map[string]map[string]string `bson:"blackSubidListV2,omitempty" json:"blackSubidListV2,omitempty"`
	BtV4               *BtV4                        `bson:"btV4,omitempty" json:"btV4,omitempty"`
	OpenType           *int32                       `bson:"openType,omitempty" json:"openType,omitempty"`
	SubCategoryName    *[]string                    `bson:"subCategoryName,omitempty" json:"subCategoryName,omitempty"`
	IsCampaignCreative *int32                       `bson:"isCampaignCreative,omitempty" json:"isCampaignCreative,omitempty"`
	CostType           *int32                       `bson:"costType,omitempty" json:"costType,omitempty"`
	Source             *int32                       `bson:"source,omitempty" json:"source,omitempty"`
	JUMPTYPECONFIG     *map[string]int32            `bson:"JUMP_TYPE_CONFIG,omitempty" json:"JUMP_TYPE_CONFIG,omitempty"`
	ChnID              *int                         `bson:"chnId,omitempty" json:"chnId,omitempty"`
	ThirdParty         string                       `bson:"thirdParty,omitempty" json:"thirdParty,omitempty"`
	TCQF               *TCQF                        `bson:"tcqf,omitempty" json:"tcqf,omitempty"`
}

type AppPostList struct {
	Include []string `bson:"include,omitempty" json:"include"`
	Exclude []string `bson:"exclude,omitempty" json:"exclude"`
}

type LoopBack struct {
	Domain *string `bson:"domain,omitempty" json:"domain,omitempty"`
	Key    *string `bson:"key,omitempty" json:"key,omitempty"`
	Value  *string `bson:"value,omitempty" json:"value,omitempty"`
	Rate   *int32  `bson:"rate,omitempty" json:"rate,omitempty"`
}

type EndCard struct {
	Urls             *[]EndCardUrls          `bson:"urls,omitempty" json:"urls,omitempty"`
	Status           *int32                  `bson:"status,omitempty" json:"status,omitempty"`
	Orientation      *int32                  `bson:"orientation,omitempty" json:"orientation,omitempty"`
	VideoTemplateUrl *[]VideoTemplateUrlItem `bson:"videoTemplateUrl,omitempty" json:"videoTemplateUrl,omitempty"`
	EndcardProtocal  *int                    `bson:"endcardProtocol,omitempty" json:"endcardProtocol,omitempty"`
	EndcardRate      *map[int]int            `bson:"endcardRate,omitempty" json:"endcardRate,omitempty"`
}

type VideoTemplateUrlItem struct {
	ID           *int32  `bson:"id,omitempty" json:"id,omitempty"`
	URL          *string `bson:"url,omitempty" json:"url,omitempty"`
	URLZip       *string `bson:"url_zip,omitempty" json:"url_zip,omitempty"`
	Weight       *int32  `bson:"weight,omitempty" json:"weight,omitempty"`
	PausedURL    *string `bson:"paused_url,omitempty" json:"paused_url,omitempty"`
	PausedURLZip *string `bson:"paused_url_zip,omitempty" json:"paused_url_zip,omitempty"`
}

type EndCardUrls struct {
	Id     *int32  `bson:"id,omitempty" json:"id,omitempty"`
	Url    *string `bson:"url,omitempty" json:"url,omitempty"`
	Weight *int32  `bson:"weight,omitempty" json:"weight,omitempty"`
	UrlV2  *string `bson:"url_v2,omitempty" json:"url_v2,omitempty"`
}

type AdvImp struct {
	Sec *int32  `bson:"sec,omitempty" json:"sec,omitempty"`
	Url *string `bson:"url,omitempty" json:"url,omitempty"`
}

type TCQF struct {
	SubIds map[string]SubInfo `bson:"subIds,omitempty" json:"subIds"`
}

type BtClass struct {
	Percent   float64 `bson:"percent,omitempty" json:"percent"`
	CapMargin int32   `bson:"capMargin,omitempty" json:"capMargin"`
	Status    int32   `bson:"status,omitempty" json:"status"`
}

type BtV4 struct {
	SubIds  map[string]SubInfo `bson:"subIds,omitempty" json:"subIds"`
	BtClass map[string]BtClass `bson:"btClass,omitempty" json:"btClass"`
}

type SubInfo struct {
	Rate        int                   `bson:"rate,omitempty" json:"rate"`
	PackageName string                `bson:"packageName,omitempty" json:"packageName"`
	DspSubIds   map[string]DspSubInfo `bson:"dspSubIds,omitempty" json:"dspSubIds"`
}

type DspSubInfo struct {
	Rate        int    `bson:"rate,omitempty" json:"rate"`
	PackageName string `bson:"packageName,omitempty" json:"packageName"`
}

type ConfigVBA struct {
	UseVBA       int `bson:"useVBA,omitempty" json:"useVBA"`
	FrequencyCap int `bson:"frequencyCap,omitempty" json:"frequencyCap"`
	Status       int `bson:"status,omitempty" json:"status"`
}
