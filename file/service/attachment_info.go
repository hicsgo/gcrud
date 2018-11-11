
package service

import (
	DOMAIN_QQQQQQ "hecate/domain/post"
	MODEL_WWWWWW "hecate/model/post"
	"time"
)
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据Id标识获取数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) GetAttachmentInfoById(id int64) (*DOMAIN_QQQQQQ.AttachmentInfo, error) {
	if id == 0 {
		return nil, common.ArgumentError
	}

	attachmentInfoModel := new(MODEL_WWWWWW.AttachmentInfoModel)
	if err := attachmentInfoModel.SelectById(id); err != nil {
		return nil, err
	}

	return AttachmentInfoModelToAttachmentInfo(attachmentInfoModel), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) GetAttachmentInfoList(attachmentInfoCondition *DOMAIN_QQQQQQ.AttachmentInfoCondition) (DOMAIN_QQQQQQ.AttachmentInfoList, error) {
	if attachmentInfoCondition == nil {
		return nil, common.ArgumentError
	}

	fields := make([]string, 0)
	args := make([]interface{}, 0)

	


if len(attachmentInfoCondition.Id) > 0 {
		fields = append(fields, "id = ?")
		args = append(args, attachmentInfoModelCondition.Id)
	}

if len(attachmentInfoCondition.TargetId) > 0 {
		fields = append(fields, "target_id = ?")
		args = append(args, attachmentInfoModelCondition.TargetId)
	}

if attachmentInfoCondition.IsDeleted!=nil {
        isTrue:=0
         if *attachmentInfoCondition.IsDeleted {
         			isTrue = 1
         		}
		fields = append(fields, "is_deleted = ?")
		args = append(args,isTrue)
	}

if !attachmentInfoCondition.MinLastModifiedDate.IsZero()  {
		fields = append(fields, "last_modified_date >= ?")
		args = append(args, attachmentInfoModelCondition.MinLastModifiedDate)
	}
if !attachmentInfoCondition.MaxLastModifiedDate.IsZero()  {
		fields = append(fields, "last_modified_date < ?")
		args = append(args, attachmentInfoModelCondition.MaxLastModifiedDate)
	}

if !attachmentInfoCondition.MinCreationDate.IsZero()  {
		fields = append(fields, "creation_date >= ?")
		args = append(args, attachmentInfoModelCondition.MinCreationDate)
	}
if !attachmentInfoCondition.MaxCreationDate.IsZero()  {
		fields = append(fields, "creation_date < ?")
		args = append(args, attachmentInfoModelCondition.MaxCreationDate)
	}


	query := glib.StringSliceToString(fields, " AND ")

	attachmentInfoModel := new(MODEL_WWWWWW.AttachmentInfoModel)
	attachmentInfoModelList, err := attachmentInfoModel.SelectAll(attachmentInfoCondition.Paging, query, args...)
	if err != nil {
		return nil, err
	}

	return AttachmentInfoModelListToAttachmentInfoList(attachmentInfoModelList), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) CreateAttachmentInfo(attachmentInfo *DOMAIN_QQQQQQ.AttachmentInfo) error {
	if AttachmentInfo == nil ||AttachmentInfo.Id==0 { 
		return common.ArgumentError
	}

	//插入数据
	attachmentInfoModel := AttachmentInfoToAttachmentInfoModel(AttachmentInfo)
	if err := attachmentInfoModel.Insert(); err != nil {
		return err
	}

	AttachmentInfo.Id = attachmentInfoModel.Id

	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 更新数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) UpdateAttachmentInfoById(id int64, data map[string]interface{}) error {
	if id == 0 || len(data) == 0 {
		return common.ArgumentError
	}

	attachmentInfoModel := new(MODEL_WWWWWW.AttachmentInfoModel)
	attachmentInfoModel.Id = id
	_, err := attachmentInfoModel.Update(data)

	return err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 删除数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Service_EEEEEE) DeleteAttachmentInfoById(id int64) error {
	if id == 0 {
		return common.ArgumentError
	}

	attachmentInfoModel := new(MODEL_WWWWWW.AttachmentInfoModel)
	attachmentInfoModel.Id = id
	_, err := attachmentInfoModel.Delete()

	return err
}
