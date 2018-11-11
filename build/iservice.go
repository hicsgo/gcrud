package build

import (
	"gcrud/model"
)

func BuildIService(parent *model.BussinessParent) *model.WriteModel {
	strs := make([]string, 0)
	iservice := `
package service

import (
	DOMAIN_QQQQQQ "hecate/domain/post"
	MODEL_WWWWWW "hecate/model/post"
	"time"
)
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * ` + parent.TableComment + `服务接口
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type I` + parent.DomainStructName + `Service interface{
    Get` + parent.DomainStructName + `ById(id int64) (*` + parent.PkgDomain + `.` + parent.DomainStructName + `, error)
	Get` + parent.DomainStructName + `List(` + parent.FirstLowerDomainName + `Condition *` + parent.PkgDomain + `.` + parent.DomainStructName + `Condition) (` + parent.PkgDomain + `.` + parent.DomainStructName + `List, error)
	Create` + parent.DomainStructName + `(` + parent.FirstLowerDomainName + ` *` + parent.PkgDomain + `.` + parent.DomainStructName + `) error
	Update` + parent.DomainStructName + `ById(id int64, data map[string]interface{}) error
	Delete` + parent.DomainStructName + `ById(id int64) error	
}
`
	strs = append(strs, iservice)
	//写入文件
	result := &model.WriteModel{
		FileName: parent.TableKey,
		Path:     "./file/iservice/" + parent.TableName + ".go",
		Content:  strs,
	}

	return result
}
