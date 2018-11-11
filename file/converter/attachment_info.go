
package model

import (
	DOMAIN_QQQQQQ "hecate/domain/post"
	MODEL_WWWWWW "hecate/model/post"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 *AttachmentInfoModelToAttachmentInfo 
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
 func AttachmentInfoModelToAttachmentInfo(attachmentInfoModel *MODEL_WWWWWW.AttachmentInfoModel) *DOMAIN_QQQQQQ.AttachmentInfo {
attachmentInfo:= new(DOMAIN_QQQQQQ.AttachmentInfo)
if attachmentInfoModel == nil {
		return attachmentInfo
	}
attachmentInfo.Id=attachmentInfoModel.Id
attachmentInfo.TargetId=attachmentInfoModel.TargetId
attachmentInfo.FilePath=attachmentInfoModel.FilePath
attachmentInfo.TypeCode=attachmentInfoModel.TypeCode
attachmentInfo.CategoryCode=attachmentInfoModel.CategoryCode
attachmentInfo.IsDeleted=attachmentInfoModel.IsDeleted
attachmentInfo.LastModifiedDate=attachmentInfoModel.LastModifiedDate
attachmentInfo.CreationDate=attachmentInfoModel.CreationDate

return attachmentInfo
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 *AttachmentInfoToAttachmentInfoModel 
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
 func AttachmentInfoToAttachmentInfoModel(attachmentInfo *DOMAIN_QQQQQQ.AttachmentInfo) *MODEL_WWWWWW.AttachmentInfoModel {
attachmentInfoModel:= new(MODEL_WWWWWW.AttachmentInfoModel)
if attachmentInfo == nil {
		return attachmentInfoModel
	}

attachmentInfoModel.Id = attachmentInfoModel.Id
	if len(attachmentInfoModel.Id) == 0 {
		guid, _ := uuid.NewV4()
		attachmentInfoModel.Id = guid.String()
	}
attachmentInfoModel.Id=attachmentInfo.Id
attachmentInfoModel.TargetId=attachmentInfo.TargetId
attachmentInfoModel.FilePath=attachmentInfo.FilePath
attachmentInfoModel.TypeCode=attachmentInfo.TypeCode
attachmentInfoModel.CategoryCode=attachmentInfo.CategoryCode
attachmentInfoModel.IsDeleted=attachmentInfo.IsDeleted
attachmentInfoModel.LastModifiedDate = attachmentInfo.LastModifiedDate

	if attachmentInfoModel.LastModifiedDate.IsZero() {
		attachmentInfoModel.LastModifiedDate = time.Now()
	}
attachmentInfoModel.CreationDate = attachmentInfo.CreationDate

	if attachmentInfoModel.CreationDate.IsZero() {
		attachmentInfoModel.CreationDate = time.Now()
	}

return attachmentInfoModel
}
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * AttachmentInfoModelListToAttachmentInfoList
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func AttachmentInfoModelListToAttachmentInfoList(attachmentInfoModelList []*MODEL_WWWWWW.AttachmentInfoModel) DOMAIN_QQQQQQ.AttachmentInfoList {
	attachmentInfoList := make(DOMAIN_QQQQQQ.AttachmentInfoList, 0, len(attachmentInfoModelList))

	for _, attachmentInfoModel:= range attachmentInfoModelList {
		attachmentInfo := AttachmentInfoModelToAttachmentInfo(attachmentInfoModel)
		attachmentInfoList = append(attachmentInfoList, attachmentInfo)
	}

	return attachmentInfoList
}
