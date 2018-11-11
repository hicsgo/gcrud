
package service

import (
	DOMAIN_QQQQQQ "hecate/domain/post"
	MODEL_WWWWWW "hecate/model/post"
	"time"
)
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据Id标识获取数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) GetDebitDetailById(id int64) (*DOMAIN_QQQQQQ.DebitDetail, error) {
	if id == 0 {
		return nil, common.ArgumentError
	}

	debitDetailModel := new(MODEL_WWWWWW.DebitDetailModel)
	if err := debitDetailModel.SelectById(id); err != nil {
		return nil, err
	}

	return DebitDetailModelToDebitDetail(debitDetailModel), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) GetDebitDetailList(debitDetailCondition *DOMAIN_QQQQQQ.DebitDetailCondition) (DOMAIN_QQQQQQ.DebitDetailList, error) {
	if debitDetailCondition == nil {
		return nil, common.ArgumentError
	}

	fields := make([]string, 0)
	args := make([]interface{}, 0)

	


if len(debitDetailCondition.Id) > 0 {
		fields = append(fields, "id = ?")
		args = append(args, debitDetailModelCondition.Id)
	}

if len(debitDetailCondition.DebitId) > 0 {
		fields = append(fields, "debit_id = ?")
		args = append(args, debitDetailModelCondition.DebitId)
	}

if len(debitDetailCondition.PurchasingId) > 0 {
		fields = append(fields, "purchasing_id = ?")
		args = append(args, debitDetailModelCondition.PurchasingId)
	}

if len(debitDetailCondition.PurchasingDetailId) > 0 {
		fields = append(fields, "purchasing_detail_id = ?")
		args = append(args, debitDetailModelCondition.PurchasingDetailId)
	}

if !debitDetailCondition.MinCreationDate.IsZero()  {
		fields = append(fields, "creation_date >= ?")
		args = append(args, debitDetailModelCondition.MinCreationDate)
	}
if !debitDetailCondition.MaxCreationDate.IsZero()  {
		fields = append(fields, "creation_date < ?")
		args = append(args, debitDetailModelCondition.MaxCreationDate)
	}

if !debitDetailCondition.MinLastModifiedDate.IsZero()  {
		fields = append(fields, "last_modified_date >= ?")
		args = append(args, debitDetailModelCondition.MinLastModifiedDate)
	}
if !debitDetailCondition.MaxLastModifiedDate.IsZero()  {
		fields = append(fields, "last_modified_date < ?")
		args = append(args, debitDetailModelCondition.MaxLastModifiedDate)
	}


	query := glib.StringSliceToString(fields, " AND ")

	debitDetailModel := new(MODEL_WWWWWW.DebitDetailModel)
	debitDetailModelList, err := debitDetailModel.SelectAll(debitDetailCondition.Paging, query, args...)
	if err != nil {
		return nil, err
	}

	return DebitDetailModelListToDebitDetailList(debitDetailModelList), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) CreateDebitDetail(debitDetail *DOMAIN_QQQQQQ.DebitDetail) error {
	if DebitDetail == nil ||DebitDetail.Id==0 { 
		return common.ArgumentError
	}

	//插入数据
	debitDetailModel := DebitDetailToDebitDetailModel(DebitDetail)
	if err := debitDetailModel.Insert(); err != nil {
		return err
	}

	DebitDetail.Id = debitDetailModel.Id

	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 更新数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) UpdateDebitDetailById(id int64, data map[string]interface{}) error {
	if id == 0 || len(data) == 0 {
		return common.ArgumentError
	}

	debitDetailModel := new(MODEL_WWWWWW.DebitDetailModel)
	debitDetailModel.Id = id
	_, err := debitDetailModel.Update(data)

	return err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 删除数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) DeleteDebitDetailById(id int64) error {
	if id == 0 {
		return common.ArgumentError
	}

	debitDetailModel := new(MODEL_WWWWWW.DebitDetailModel)
	debitDetailModel.Id = id
	_, err := debitDetailModel.Delete()

	return err
}
