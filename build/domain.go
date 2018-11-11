package build

import (
	"gcrud/model"
	"fmt"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 生成业务领域模型数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func BuildDomain(parent *model.BussinessParent) *model.WriteModel {
	strs := make([]string, 0)

	//导入模块
	imports := `
	package post

	import (
		"github.com/hicsgo/ging"
	)
`
	strs = append(strs, imports)

	title := `
/* ================================================================================
 * ` + parent.TitleName + `映射数据域结构
 * author: jcheng
 * ================================================================================ */
type ` + parent.DomainStructName + `List []*` + parent.DomainStructName + `
type ` + parent.DomainStructName + ` struct {
`
	strs = append(strs, title)

	for _, val := range parent.Childs {
		strs = append(strs, `` + val.DomainName + ` ` + val.DomainType + ` ` + fmt.Sprintf("`form:\"%v\" json:\"%v\"`//%v", val.ColumnName, val.ColumnName, val.DomainTag)+ `
`)
	}
	strs = append(strs, `
}`)
	conditionTitle := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * ` + parent.TitleName + `映射查询数据域结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type ` + parent.DomainStructName + `Condition struct {
`

	strs = append(strs, conditionTitle)
	for _, val := range parent.Childs {

		//生成Max Min字段
		if val.IsCondition && val.ConditionType == model.CONDITION_UEQUAL {
			strs = append(strs, `Min` + val.ModelName + `  ` + val.DomainType + ` ` + fmt.Sprintf("`form:\"min_%v\" json:\"min_%v\"`", val.ColumnName, val.ColumnName)+ `
Max`+ val.ModelName+ `  `+ val.ModelType+ ` `+ fmt.Sprintf("`form:\"max_%v\" json:\"max_%v\"`", val.ColumnName, val.ColumnName)+ `
`)
		}
		if val.IsCondition && val.ConditionType == model.CONDITION_BOOL {
			strs = append(strs, `` + val.ModelName + `  *` + val.DomainType + ` ` + fmt.Sprintf("`form:\"%v\" json:\"%v\"`", val.ColumnName, val.ColumnName)+ `
`)
		}
		if val.IsCondition && val.ConditionType == model.CONDTION_EQUAL {
			strs = append(strs, `` + val.ModelName + `  ` + val.DomainType + ` ` + fmt.Sprintf("`form:\"%v\" json:\"%v\"`", val.ColumnName, val.ColumnName)+ `
`)
		}
	}
	strs = append(strs, `Paging  *ging.Paging  ` + fmt.Sprintf("`form:\"paging\" json:\"paging\"`")+ `
}
`)

	//写入文件
	result := &model.WriteModel{
		FileName: parent.TableKey,
		Path:     "./file/domain/" + parent.TableName + ".go",
		Content:  strs,
	}

	return result
}
