package build

import (
	"gcrud/model"
	"fmt"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 生成模型数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func BuildModel(parent *model.BussinessParent) *model.WriteModel {
	strs := make([]string, 0)

	//导入模块
	imports := `
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
    `
	strs = append(strs, imports)

	//头部说明
	title := `
     /* ================================================================================
      * ` + parent.TitleName + ` 
      * author:jcheng 
      * ================================================================================ */
`
	strs = append(strs, title)

	//结构体说明
	structs := `
type ` + parent.ModelStuctName + ` struct{
model.BaseModel
`
	strs = append(strs, structs)

	//生成内容体
	for _, column := range parent.Childs {
		//if column.ColumnName == "id" {
		//	//id在model.BaseModel中
		//	continue
		//}
		strs = append(strs, `` + column.ModelName + ` ` + column.ModelType + ` ` + fmt.Sprintf("`gorm:\"column:%v\"`//%v", column.ColumnName, column.ModelTag)+ `  
`)
	}
	strs = append(strs, "}")

	//gorm表指定名称
	tableName := `
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) TableName () string {
	 	return "` + parent.TableName + `"
}

func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) GetProjectName() string {
	return common.TAPAI_PROJECT_NAME
}
`
	strs = append(strs, tableName)

	//判断是否存在
	isExists := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断数据是否存在（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) IsExists(
	query interface{},
	args ...interface{}) (bool, error) {

	if count, err := ` + parent.FirstLowerModel + `.GetCount(query, args...); err != nil {
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
`
	strs = append(strs, isExists)

	//获取数量
	count := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据记录数（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) GetCount(query interface{}, args ...interface{}) (int64, error) {
	` + parent.FirstLowerModel + `.ReadModelDbMap(` + parent.FirstLowerModel + `)
	defer ` + parent.FirstLowerModel + `.DbMap.Close()

	paging := ging.Paging{
		PagingIndex: 1,
		PagingSize:  1,
	}

	if err := ` + parent.FirstLowerModel + `.DbMap.Model(` + parent.FirstLowerModel + `).Where(
		query, args...).Count(&paging.TotalRecord).Error; err != nil {
		return 0, err
	}

	return paging.TotalRecord, nil
}
`
	strs = append(strs, count)
	single := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（单个简单条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) Select(fieldName string, fieldValue interface{}) error {
	` + parent.FirstLowerModel + `.ReadModelDbMap(` + parent.FirstLowerModel + `)
	defer ` + parent.FirstLowerModel + `.DbMap.Close()

	query := map[string]interface{}{}
	query[fieldName] = fieldValue

	if err := ` + parent.FirstLowerModel + `.DbMap.Find(` + parent.FirstLowerModel + `, query).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NotFoundError
		}
		return err
	}
	return nil
}`
	strs = append(strs, single)

	selectSingle := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) SelectQuery(query interface{}, args ...interface{}) error {
	` + parent.FirstLowerModel + `.ReadModelDbMap(` + parent.FirstLowerModel + `)
	defer ` + parent.FirstLowerModel + `.DbMap.Close()

	if err := ` + parent.FirstLowerModel + `.DbMap.Where(
		query, args...).Find(` + parent.FirstLowerModel + `).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NotFoundError
		}
		return err
	}
	return nil
}
`
	strs = append(strs, selectSingle)
	selectSingleOrder := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（多个自定义复杂条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) SelectOrderQuery(query interface{}, sortorder string, args ...interface{}) error {
	` + parent.FirstLowerModel + `.ReadModelDbMap(` + parent.FirstLowerModel + `)
	defer ` + parent.FirstLowerModel + `.DbMap.Close()

	if err := ` + parent.FirstLowerModel + `.DbMap.Where(
		query, args...).Order(sortorder).Limit(1).Find(` + parent.FirstLowerModel + `).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NotFoundError
		}
		return err
	}
	return nil
}
`
	strs = append(strs, selectSingleOrder)

	selectById := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取单条数据（主键标识简单条件查询）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) SelectById(id interface{}) error {
	
	` + parent.FirstLowerModel + `.ReadModelDbMap(` + parent.FirstLowerModel + `)
	defer ` + parent.FirstLowerModel + `.DbMap.Close()

	query := map[string]interface{}{
		"` + parent.TableKey + `": id,
	}

	if err := ` + parent.FirstLowerModel + `.DbMap.Find(` + parent.FirstLowerModel + `, query).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NotFoundError
		}
		return err
	}

	return nil
}
`
	strs = append(strs, selectById)
	selectAll := `
/* ================================================================================
 * query: []interface{} || map[string]interface{} || string
 * args: if string: interface{}...
 * ================================================================================ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) SelectAll(paging *ging.Paging, query interface{}, args ...interface{}) ([]*` + parent.ModelStuctName + `, error) {
	` + parent.FirstLowerModel + `.ReadModelDbMap(` + parent.FirstLowerModel + `)
	defer ` + parent.FirstLowerModel + `.DbMap.Close()

	var ` + parent.FirstLowerModelListName + ` = make([]*` + parent.ModelStuctName + `, 0)
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
				err = ` + parent.FirstLowerModel + `.DbMap.Model(` + parent.FirstLowerModel + `).
					Where(query, args...).
					Count(&paging.TotalRecord).
					Order(paging.Sortorder).
					Offset(paging.Offset()).
					Limit(paging.PagingSize).
					Find(&` + parent.FirstLowerModelListName + `).Error
			} else {
				err = ` + parent.FirstLowerModel + `.DbMap.Model(` + parent.FirstLowerModel + `).
					Where(query, args...).
					Group(paging.Group).
					Count(&paging.TotalRecord).
					Order(paging.Sortorder).
					Offset(paging.Offset()).
					Limit(paging.PagingSize).
					Find(&` + parent.FirstLowerModelListName + `).Error
			}

			paging.SetTotalRecord(paging.TotalRecord)
		} else {
			if len(paging.Group) == 0 {
				err = ` + parent.FirstLowerModel + `.DbMap.Model(` + parent.FirstLowerModel + `).
					Where(query, args...).
					Order(paging.Sortorder).
					Find(&` + parent.FirstLowerModelListName + `).Error
			} else {
				err = ` + parent.FirstLowerModel + `.DbMap.Model(` + parent.FirstLowerModel + `).
					Where(query, args...).
					Group(paging.Group).
					Order(paging.Sortorder).
					Find(&` + parent.FirstLowerModelListName + `).Error
			}
		}
	} else {
		err = ` + parent.FirstLowerModel + `.DbMap.Where(query, args...).Find(&` + parent.FirstLowerModelListName + `).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = common.NotFoundError
		}
	}

	return ` + parent.FirstLowerModelListName + `, err
}`
	strs = append(strs, selectAll)
	insertId := ""
	if parent.TableKeyType == "string" {
		insertId = `if len(` + parent.FirstLowerModel + `.` + parent.TableKeyModelName + `)==0{
			return common.ArgumentError
		}`
	} else {
		insertId = `if` + parent.FirstLowerModel + `. ` + parent.TableKeyModelName + `==0{
			return common.ArgumentError
		}`
	}
	insert := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 插入数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) Insert() error {
    ` + insertId + `
	if ` + parent.FirstLowerModel + `.DbMap == nil {
		` + parent.FirstLowerModel + `.WriteModelDbMap(` + parent.FirstLowerModel + `)
		defer ` + parent.FirstLowerModel + `.DbMap.Close()
	}

	if err := ` + parent.FirstLowerModel + `.DbMap.Create(` + parent.FirstLowerModel + `).Error; err != nil {
		return err
	}
	return nil
}`
	strs = append(strs, insert)
	updateId := ""
	if parent.TableKeyType == "string" {
		updateId = `
if len(` + parent.FirstLowerModel + `.` + parent.TableKeyModelName + `) == 0 || len(data) == 0 {
		return 0, common.ArgumentError
	}
`
	} else {
		updateId = `
if ` + parent.FirstLowerModel + `.` + parent.TableKeyModelName + ` == 0 || len(data) == 0 {
		return 0, common.ArgumentError
	}
`
	}
	update := `
/* ================================================================================
 * data type:
 * Model{"fieldName":"value"...}
 * map[string]interface{}
 * key1,value1,key2,value2
 * ================================================================================ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) Update(data ...interface{}) (int64, error) {
	` + updateId + `

	if ` + parent.FirstLowerModel + `.DbMap == nil {
		` + parent.FirstLowerModel + `.WriteModelDbMap(` + parent.FirstLowerModel + `)
		defer ` + parent.FirstLowerModel + `.DbMap.Close()
	}

	dbContext := ` + parent.FirstLowerModel + `.DbMap.Model(` + parent.FirstLowerModel + `).UpdateColumns(data)
	rowsAffected, err := dbContext.RowsAffected, dbContext.Error

	return rowsAffected, err
}
`
	strs = append(strs, update)
	deleteId := ""
	if parent.TableKeyType == "string" {
		deleteId = `if len(` + parent.FirstLowerModel + `.` + parent.TableKeyModelName + `)==0{
			return 0,common.ArgumentError
		}`
	} else {
		deleteId = `if ` + parent.FirstLowerModel + `.` + parent.TableKeyModelName + `==0{
			return 0,common.ArgumentError
		}`
	}
	delete := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 删除数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (` + parent.FirstLowerModel + " *" + parent.ModelStuctName + `) Delete() (int64, error) {
     ` + deleteId + `	

	if ` + parent.FirstLowerModel + `.DbMap == nil {
		` + parent.FirstLowerModel + `.WriteModelDbMap(` + parent.FirstLowerModel + `)
		defer ` + parent.FirstLowerModel + `.DbMap.Close()
	}

	dbContext := ` + parent.FirstLowerModel + `.DbMap.Delete(` + parent.FirstLowerModel + `)
	rowsAffected, err := dbContext.RowsAffected, dbContext.Error

	return rowsAffected, err
}`
	strs = append(strs, delete)

	//写入文件
	result := &model.WriteModel{
		FileName: parent.TableKey,
		Path:     "./file/model/" + parent.TableName + ".go",
		Content:  strs,
	}

	return result
}
