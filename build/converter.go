package build

import (
	"gcrud/model"
)

func Converter(parent *model.BussinessParent) *model.WriteModel {
	strs := make([]string, 0)

	//modeltodomain
	modeltodomain := `
package model

import (
	DOMAIN_QQQQQQ "hecate/domain/post"
	MODEL_WWWWWW "hecate/model/post"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 *` + parent.ModelStuctName + `To` + parent.DomainStructName + ` 
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
 func ` + parent.ModelStuctName + `To` + parent.DomainStructName + `(` + parent.FirstLowerModel + ` *` + parent.PkgModel + `.` + parent.ModelStuctName + `) *` + parent.PkgDomain + `.` + parent.DomainStructName + ` {
` + parent.FirstLowerDomainName + `:= new(` + parent.PkgDomain + `.` + parent.DomainStructName + `)
if ` + parent.FirstLowerModel + ` == nil {
		return ` + parent.FirstLowerDomainName + `
	}
`
	strs = append(strs, modeltodomain)
	for _, val := range parent.Childs {
		strs = append(strs, `` + parent.FirstLowerDomainName + `.` + val.DomainName + `=` + parent.FirstLowerModel + `.` + val.ModelName+ `
`)

	}
	strs = append(strs, `
return `+ parent.FirstLowerDomainName+ `
}
`)

	//domaintomodel
	domaintomodel := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 *` + parent.DomainStructName + `To` + parent.ModelStuctName + ` 
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
 func ` + parent.DomainStructName + `To` + parent.ModelStuctName + `(` + parent.FirstLowerDomainName + ` *` + parent.PkgDomain + `.` + parent.DomainStructName + `) *` + parent.PkgModel + `.` + parent.ModelStuctName + ` {
` + parent.FirstLowerModel + `:= new(` + parent.PkgModel + `.` + parent.ModelStuctName + `)
if ` + parent.FirstLowerDomainName + ` == nil {
		return ` + parent.FirstLowerModel + `
	}
`
	strs = append(strs, domaintomodel)

	//处理Id
	if parent.TableKeyType == "string" {
		strs = append(strs, `
`+ parent.FirstLowerModel+ `.`+ parent.TableKeyModelName+ ` = `+ parent.FirstLowerModel+ `.`+ parent.TableKeyModelName+ `
	if len(`+ parent.FirstLowerModel+ `.`+ parent.TableKeyModelName+ `) == 0 {
		guid, _ := uuid.NewV4()
		`+ parent.FirstLowerModel+ `.`+ parent.TableKeyModelName+ ` = guid.String()
	}
`)
		//处理内部字段
		for _, val := range parent.Childs {
			if val.ModelType == "time.Time" {
				strs = append(strs, `` + parent.FirstLowerModel + `.` + val.ModelName + ` = ` + parent.FirstLowerDomainName + `.` + val.DomainName+ `

	if `+ parent.FirstLowerModel+ `.`+ val.ModelName+ `.IsZero() {
		`+ parent.FirstLowerModel+ `.`+ val.ModelName+ ` = time.Now()
	}
`)
				continue
			} else {
				strs = append(strs, `` + parent.FirstLowerModel + `.` + val.ModelName + `=` + parent.FirstLowerDomainName + `.` + val.DomainName+ `
`)
			}
		}

	}
	strs = append(strs, `
return `+ parent.FirstLowerModel+ `
}`)

	//modelListToDomainList
	list := `
/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * ` + parent.ModelStuctName + `ListTo` + parent.DomainStructName + `List
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ` + parent.ModelStuctName + `ListTo` + parent.DomainStructName + `List(` + parent.FirstLowerModel + `List []*` + parent.PkgModel + `.` + parent.ModelStuctName + `) ` + parent.PkgDomain + `.` + parent.DomainStructName + `List {
	` + parent.FirstLowerDomainName + `List := make(` + parent.PkgDomain + `.` + parent.DomainStructName + `List, 0, len(` + parent.FirstLowerModel + `List))

	for _, ` + parent.FirstLowerModel + `:= range ` + parent.FirstLowerModel + `List {
		` + parent.FirstLowerDomainName + ` := ` + parent.ModelStuctName + `To` + parent.DomainStructName + `(` + parent.FirstLowerModel + `)
		` + parent.FirstLowerDomainName + `List = append(` + parent.FirstLowerDomainName + `List, ` + parent.FirstLowerDomainName + `)
	}

	return ` + parent.FirstLowerDomainName + `List
}
`
	strs = append(strs, list)

	//写入文件
	result := &model.WriteModel{
		FileName: parent.TableKey,
		Path:     "./file/converter/" + parent.TableName + ".go",
		Content:  strs,
	}

	return result
}
