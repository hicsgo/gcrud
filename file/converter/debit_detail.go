
package model

import (
	DOMAIN_QQQQQQ "hecate/domain/post"
	MODEL_WWWWWW "hecate/model/post"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 *DebitDetailModelToDebitDetail 
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
 func DebitDetailModelToDebitDetail(debitDetailModel *MODEL_WWWWWW.DebitDetailModel) *DOMAIN_QQQQQQ.DebitDetail {
debitDetail:= new(DOMAIN_QQQQQQ.DebitDetail)
if debitDetailModel == nil {
		return debitDetail
	}
debitDetail.Id=debitDetailModel.Id
debitDetail.DebitId=debitDetailModel.DebitId
debitDetail.PurchasingId=debitDetailModel.PurchasingId
debitDetail.PurchasingDetailId=debitDetailModel.PurchasingDetailId
debitDetail.BuyingPrice=debitDetailModel.BuyingPrice
debitDetail.SalePrice=debitDetailModel.SalePrice
debitDetail.Profit=debitDetailModel.Profit
debitDetail.Quantity=debitDetailModel.Quantity
debitDetail.Remark=debitDetailModel.Remark
debitDetail.CreationDate=debitDetailModel.CreationDate
debitDetail.LastModifiedDate=debitDetailModel.LastModifiedDate

return debitDetail
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 *DebitDetailToDebitDetailModel 
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
 func DebitDetailToDebitDetailModel(debitDetail *DOMAIN_QQQQQQ.DebitDetail) *MODEL_WWWWWW.DebitDetailModel {
debitDetailModel:= new(MODEL_WWWWWW.DebitDetailModel)
if debitDetail == nil {
		return debitDetailModel
	}

debitDetailModel.Id = debitDetailModel.Id
	if len(debitDetailModel.Id) == 0 {
		guid, _ := uuid.NewV4()
		debitDetailModel.Id = guid.String()
	}
debitDetailModel.Id=debitDetail.Id
debitDetailModel.DebitId=debitDetail.DebitId
debitDetailModel.PurchasingId=debitDetail.PurchasingId
debitDetailModel.PurchasingDetailId=debitDetail.PurchasingDetailId
debitDetailModel.BuyingPrice=debitDetail.BuyingPrice
debitDetailModel.SalePrice=debitDetail.SalePrice
debitDetailModel.Profit=debitDetail.Profit
debitDetailModel.Quantity=debitDetail.Quantity
debitDetailModel.Remark=debitDetail.Remark
debitDetailModel.CreationDate = debitDetail.CreationDate

	if debitDetailModel.CreationDate.IsZero() {
		debitDetailModel.CreationDate = time.Now()
	}
debitDetailModel.LastModifiedDate = debitDetail.LastModifiedDate

	if debitDetailModel.LastModifiedDate.IsZero() {
		debitDetailModel.LastModifiedDate = time.Now()
	}

return debitDetailModel
}
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * DebitDetailModelListToDebitDetailList
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DebitDetailModelListToDebitDetailList(debitDetailModelList []*MODEL_WWWWWW.DebitDetailModel) DOMAIN_QQQQQQ.DebitDetailList {
	debitDetailList := make(DOMAIN_QQQQQQ.DebitDetailList, 0, len(debitDetailModelList))

	for _, debitDetailModel:= range debitDetailModelList {
		debitDetail := DebitDetailModelToDebitDetail(debitDetailModel)
		debitDetailList = append(debitDetailList, debitDetail)
	}

	return debitDetailList
}
