
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

type AttachmentInfoModel struct{
model.BaseModel
Id string `gorm:"column:id"`//  
TargetId string `gorm:"column:target_id"`//目标id  
FilePath string `gorm:"column:file_path"`//  
TypeCode string `gorm:"column:type_code"`//类型编码（payment:付款凭证 | buy_contract:购车合同 | sale_contract:售车合同 | account:收款帐号）  
CategoryCode string `gorm:"column:category_code"`//目标类别编码（purchasing:采购单 | debit:请款单 | returned:回款单）  
IsDeleted bool `gorm:"column:is_deleted"`//是否删除  
LastModifiedDate time.Time `gorm:"column:last_modified_date"`//  
CreationDate time.Time `gorm:"column:creation_date"`//  
}
func (attachmentInfoModel *AttachmentInfoModel) TableName () string {
	 	return "attachment_info"
}

func (attachmentInfoModel *AttachmentInfoModel) GetProjectName() string {
	return common.TAPAI_PROJECT_NAME
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断数据是否存在（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (attachmentInfoModel *AttachmentInfoModel) IsExists(
	query interface{},
	args ...interface{}) (bool, error) {

	if count, err := attachmentInfoModel.GetCount(query, args...); err != nil {
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
func (attachmentInfoModel *AttachmentInfoModel) GetCount(query interface{}, args ...interface{}) (int64, error) {
	attachmentInfoModel.ReadModelDbMap(attachmentInfoModel)
	defer attachmentInfoModel.DbMap.Close()

	paging := ging.Paging{
		PagingIndex: 1,
		PagingSize:  1,
	}

	if err := attachmentInfoModel.DbMap.Model(attachmentInfoModel).Where(
		query, args...).Count(&paging.TotalRecord).Error; err != nil {
		return 0, err
	}

	return paging.TotalRecord, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（单个简单条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (attachmentInfoModel *AttachmentInfoModel) Select(fieldName string, fieldValue interface{}) error {
	attachmentInfoModel.ReadModelDbMap(attachmentInfoModel)
	defer attachmentInfoModel.DbMap.Close()

	query := map[string]interface{}{}
	query[fieldName] = fieldValue

	if err := attachmentInfoModel.DbMap.Find(attachmentInfoModel, query).Error; err != nil {
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
func (attachmentInfoModel *AttachmentInfoModel) SelectQuery(query interface{}, args ...interface{}) error {
	attachmentInfoModel.ReadModelDbMap(attachmentInfoModel)
	defer attachmentInfoModel.DbMap.Close()

	if err := attachmentInfoModel.DbMap.Where(
		query, args...).Find(attachmentInfoModel).Error; err != nil {
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
func (attachmentInfoModel *AttachmentInfoModel) SelectOrderQuery(query interface{}, sortorder string, args ...interface{}) error {
	attachmentInfoModel.ReadModelDbMap(attachmentInfoModel)
	defer attachmentInfoModel.DbMap.Close()

	if err := attachmentInfoModel.DbMap.Where(
		query, args...).Order(sortorder).Limit(1).Find(attachmentInfoModel).Error; err != nil {
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
func (attachmentInfoModel *AttachmentInfoModel) SelectById(id interface{}) error {
	
	attachmentInfoModel.ReadModelDbMap(attachmentInfoModel)
	defer attachmentInfoModel.DbMap.Close()

	query := map[string]interface{}{
		"id": id,
	}

	if err := attachmentInfoModel.DbMap.Find(attachmentInfoModel, query).Error; err != nil {
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
func (attachmentInfoModel *AttachmentInfoModel) SelectAll(paging *ging.Paging, query interface{}, args ...interface{}) ([]*AttachmentInfoModel, error) {
	attachmentInfoModel.ReadModelDbMap(attachmentInfoModel)
	defer attachmentInfoModel.DbMap.Close()

	var attachmentInfoModelList = make([]*AttachmentInfoModel, 0)
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
				err = attachmentInfoModel.DbMap.Model(attachmentInfoModel).
					Where(query, args...).
					Count(&paging.TotalRecord).
					Order(paging.Sortorder).
					Offset(paging.Offset()).
					Limit(paging.PagingSize).
					Find(&attachmentInfoModelList).Error
			} else {
				err = attachmentInfoModel.DbMap.Model(attachmentInfoModel).
					Where(query, args...).
					Group(paging.Group).
					Count(&paging.TotalRecord).
					Order(paging.Sortorder).
					Offset(paging.Offset()).
					Limit(paging.PagingSize).
					Find(&attachmentInfoModelList).Error
			}

			paging.SetTotalRecord(paging.TotalRecord)
		} else {
			if len(paging.Group) == 0 {
				err = attachmentInfoModel.DbMap.Model(attachmentInfoModel).
					Where(query, args...).
					Order(paging.Sortorder).
					Find(&attachmentInfoModelList).Error
			} else {
				err = attachmentInfoModel.DbMap.Model(attachmentInfoModel).
					Where(query, args...).
					Group(paging.Group).
					Order(paging.Sortorder).
					Find(&attachmentInfoModelList).Error
			}
		}
	} else {
		err = attachmentInfoModel.DbMap.Where(query, args...).Find(&attachmentInfoModelList).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = common.NotFoundError
		}
	}

	return attachmentInfoModelList, err
}
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 插入数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (attachmentInfoModel *AttachmentInfoModel) Insert() error {
    if len(attachmentInfoModel.Id)==0{
			return common.ArgumentError
		}
	if attachmentInfoModel.DbMap == nil {
		attachmentInfoModel.WriteModelDbMap(attachmentInfoModel)
		defer attachmentInfoModel.DbMap.Close()
	}

	if err := attachmentInfoModel.DbMap.Create(attachmentInfoModel).Error; err != nil {
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
func (attachmentInfoModel *AttachmentInfoModel) Update(data ...interface{}) (int64, error) {
	
if len(attachmentInfoModel.Id) == 0 || len(data) == 0 {
		return 0, common.ArgumentError
	}


	if attachmentInfoModel.DbMap == nil {
		attachmentInfoModel.WriteModelDbMap(attachmentInfoModel)
		defer attachmentInfoModel.DbMap.Close()
	}

	dbContext := attachmentInfoModel.DbMap.Model(attachmentInfoModel).UpdateColumns(data)
	rowsAffected, err := dbContext.RowsAffected, dbContext.Error

	return rowsAffected, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 删除数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (attachmentInfoModel *AttachmentInfoModel) Delete() (int64, error) {
     if len(attachmentInfoModel.Id)==0{
			return 0,common.ArgumentError
		}	

	if attachmentInfoModel.DbMap == nil {
		attachmentInfoModel.WriteModelDbMap(attachmentInfoModel)
		defer attachmentInfoModel.DbMap.Close()
	}

	dbContext := attachmentInfoModel.DbMap.Delete(attachmentInfoModel)
	rowsAffected, err := dbContext.RowsAffected, dbContext.Error

	return rowsAffected, err
}