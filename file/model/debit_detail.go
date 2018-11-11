
     package model

    import (
       "time"
    )

     import (
	   "github.com/jinzhu/gorm"
       "github.com/hicsgo/ging"
    )

     import (
        "hecate/domain/common"
	    "hecate/model"
    )
    
     /* ================================================================================
      * 模型 
      * author:jcheng 
      * ================================================================================ */

type DebitDetailModel struct{
model.BaseModel
Id string `gorm:"column:id"`//  
DebitId string `gorm:"column:debit_id"`//请款单id  
PurchasingId string `gorm:"column:purchasing_id"`//采购单id  
PurchasingDetailId string `gorm:"column:purchasing_detail_id"`//采购单明细id  
BuyingPrice float64 `gorm:"column:buying_price"`//购车单价  
SalePrice float64 `gorm:"column:sale_price"`//售车单价  
Profit float64 `gorm:"column:profit"`//单台利润  
Quantity int64 `gorm:"column:quantity"`//请款台数  
Remark string `gorm:"column:remark"`//备注  
CreationDate time.Time `gorm:"column:creation_date"`//  
LastModifiedDate time.Time `gorm:"column:last_modified_date"`//  
}
func (debitDetailModel *DebitDetailModel) TableName () string {
	 	return "debit_detail"
}

func (debitDetailModel *DebitDetailModel) GetProjectName() string {
	return common.TAPAI_PROJECT_NAME
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断数据是否存在（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (debitDetailModel *DebitDetailModel) IsExists(
	query interface{},
	args ...interface{}) (bool, error) {

	if count, err := debitDetailModel.GetCount(query, args...); err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return true, err
	} else {
		if count > 0 {
			return true, nil
		}
	}

	return false, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据记录数（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (debitDetailModel *DebitDetailModel) GetCount(query interface{}, args ...interface{}) (int64, error) {
	debitDetailModel.ReadModelDbMap(debitDetailModel)
	defer debitDetailModel.DbMap.Close()

	paging := ging.Paging{
		PagingIndex: 1,
		PagingSize:  1,
	}

	if err := debitDetailModel.DbMap.Model(debitDetailModel).Where(
		query, args...).Count(&paging.TotalRecord).Error; err != nil {
		return 0, err
	}

	return paging.TotalRecord, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（单个简单条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (debitDetailModel *DebitDetailModel) Select(fieldName string, fieldValue interface{}) error {
	debitDetailModel.ReadModelDbMap(debitDetailModel)
	defer debitDetailModel.DbMap.Close()

	query := map[string]interface{}{}
	query[fieldName] = fieldValue

	if err := debitDetailModel.DbMap.Find(debitDetailModel, query).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NotFoundError
		}
		return err
	}
	return nil
}
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (debitDetailModel *DebitDetailModel) SelectQuery(query interface{}, args ...interface{}) error {
	debitDetailModel.ReadModelDbMap(debitDetailModel)
	defer debitDetailModel.DbMap.Close()

	if err := debitDetailModel.DbMap.Where(
		query, args...).Find(debitDetailModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NotFoundError
		}
		return err
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (debitDetailModel *DebitDetailModel) SelectOrderQuery(query interface{}, sortorder string, args ...interface{}) error {
	debitDetailModel.ReadModelDbMap(debitDetailModel)
	defer debitDetailModel.DbMap.Close()

	if err := debitDetailModel.DbMap.Where(
		query, args...).Order(sortorder).Limit(1).Find(debitDetailModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NotFoundError
		}
		return err
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（主键标识简单条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (debitDetailModel *DebitDetailModel) SelectById(id interface{}) error {
	
	debitDetailModel.ReadModelDbMap(debitDetailModel)
	defer debitDetailModel.DbMap.Close()

	query := map[string]interface{}{
		"id": id,
	}

	if err := debitDetailModel.DbMap.Find(debitDetailModel, query).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NotFoundError
		}
		return err
	}

	return nil
}

/* ================================================================================
 * query: []interface{} || map[string]interface{} || string
 * args: if string: interface{}...
 * ================================================================================ */
func (debitDetailModel *DebitDetailModel) SelectAll(paging *ging.Paging, query interface{}, args ...interface{}) ([]*DebitDetailModel, error) {
	debitDetailModel.ReadModelDbMap(debitDetailModel)
	defer debitDetailModel.DbMap.Close()

	var debitDetailModelList = make([]*DebitDetailModel, 0)
	var err error = nil

	if paging != nil {
		isTotalRecord := true
		if paging.IsTotalOnce {
			if paging.PagingIndex > 1 {
				isTotalRecord = false
			}
		}

		if isTotalRecord && paging.PagingSize > 0 {
			if len(paging.Group) == 0 {
				err = debitDetailModel.DbMap.Model(debitDetailModel).
					Where(query, args...).
					Count(&paging.TotalRecord).
					Order(paging.Sortorder).
					Offset(paging.Offset()).
					Limit(paging.PagingSize).
					Find(&debitDetailModelList).Error
			} else {
				err = debitDetailModel.DbMap.Model(debitDetailModel).
					Where(query, args...).
					Group(paging.Group).
					Count(&paging.TotalRecord).
					Order(paging.Sortorder).
					Offset(paging.Offset()).
					Limit(paging.PagingSize).
					Find(&debitDetailModelList).Error
			}

			paging.SetTotalRecord(paging.TotalRecord)
		} else {
			if len(paging.Group) == 0 {
				err = debitDetailModel.DbMap.Model(debitDetailModel).
					Where(query, args...).
					Order(paging.Sortorder).
					Find(&debitDetailModelList).Error
			} else {
				err = debitDetailModel.DbMap.Model(debitDetailModel).
					Where(query, args...).
					Group(paging.Group).
					Order(paging.Sortorder).
					Find(&debitDetailModelList).Error
			}
		}
	} else {
		err = debitDetailModel.DbMap.Where(query, args...).Find(&debitDetailModelList).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = common.NotFoundError
		}
	}

	return debitDetailModelList, err
}
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 插入数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (debitDetailModel *DebitDetailModel) Insert() error {
    if len(debitDetailModel.Id)==0{
			return common.ArgumentError
		}
	if debitDetailModel.DbMap == nil {
		debitDetailModel.WriteModelDbMap(debitDetailModel)
		defer debitDetailModel.DbMap.Close()
	}

	if err := debitDetailModel.DbMap.Create(debitDetailModel).Error; err != nil {
		return err
	}
	return nil
}
/* ================================================================================
 * data type:
 * Model{"fieldName":"value"...}
 * map[string]interface{}
 * key1,value1,key2,value2
 * ================================================================================ */
func (debitDetailModel *DebitDetailModel) Update(data ...interface{}) (int64, error) {
	
if len(debitDetailModel.Id) == 0 || len(data) == 0 {
		return 0, common.ArgumentError
	}


	if debitDetailModel.DbMap == nil {
		debitDetailModel.WriteModelDbMap(debitDetailModel)
		defer debitDetailModel.DbMap.Close()
	}

	dbContext := debitDetailModel.DbMap.Model(debitDetailModel).UpdateColumns(data)
	rowsAffected, err := dbContext.RowsAffected, dbContext.Error

	return rowsAffected, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 删除数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (debitDetailModel *DebitDetailModel) Delete() (int64, error) {
     if len(debitDetailModel.Id)==0{
			return 0,common.ArgumentError
		}	

	if debitDetailModel.DbMap == nil {
		debitDetailModel.WriteModelDbMap(debitDetailModel)
		defer debitDetailModel.DbMap.Close()
	}

	dbContext := debitDetailModel.DbMap.Delete(debitDetailModel)
	rowsAffected, err := dbContext.RowsAffected, dbContext.Error

	return rowsAffected, err
}