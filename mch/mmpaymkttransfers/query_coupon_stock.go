package mmpaymkttransfers

import (
	"gopkg.in/skOak/wechat.v2/mch/core"
)

// 查询代金券批次信息.
func QueryCouponStock(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL(clt.Sandbox())+"/mmpaymkttransfers/query_coupon_stock", req)
}
