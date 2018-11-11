
package service

import (
	DOMAIN_QQQQQQ "hecate/domain/post"
	MODEL_WWWWWW "hecate/model/post"
	"time"
)
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 服务接口
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type IDebitDetailService interface{
    GetDebitDetailById(id int64) (*DOMAIN_QQQQQQ.DebitDetail, error)
	GetDebitDetailList(debitDetailCondition *DOMAIN_QQQQQQ.DebitDetailCondition) (DOMAIN_QQQQQQ.DebitDetailList, error)
	CreateDebitDetail(debitDetail *DOMAIN_QQQQQQ.DebitDetail) error
	UpdateDebitDetailById(id int64, data map[string]interface{}) error
	DeleteDebitDetailById(id int64) error	
}
