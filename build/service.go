package build

import (
	"gcrud/model"
)

func BuildService(parent *model.BussinessParent) *model.WriteModel {
	strs := make([]string, 0)
	selectById := `
package service

import (
	DOMAIN_QQQQQQ "hecate/domain/post"
	MODEL_WWWWWW "hecate/model/post"
	"time"
)
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据Id标识获取数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *` + parent.ServiceName + `) Get` + parent.DomainStructName + `ById(id int64) (*` + parent.PkgDomain + `.` + parent.DomainStructName + `, error) {
	if id == 0 {
		return nil, common.ArgumentError
	}

	` + parent.FirstLowerModel + ` := new(` + parent.PkgModel + `.` + parent.ModelStuctName + `)
	if err := ` + parent.FirstLowerModel + `.SelectById(id); err != nil {
		return nil, err
	}

	return ` + parent.ModelStuctName + `To` + parent.DomainStructName + `(` + parent.FirstLowerModel + `), nil
}
`
	strs = append(strs, selectById)
	lists := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *` + parent.ServiceName + `) Get` + parent.DomainStructName + `List(` + parent.FirstLowerDomainName + `Condition *` + parent.PkgDomain + `.` + parent.DomainStructName + `Condition) (` + parent.PkgDomain + `.` + parent.DomainStructName + `List, error) {
	if ` + parent.FirstLowerDomainName + `Condition == nil {
		return nil, common.ArgumentError
	}

	fields := make([]string, 0)
	args := make([]interface{}, 0)

	

`
	strs = append(strs, lists)
	for _, val := range parent.Childs {
		//判断string类型的等于条件
		if val.IsCondition && val.ConditionType == model.CONDTION_EQUAL && val.ModelType == "string" {
			strs = append(strs, `
if len(`+ parent.FirstLowerDomainName+ `Condition.`+ val.ModelName+ `) > 0 {
		fields = append(fields, "`+ val.ColumnName+ ` = ?")
		args = append(args, `+ parent.FirstLowerModel+ `Condition.`+ val.ModelName+ `)
	}
`)
		}
		//判断非string类的等于条件
		if val.IsCondition && val.ConditionType == model.CONDTION_EQUAL && val.ModelType != "string" {
			strs = append(strs, `
if `+ parent.FirstLowerDomainName+ `Condition.`+ val.ModelName+ ` > 0 {
		fields = append(fields, "`+ val.ColumnName+ ` = ?")
		args = append(args, `+ parent.FirstLowerModel+ `Condition.`+ val.ModelName+ `)
	}
`)
		}
		//判断时间类型的大于小于条件
		if val.IsCondition && val.ConditionType == model.CONDITION_UEQUAL && val.ModelType == "time.Time" {
			strs = append(strs, `
if !`+ parent.FirstLowerDomainName+ `Condition.Min`+ val.ModelName+ `.IsZero()  {
		fields = append(fields, "`+ val.ColumnName+ ` >= ?")
		args = append(args, `+ parent.FirstLowerModel+ `Condition.Min`+ val.ModelName+ `)
	}
if !`+ parent.FirstLowerDomainName+ `Condition.Max`+ val.ModelName+ `.IsZero()  {
		fields = append(fields, "`+ val.ColumnName+ ` < ?")
		args = append(args, `+ parent.FirstLowerModel+ `Condition.Max`+ val.ModelName+ `)
	}
`)
		}

		//判断bool类型的条件
		if val.IsCondition && val.ConditionType == model.CONDITION_BOOL {
			strs = append(strs, `
if `+ parent.FirstLowerDomainName+ `Condition.`+ val.ModelName+ `!=nil {
        isTrue:=0
         if *`+ parent.FirstLowerDomainName+ `Condition.`+ val.ModelName+ ` {
         			isTrue = 1
         		}
		fields = append(fields, "`+ val.ColumnName+ ` = ?")
		args = append(args,isTrue)
	}
`)
		}
	}
	strs = append(strs, `

	query := glib.StringSliceToString(fields, " AND ")

	`+ parent.FirstLowerModel+ ` := new(`+ parent.PkgModel+ `.`+ parent.ModelStuctName+ `)
	`+ parent.FirstLowerModel+ `List, err := `+ parent.FirstLowerModel+ `.SelectAll(`+ parent.FirstLowerDomainName+ `Condition.Paging, query, args...)
	if err != nil {
		return nil, err
	}

	return `+ parent.ModelStuctName+ `ListTo`+ parent.DomainStructName+ `List(`+ parent.FirstLowerModel+ `List), nil
}
`)
	//创建数据
	create := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 创建数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *` + parent.ServiceName + `) Create` + parent.DomainStructName + `(` + parent.FirstLowerDomainName + ` *` + parent.PkgDomain + `.` + parent.DomainStructName + `) error {
	if ` + parent.DomainStructName + ` == nil ||` + parent.DomainStructName + `.Id==0 { 
		return common.ArgumentError
	}

	//插入数据
	` + parent.FirstLowerModel + ` := ` + parent.DomainStructName + `To` + parent.ModelStuctName + `(` + parent.DomainStructName + `)
	if err := ` + parent.FirstLowerModel + `.Insert(); err != nil {
		return err
	}

	` + parent.DomainStructName + `.` + parent.TableKeyModelName + ` = ` + parent.FirstLowerModel + `.` + parent.TableKeyModelName + `

	return nil
}
`
	strs = append(strs, create)

	update := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 更新数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *` + parent.ServiceName + `) Update` + parent.DomainStructName + `ById(id int64, data map[string]interface{}) error {
	if id == 0 || len(data) == 0 {
		return common.ArgumentError
	}

	` + parent.FirstLowerModel + ` := new(` + parent.PkgModel + `.` + parent.ModelStuctName + `)
	` + parent.FirstLowerModel + `.` + parent.TableKeyModelName + ` = id
	_, err := ` + parent.FirstLowerModel + `.Update(data)

	return err
}
`
	strs = append(strs, update)

	delete := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 删除数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *` + parent.ServiceName + `) Delete` + parent.DomainStructName + `ById(id int64) error {
	if id == 0 {
		return common.ArgumentError
	}

	` + parent.FirstLowerModel + ` := new(` + parent.PkgModel + `.` + parent.ModelStuctName + `)
	` + parent.FirstLowerModel + `.` + parent.TableKeyModelName + ` = id
	_, err := ` + parent.FirstLowerModel + `.Delete()

	return err
}
`
	strs = append(strs, delete)
	//写入文件
	result := &model.WriteModel{
		FileName: parent.TableKey,
		Path:     "./file/service/" + parent.TableName + ".go",
		Content:  strs,
	}

	return result
}
