package loginservice

import (
	"encoding/json"
	"originserver2/common/collect"
	"originserver2/common/db"
	"time"

	"github.com/duanhf2012/origin/v2/log"
	"github.com/duanhf2012/origin/v2/node"
	"github.com/duanhf2012/origin/v2/service"
	"github.com/duanhf2012/origin/v2/sysservice/httpservice"
	"github.com/duanhf2012/origin/v2/util/timer"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	node.Setup(&httpservice.HttpService{})
	node.Setup(&LoginService{})
}

type AreaGate struct {
	AreaName string
	AreaId   int
	GateInfo []GateInfoResp
}

type LoginService struct {
	service.Service
	loginModule *LoginModule
	mapTcpGate  map[int]map[string]*TcpGateInfo //map[AreaId]map[nodeId]*TcpGateInfo

	//mapTcpGateSlice map[int][]GateInfoResp //map[AreaId][]GateInfoResp

	//areaGate []AreaGate
	prepareAreaGateInfo string
	dirty               bool
	//mapArea map[int32]int
}

type TcpGateInfo struct {
	AreaName string
	AreaId   int
	Weight   int32
	Url      string
	refresh  time.Time
}

func (gate *LoginService) OnInit() error {
	gate.mapTcpGate = make(map[int]map[string]*TcpGateInfo, 100)

	areaList := gate.GetServiceCfg().([]interface{})
	for _, al := range areaList {
		mapArea := al.(map[string]interface{})
		areaId := int(mapArea["AreaId"].(float64))
		areaName := mapArea["AreaName"].(string)
		gateList := mapArea["GateList"].([]interface{})
		for _, g := range gateList {
			gateInfo := g.(map[string]interface{})
			nodeId := gateInfo["NodeId"].(string)
			addr := gateInfo["Addr"].(string)
			if gate.mapTcpGate[areaId] == nil {
				gate.mapTcpGate[areaId] = make(map[string]*TcpGateInfo, 5)
			}
			gate.mapTcpGate[areaId][nodeId] = &TcpGateInfo{Weight: -1, Url: addr, AreaId: areaId, AreaName: areaName}
		}
	}

	//获取系统httpService服务
	httpService := node.GetService("HttpService").(*httpservice.HttpService)

	//新建并设置路由对象
	httpRouter := httpservice.NewHttpHttpRouter()
	httpService.SetHttpRouter(httpRouter, gate.GetEventHandler())

	gate.loginModule = &LoginModule{}
	//gate.loginModule.funcGetGateUrl = gate.GetGateInfoUrl
	gate.loginModule.funcGetAllAreaGateUrl = gate.GetAllAreaGateUrl
	gate.AddModule(gate.loginModule)

	//性能监控
	gate.OpenProfiler()
	gate.GetProfiler().SetOverTime(time.Millisecond * 100)
	gate.GetProfiler().SetMaxOverTime(time.Second * 10)

	gate.AfterFunc(time.Second*1, gate.PrepareGateService)

	//POST方法 请求url:http://127.0.0.1:9402/login
	//返回结果为：{"msg":"hello world"}
	httpRouter.POST("/login", gate.loginModule.Login)
	httpRouter.POST("/savearea", gate.loginModule.SaveArea)

	return nil
}

func (gate *LoginService) OnStart() {
	var req db.DBControllerReq
	req.CollectName = collect.AccountCollectName
	req.Type = db.OptType_Find
	req.MaxRow = 1
	req.Condition, _ = bson.Marshal(bson.D{{"Account", "aa"}})
	req.Key = "aa"
	//ret := db.DBControllerRet{}
	//gate.GetService().GetRpcHandler().Call("DBService.RPC_DBRequest", &req, &ret)
}

// GateService->LoginService同步负载
/*func (gate *LoginService) RPC_SetTcpGateBalance(balance *rpc.NodeBalance) error {
	for _, mapNode := range gate.mapTcpGate {
		v, ok := mapNode[int(balance.NodeId)]
		if !ok {
			continue
		}

		gate.dirty = true

		v.Weight = balance.Weigh
		v.refresh = time.Now()
	}

	return nil
}*/

// 超过10秒，则判断该网关已经无效
const HealthyTimeOut = 10 * time.Second //10秒
const WillMaxArea = 100

// 定时预处理网关列表
func (gate *LoginService) PrepareGateService(timer *timer.Timer) {
	// if gate.dirty == false {
	// 	return
	// }

	areaGateSlice := make([]AreaGate, 0, WillMaxArea)
	for areaId, info := range gate.mapTcpGate {
		var areaName string
		var gateInfoSlice []GateInfoResp
		for _, tcpGateInfo := range info {
			// if time.Now().Sub(tcpGateInfo.refresh) > HealthyTimeOut { //10秒都没有同步
			// 	continue
			// }
			areaName = tcpGateInfo.AreaName
			gateInfoSlice = append(gateInfoSlice, GateInfoResp{Weight: tcpGateInfo.Weight, Url: tcpGateInfo.Url})
		}
		var areaGate AreaGate
		areaGate.AreaName = areaName
		areaGate.GateInfo = gateInfoSlice
		areaGate.AreaId = areaId
		if len(gateInfoSlice) > 0 {
			areaGateSlice = append(areaGateSlice, areaGate)
		}
	}

	prepareAreaGateInfo, err := json.Marshal(areaGateSlice)
	if err != nil {
		log.Error("param gate info fail %+v", gate.prepareAreaGateInfo)
		return
	}
	gate.prepareAreaGateInfo = string(prepareAreaGateInfo)
	gate.dirty = false
}

func (gate *LoginService) GetAllAreaGateUrl() string {
	return gate.prepareAreaGateInfo
}
