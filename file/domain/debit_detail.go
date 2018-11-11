
	package post

	import (
		"github.com/hicsgo/ging"
	)

/* ================================================================================
 * 模型映射数据域结构
 * author: jcheng
 * ================================================================================ */
type DebitDetailList []*DebitDetail
type DebitDetail struct {
Id string `form:"id" json:"id"`//
DebitId string `form:"debit_id" json:"debit_id"`//请款单id
PurchasingId string `form:"purchasing_id" json:"purchasing_id"`//采购单id
PurchasingDetailId string `form:"purchasing_detail_id" json:"purchasing_detail_id"`//采购单明细id
BuyingPrice float64 `form:"buying_price" json:"buying_price"`//购车单价
SalePrice float64 `form:"sale_price" json:"sale_price"`//售车单价
Profit float64 `form:"profit" json:"profit"`//单台利润
Quantity int64 `form:"quantity" json:"quantity"`//请款台数
Remark string `form:"remark" json:"remark"`//备注
CreationDate time.Time `form:"creation_date" json:"creation_date"`//
LastModifiedDate time.Time `form:"last_modified_date" json:"last_modified_date"`//

}
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 模型映射查询数据域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type DebitDetailCondition struct {
Id  string `form:"id" json:"id"`
DebitId  string `form:"debit_id" json:"debit_id"`
PurchasingId  string `form:"purchasing_id" json:"purchasing_id"`
PurchasingDetailId  string `form:"purchasing_detail_id" json:"purchasing_detail_id"`
MinCreationDate  time.Time `form:"min_creation_date" json:"min_creation_date"`
MaxCreationDate  time.Time `form:"max_creation_date" json:"max_creation_date"`
MinLastModifiedDate  time.Time `form:"min_last_modified_date" json:"min_last_modified_date"`
MaxLastModifiedDate  time.Time `form:"max_last_modified_date" json:"max_last_modified_date"`
Paging  *ging.Paging  `form:"paging" json:"paging"`
}
