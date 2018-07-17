package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	//"github.com/google/brotli/go/cbrotli"
	//"github.com/vmihailenco/msgpack"
	"github.com/json-iterator/go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/golang/protobuf/proto"
	"./protobuf"
	"./redis"
	"./util"
	//"gitlab.mobvista.com/ADN/test/protobuf"
)

const delay = 2 * time.Second

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	start := time.Now()

	//  设置并发度
	CORE_NUM := runtime.NumCPU() //number of core
	runtime.GOMAXPROCS(CORE_NUM - 1)

	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGTERM)
	signal.Notify(done, syscall.SIGINT)

	//  处理命令行参数
	//confPath := flag.String("c", "./conf/zeusextractor.conf", "configure file path")
	limit := flag.Int("limit", 200, "find limit, page size")
	findAll := flag.Bool("all", false, "find all")
	flag.Parse()

	// 读取配置文件
	//config := util.Config
	/*if err := config.LoadConfigFile(*confPath); err != nil {
		fmt.Println("load config failed: ", err.Error())
		os.Exit(1)
	}*/
	// redis
	if err := redis.InitRedis("127.0.0.1:6380"); err != nil {
		fmt.Println("init redis client failed: ", err.Error())
		//logger.Runtime.Error("init redis client failed: ", err.Error())
		//util.WriteDisk()
		os.Exit(1)
	}
	// mongodb
	session, err := mgo.Dial("adn-cpmongo-slave-sg.rayjump.com:27017")
	if err != nil {
		panic(err)
	}
	c := session.DB("new_adn").C("campaign")

	go func() {
		defer session.Close()
		sig := <-done
		fmt.Printf("caught sig: %+v", sig)
		//fmt.Println("Wait for 2 second to finish processing")
		//time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
	var item *mgo.Iter
	if *findAll {
		item = c.Find(bson.M{"status": 1}).Batch(200).Prefetch(0.25).Iter()
		flushALLCampaign(item)
	} else {
		totalDocs, err := c.Find(bson.M{"status": 1}).Count()
		if err != nil {
			panic(err)
		}
		for i := 0; i * *limit < totalDocs; i++ {
			//fmt.Println(i)
			item = c.Find(bson.M{"status": 1}).Skip(*limit * i).Limit(*limit).Batch(200).Prefetch(0.25).Iter()
			flushALLCampaign(item)
			fmt.Println("**********************************************")
			fmt.Printf("***** campaign all: [%d], flush [%d]. *****\n", totalDocs, *limit*i)
			fmt.Println("**********************************************")
			time.Sleep(delay)
			//break;
		}
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println("crtl + c exit processing")
	<-done
}

func fillCampaignPb(camInfo util.CampaignInfo) *protobuf.Campaign {
   camPb := &protobuf.Campaign {
	CampaignId: camInfo.CampaignId,
	Status:camInfo.Status,
   }
   if camInfo.AdvertiserId != nil {
	camPb.AdvertiserId = *camInfo.AdvertiserId
   }
   if camInfo.TrackingUrl != nil {
	camPb.TrackingUrl = *camInfo.TrackingUrl
   }
   if camInfo.DirectUrl != nil {
	camPb.DirectUrl = *camInfo.DirectUrl
   }
   if camInfo.Price != nil {
	camPb.Price = *camInfo.Price
   }
   if camInfo.OriPrice != nil {
	camPb.OriPrice = *camInfo.OriPrice
   }
   if camInfo.Network != nil {
	camPb.Network = *camInfo.Network
   }
   if camInfo.PreviewUrl != nil {
	camPb.PreviewUrl = *camInfo.PreviewUrl
   }
   if camInfo.PackageName != nil {
	camPb.PackageName = *camInfo.PackageName
   }
   if camInfo.CampaignType != nil {
	camPb.CampaignType = *camInfo.CampaignType
   }
   if camInfo.Ctype != nil {
	camPb.CType = *camInfo.Ctype
   }
   if camInfo.AppSize != nil {
	camPb.AppSize = *camInfo.AppSize
   }
   if camInfo.Tag != nil {
	camPb.Tag = *camInfo.Tag
   }
   if camInfo.AdSourceId != nil {
	camPb.AdSourceId = *camInfo.AdSourceId
   }
   if camInfo.PublisherId != nil {
	camPb.PublisherId = *camInfo.PublisherId
   }
   if camInfo.PreClickCacheTime != nil {
	camPb.PreClickCacheTime = *camInfo.PreClickCacheTime
   }
   if camInfo.FrequencyCap != nil {
	camPb.FrequencyCap = *camInfo.FrequencyCap
   }
   if camInfo.DirectPackageName != nil {
	camPb.DirectPackageName = *camInfo.DirectPackageName
   }
   if camInfo.SdkPackageName != nil {
	camPb.SdkPackageName = *camInfo.SdkPackageName
   }
   if camInfo.JumpType != nil {
	camPb.JumpType = *camInfo.JumpType
   }
   if camInfo.VbaConnecting != nil {
	camPb.VbaConnecting = *camInfo.VbaConnecting
   }
   if camInfo.VbaTrackingLink != nil {
	camPb.VbaTrackingLink = *camInfo.VbaTrackingLink
   }
   if camInfo.RetargetingDevice != nil {
	camPb.RetargetingDevice = *camInfo.RetargetingDevice
   }
   if camInfo.SendDeviceidRate != nil {
	camPb.SendDeviceidRate = *camInfo.SendDeviceidRate
   }
   if camInfo.BelongType != nil {
	camPb.BelongType = *camInfo.BelongType
   }
   if camInfo.OpenType != nil {
	camPb.OpenType = *camInfo.OpenType
   }
   if camInfo.IsCampaignCreative != nil {
	camPb.IsCampaignCreative = *camInfo.IsCampaignCreative
   }
   if camInfo.CostType != nil {
	camPb.CostType = *camInfo.CostType
   }
   if camInfo.Source != nil {
	camPb.Source = *camInfo.Source
   }
   if camInfo.ChnID != nil {
	camPb.ChnId = int32(*camInfo.ChnID)
   }
   if len(camInfo.ThirdParty) != 0 {
	camPb.ThirdParty = camInfo.ThirdParty
   }
   if camInfo.AdUrlList != nil && len(*camInfo.AdUrlList) > 0 {
	camPb.AdUrlList = *camInfo.AdUrlList
   }
   if camInfo.SubCategoryName != nil && len(*camInfo.SubCategoryName)>0 {
	camPb.SubCategoryName = *camInfo.SubCategoryName
   }
   if camInfo.JUMPTYPECONFIG != nil && len(*camInfo.JUMPTYPECONFIG) > 0 {
	camPb.JumpTypeConfig = *camInfo.JUMPTYPECONFIG
   }
   if camInfo.BlackSubidListV2 != nil && len(camInfo.BlackSubidListV2)>0 {
	camPb.BlackSubIdListV2 = make(map[string]*protobuf.MapValue, len(camInfo.BlackSubidListV2))
	for k, v := range camInfo.BlackSubidListV2 {
	    camPb.BlackSubIdListV2[k] = &protobuf.MapValue{MapValue:v}
	}
   }
   if camInfo.AdvImp != nil && len(*camInfo.AdvImp) > 0 {
	for _, advImp := range *camInfo.AdvImp {
	    camPb.AdvImp = append(camPb.AdvImp, &protobuf.AdvImp{Sec:*advImp.Sec, Url:*advImp.Url})
	}
   }
   if camInfo.Loopback != nil {
	camPb.LoopBack = &protobuf.LoopBack{
	    Domain: *camInfo.Loopback.Domain,
           Key : *camInfo.Loopback.Key,
           Value : *camInfo.Loopback.Value,
	    Rate : *camInfo.Loopback.Rate,
	}
   }
   if camInfo.AppPostList != nil {
	camPb.AppPostList = &protobuf.AppPostList {
	    Include : camInfo.AppPostList.Include,
           Exclude : camInfo.AppPostList.Exclude,
	}
   }
   if camInfo.ConfigVBA != nil {
	camPb.ConfigVBA = &protobuf.ConfigVBA{
	    UseVBA: int32(camInfo.ConfigVBA.UseVBA),
           FrequencyCap : int32(camInfo.ConfigVBA.FrequencyCap),
	    Status : int32(camInfo.ConfigVBA.Status),
	}
   }
   if camInfo.TCQF != nil && len(camInfo.TCQF.SubIds) > 0 {
	SubIds := make(map[string]*protobuf.SubInfo, len(camInfo.TCQF.SubIds))
	for k, v := range camInfo.TCQF.SubIds {
	     subInfo := protobuf.SubInfo{
		Rate : int32(v.Rate),
               PackageName : v.PackageName,
	     }
	     if len(v.DspSubIds) == 0 {
		SubIds[k] = &subInfo
		continue
	     }
	     subInfo.DspSubInfo = make(map[string]*protobuf.DspSubInfo, len(v.DspSubIds))
	     for dk, dv := range v.DspSubIds {
	         subInfo.DspSubInfo[dk] = &protobuf.DspSubInfo{
			Rate : int32(dv.Rate),
                       PackageName : dv.PackageName,
		 }
	     }
	     SubIds[k] = &subInfo
	}
	camPb.Tcqf = &protobuf.TCQF{
	    SubIds : SubIds,
	}
   }
   if camInfo.BtV4 != nil && (len(camInfo.BtV4.SubIds) > 0 || len(camInfo.BtV4.BtClass)>0) {
	btV4 := protobuf.BtV4{}
	if len(camInfo.BtV4.SubIds) > 0 {
	SubIds := make(map[string]*protobuf.SubInfo, len(camInfo.BtV4.SubIds))
	for k, v := range camInfo.BtV4.SubIds {
	     subInfo := protobuf.SubInfo{
		Rate : int32(v.Rate),
               PackageName : v.PackageName,
	     }
	     if len(v.DspSubIds) == 0 {
		SubIds[k] = &subInfo
		continue
	     }
	     subInfo.DspSubInfo = make(map[string]*protobuf.DspSubInfo, len(v.DspSubIds))
	     for dk, dv := range v.DspSubIds {
	         subInfo.DspSubInfo[dk] = &protobuf.DspSubInfo{
			Rate : int32(dv.Rate),
                       PackageName : dv.PackageName,
		 }
	     }
	     SubIds[k] = &subInfo
	}

	btV4.SubIds = SubIds
	}
	if len(camInfo.BtV4.BtClass) > 0 {
	    btClass := make(map[string]*protobuf.BtClass, len(camInfo.BtV4.BtClass))
	    for k, v := range camInfo.BtV4.BtClass {
		btClass[k] = &protobuf.BtClass{
		    Percent : v.Percent,
                   CapMargin: v.CapMargin,
                   Status : v.Status,
		}
	    }
	    btV4.BtClass = btClass
	}
	camPb.BtV4 = &btV4
   }
   if camInfo.Endcard != nil && len(*camInfo.Endcard) > 0 {
	endcard := make(map[string]*protobuf.EndCard, len(*camInfo.Endcard))
	for k, v := range *camInfo.Endcard {
	    endcardTmp := protobuf.EndCard{
		Status : *v.Status,
               Orientation: *v.Orientation,
	        EndcardProtocal : int32(*v.EndcardProtocal),
	    }

	    if len(*v.EndcardRate) > 0 {
		endcardRate := make(map[int32]int32, len(*v.EndcardRate))
		for rk,rv := range *v.EndcardRate {
		   endcardRate[int32(rk)] = int32(rv)
		}
		endcardTmp.EndcardRate = endcardRate
	    }

	    if len(*v.Urls) > 0 {
		urls := make([]*protobuf.EndCardUrls, 0)
		for _, u := range *v.Urls {
		   urls = append(urls, &protobuf.EndCardUrls{
		 	Id : *u.Id,
			Url : *u.Url,
                       Weight: *u.Weight,
                       UrlV2 : *u.UrlV2,
		   })
		}
		endcardTmp.Urls = urls
	    }
	    if len(*v.VideoTemplateUrl) > 0  {
		vUrls := make([]*protobuf.VideoTemplateUrlItem, len(*v.VideoTemplateUrl))
		for _, vu := range *v.VideoTemplateUrl {
		    vUrls = append(vUrls, &protobuf.VideoTemplateUrlItem{
			Id : *vu.ID,
			Url : *vu.URL,
			UrlZip : *vu.URLZip,
			Weight : *vu.Weight,
			PausedUrl : *vu.PausedURL,
			PausedUrlZip : *vu.PausedURLZip,
		    })
		}
		endcardTmp.VideoTemplateUrl = vUrls
	    }
	    endcard[k] = &endcardTmp

	}
	camPb.EndCard = endcard
   }
   return camPb
}

func flushALLCampaign(item *mgo.Iter) {
	updated := 0
	campaignInfo := util.CampaignInfo{}
	//redisKey := "campaign"
	// 删除
	//if err := redis.RedisDel(redisKey, 0); err != nil {
	//	fmt.Errorf("redis del %s error:%s", redisKey, err.Error())
	//}
	//i := 0
	for item.Next(&campaignInfo) {
		//fmt.Printf("%+v", campaignInfo)
		//if i == 2 {
		//    break;
		//}
		//continue;
		// 插入
		//rdb, err := json.Marshal(&campaignInfo)
		//if err != nil {
		//	fmt.Errorf("unit json marshal error: %s", err.Error())
		//}
		cpb := fillCampaignPb(campaignInfo)
		pdb, err := proto.Marshal(cpb)
		if err != nil {
			fmt.Errorf("campaign protobuf marshal error: %s", err.Error())
		}
		//crdb, err := cbrotli.Encode(rdb, cbrotli.WriterOptions{Quality: 5})
		//if err != nil {
		//	fmt.Errorf("brotli compression error: %s", err.Error())
		//}
		if err := redis.RedisSend("HSET", "campaign", strconv.FormatInt(campaignInfo.CampaignId, 10), pdb); err != nil {
			fmt.Errorf("redis write command to buffer error: %s", err.Error())
		}
		updated++
		//fmt.Println("********updated=[%d]", updated)
	}
	
	if err := redis.RedisFlush(updated); err != nil {
		fmt.Errorf("redis flush buffer error: %s", err.Error())
	}

	fmt.Println("*****************updated=%d", updated)
}
