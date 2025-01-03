package collect

import (
	"originserver2/common/db"

	"go.mongodb.org/mongo-driver/bson"
)

type EServerMark = int32

const (
	ESMNormal    EServerMark = 0 //普通
	ESMNew       EServerMark = 1 //新服
	ESMRecommend EServerMark = 2 //推荐
)

type EServerStatus = int32

const (
	ESSRight    EServerStatus = 0 //正常
	ESSMaintain EServerStatus = 1 //维护
)

type CRealAreaInfo struct {
	BaseMultiCollection `bson:"-"`
	RealAreaId          int `bson:"_id"` //真实区服ID

	GateList []string `bson:"GateList"` //地址列表 ip:端口
}

var realAreaInfo = "RealAreaInfo"
var emptyRealAreaInfo CRealAreaInfo

func (ra *CRealAreaInfo) GetCollName() string {
	return realAreaInfo
}

func (ra *CRealAreaInfo) Clean() {
	*ra = emptyRealAreaInfo
}

func (ra *CRealAreaInfo) GetId() interface{} {
	return ra.RealAreaId
}

func (ra *CRealAreaInfo) GetCollectionType() MultiCollectionType {
	return MTCRealArea
}

func (ra *CRealAreaInfo) MakeRow() IMultiCollection {
	return &CRealAreaInfo{}
}

func (ra *CRealAreaInfo) GetCondition(value interface{}) bson.D {
	return bson.D{}
}

func (ra *CRealAreaInfo) GetSort() []*db.Sort {
	return nil
}

func (ra *CRealAreaInfo) OnLoadSucc(notFound bool, userID uint64) {

}

func (ra *CRealAreaInfo) GetUpdateData() bson.M {
	return bson.M{}
}
