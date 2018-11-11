
	package post

	import (
		"github.com/hicsgo/ging"
	)

/* ================================================================================
 * 模型映射数据域结构
 * author: jcheng
 * ================================================================================ */
type AttachmentInfoList []*AttachmentInfo
type AttachmentInfo struct {
Id string `form:"id" json:"id"`//
TargetId string `form:"target_id" json:"target_id"`//目标id
FilePath string `form:"file_path" json:"file_path"`//
TypeCode string `form:"type_code" json:"type_code"`//类型编码（payment:付款凭证 | buy_contract:购车合同 | sale_contract:售车合同 | account:收款帐号）
CategoryCode string `form:"category_code" json:"category_code"`//目标类别编码（purchasing:采购单 | debit:请款单 | returned:回款单）
IsDeleted bool `form:"is_deleted" json:"is_deleted"`//是否删除
LastModifiedDate time.Time `form:"last_modified_date" json:"last_modified_date"`//
CreationDate time.Time `form:"creation_date" json:"creation_date"`//

}
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 模型映射查询数据域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type AttachmentInfoCondition struct {
Id  string `form:"id" json:"id"`
TargetId  string `form:"target_id" json:"target_id"`
IsDeleted  *bool `form:"is_deleted" json:"is_deleted"`
MinLastModifiedDate  time.Time `form:"min_last_modified_date" json:"min_last_modified_date"`
MaxLastModifiedDate  time.Time `form:"max_last_modified_date" json:"max_last_modified_date"`
MinCreationDate  time.Time `form:"min_creation_date" json:"min_creation_date"`
MaxCreationDate  time.Time `form:"max_creation_date" json:"max_creation_date"`
Paging  *ging.Paging  `form:"paging" json:"paging"`
}
