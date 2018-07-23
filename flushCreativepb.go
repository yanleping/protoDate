package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"strconv"
	"syscall"
	"time"
	"unicode/utf8"

	//"github.com/google/brotli/go/cbrotli"
	"github.com/json-iterator/go"
	"gopkg.in/mgo.v2"
	//"github.com/vmihailenco/msgpack"
	"github.com/gogo/protobuf/proto"
	"gopkg.in/mgo.v2/bson"
	"./protobuf"
	"./redis"
	"./util"
)

const delay = 2 * time.Second

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var num = 0

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
	batchQuery := flag.Int("batchQuery", 200, "mongodb batch query limit")
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
	session, err := mgo.Dial("adn-cpmongo-slave-sg.rayjump.com:27017")  //连接mongo数据库
	if err != nil {
		panic(err)
	}
	c := session.DB("new_adn").C("creative")

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
		item = c.Find(bson.M{"status": 1}).Batch(*batchQuery).Prefetch(0.25).Iter()
		flushALL(item)
	} else {
		totalDocs, err := c.Find(bson.M{"status": 1}).Count()
		if err != nil {
			panic(err)
		}
		for i := 0; i**limit < totalDocs; i++ {
			fmt.Println(i)
			item = c.Find(bson.M{"status": 1}).Skip(*limit * i).Limit(*limit).Batch(*batchQuery).Prefetch(0.25).Iter()
			flushALL(item)
			fmt.Println("**********************************************")
			fmt.Printf("***** creative all: [%d], flush [%d]. *****\n", totalDocs, *limit*i)
			fmt.Println("**********************************************")
			time.Sleep(delay)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("###### num = ",num)
	fmt.Println(elapsed)
	fmt.Println("crtl + c exit processing")
	<-done
}

func fillCreativePb(content *util.Content) *protobuf.Creative {
	creativePb := &protobuf.Creative{}
	if len(content.Url) > 0 {
		creativePb.Url = content.Url
	}
	if len(content.VideoResolution) > 0 {
		creativePb.VideoResolution = content.VideoResolution
	}
	if len(content.Resolution) > 0 {
		creativePb.Resolution = content.Resolution
	}
	if len(content.Mime) > 0 {
		creativePb.Mime = content.Mime
	}
	if len(content.AdvCreativeId) > 0 {
		creativePb.AdvCreativeId = content.AdvCreativeId
	}
	if content.VideoLength > 0 {
		creativePb.VideoLength = content.VideoLength
	}
	if content.VideoSize > 0 {
		creativePb.VideoSize = content.VideoSize
	}
	if content.Width > 0 {
		creativePb.Width = content.Width
	}
	if content.Height > 0 {
		creativePb.Height = content.Height
	}
	if content.WatchMile > 0 {
		creativePb.WatchMile = content.WatchMile
	}
	if content.BitRate > 0 {
		creativePb.BitRate = content.BitRate
	}
	if content.CreativeId != nil {
		creativePb.CreativeId = *content.CreativeId
	}
	if content.Source != nil {
		creativePb.Source = *content.Source
	}
	if len(content.FMd5) > 0 {
		creativePb.FMd5 = content.FMd5
	}

	if content.Protocal > 0 {
		creativePb.Protocal = content.Protocal
	}

	if iv, ok := content.Value.(int); ok {
		creativePb.IValue = int32(iv)
		return creativePb
	}

	if sv, ok := content.Value.(string); ok {
		creativePb.SValue = sv
		return creativePb
	}

	if fv, ok := content.Value.(float64); ok {
		creativePb.FValue = fv
		return creativePb
	}
	return creativePb
}

func flushALL(item *mgo.Iter) {
	creative := util.CreativeInfo{}
	for item.Next(&creative) {
		redisKey := fmt.Sprintf("creative:%s", creative.ID.Hex())
		//fmt.Println(redisKey)
		// 删除
		/*if err := redis.RedisDel(redisKey, 0); err != nil {
			fmt.Errorf("redis del key[creative:%s] error:%s", redisKey, err.Error())
		}*/
		redisCreativeIds, err := redis.RedisHKeys(redisKey, 0)
		if err != nil {
			fmt.Errorf("redis key: [%s] by hkeys command error: %s", redisKey, err.Error())
		}
		var mongoCreativeIds []string
		updated := 0
		for cType, record := range creative.Content {
			x, _ := strconv.Atoi(cType)
			i := sort.Search(len(util.CreativeTypes), func(i int) bool { return util.CreativeTypes[i] >= x })
			if i < len(util.CreativeTypes) && util.CreativeTypes[i] == x {

				if rec, ok := record.(map[string]interface{}); ok {
					for field, val := range rec {
						//var mdb []byte
						/*rdb, err = json.Marshal(&val)
						if err != nil {
							fmt.Errorf("creative content json marshal error: %s", err.Error())
						}*/
						/*mdb, err = msgpack.Marshal(&val)
						if err != nil {
						    fmt.Errorf("creative content msgpack marshal error: %s", err.Error())
						}*/
						var buf bytes.Buffer
						encoder := json.NewEncoder(&buf)
						encoder.SetEscapeHTML(false)
						if err := encoder.Encode(&val); err != nil {
							fmt.Errorf("creative content json encoder error: %s", err.Error())
						}
						var content *util.Content
						if err := json.Unmarshal([]byte(buf.String()), &content); err != nil {
							fmt.Errorf("creative content json unmarshal error: %s", err.Error())
						}
						cpb := fillCreativePb(content)
						if x == util.APP_DESC {
							if utf8.RuneCountInString(cpb.SValue) > 90 {
								detail := []rune(cpb.SValue)
								cpb.SValue = string(detail[0:90]) + "..."
							}
						}
						pdb, err := proto.Marshal(cpb)
						//fmt.Print(pdb)
						if err != nil {
							fmt.Errorf("creative protobuf marshal error: %s", err.Error())
						}

						if err := redis.RedisSend("HSET", redisKey, field, pdb); err != nil {
							fmt.Errorf("redis write command to buffer error: %s", err.Error())
						}
						mongoCreativeIds = append(mongoCreativeIds, field)
						//fmt.Println(mongoCreativeIds)
						// for redis flush count
						updated++
						num++
					}
				} else {
					fmt.Printf("creative content not a map[string]interface{}: %+v", record)
				}
			}
		}
		//fmt.Println("********updated=",updated)
		// 删除素材表已经删除的素材类型配置
		for _, v := range redisCreativeIds.([]interface{}) {
			field := fmt.Sprintf("%s", v)
			found := util.StringInSlice(field, mongoCreativeIds)
			//fmt.Printf("redis key: %s, field: %s in mongodb:%s is %t\n", redisKey, field, mongoCreativeIds, found)
			if !found {
				if err := redis.RedisDelBatch(redisKey, field); err != nil {
					fmt.Errorf("redis key: %s on hdel field: %s command error: %s", redisKey, field, err.Error())
				}
				// for redis flush count
				updated++
			}
		}
		if err := redis.RedisFlush(updated); err != nil {
			fmt.Errorf("redis flush buffer error: %s", err.Error())
		}
	}
}
