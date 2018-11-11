
package service

import (
	DOMAIN_QQQQQQ "hecate/domain/post"
	MODEL_WWWWWW "hecate/model/post"
	"time"
)
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 服务接口
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type IAttachmentInfoService interface{
    GetAttachmentInfoById(id int64) (*DOMAIN_QQQQQQ.AttachmentInfo, error)
	GetAttachmentInfoList(attachmentInfoCondition *DOMAIN_QQQQQQ.AttachmentInfoCondition) (DOMAIN_QQQQQQ.AttachmentInfoList, error)
	CreateAttachmentInfo(attachmentInfo *DOMAIN_QQQQQQ.AttachmentInfo) error
	UpdateAttachmentInfoById(id int64, data map[string]interface{}) error
	DeleteAttachmentInfoById(id int64) error	
}
