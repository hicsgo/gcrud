package run

import (
	"log"
	"strings"
	"fmt"
)

import (
	"github.com/hicsgo/glib"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

import (
	"database/sql"
	"gcrud/model"
	"gcrud/build"
	"os"
)

/* ================================================================================
 * 数据中心
 * author: jcheng
 * ================================================================================ */

var (
	DB             *sql.DB
	DatabaseConfig *model.DatabaseConfig
)

func Run() {
	DB, _ = sql.Open("mysql", ""+DatabaseConfig.UserName+":"+DatabaseConfig.PassWord+"@("+DatabaseConfig.Host+":"+DatabaseConfig.Port+")/"+DatabaseConfig.DatabaseName+"?charset=utf8&timeout=3s&parseTime=True&loc=Local&interpolateParams=true")
	data, err := GetTableData()

	if err == nil {
		GetTableColumnsData(data)
		results := GetBusData(data)
		for i, v := range results {
			if i == 2 {
				break
			}
			WriteFile(build.BuildModel(v))
			WriteFile(build.BuildDomain(v))
			WriteFile(build.Converter(v))
			WriteFile(build.BuildService(v))
			WriteFile(build.BuildIService(v))
		}

	} else {
		fmt.Println(err)
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前库下的所有表数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetTableData() (Tables []*model.Table, err error) {
	sql := "SELECT table_name name,TABLE_COMMENT value FROM INFORMATION_SCHEMA.TABLES WHERE table_type='base table'and table_schema = '" + DatabaseConfig.DatabaseName + "'"
	columns := make([]*model.Table, 0)
	rows, _ := DB.Query(sql)
	defer rows.Close()
	for rows.Next() {
		var column = model.Table{}
		if err := rows.Scan(&column.TableName, &column.TableComment); err != nil {
			log.Fatal(err)
		}
		columns = append(columns, &column)
	}
	return columns, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前库下的所有表数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetTableColumnsData(Tables []*model.Table) {
	tableName := ""
	for _, t := range Tables {
		tableName = t.TableName
		var columns []*model.Column
		sql := `select COLUMN_NAME,DATA_TYPE,COLUMN_KEY,COLUMN_COMMENT from  information_schema.columns where TABLE_NAME='` + tableName + `'and table_schema = '` + DatabaseConfig.DatabaseName + `'`
		rows, _ := DB.Query(sql)
		defer rows.Close()
		for rows.Next() {
			var column = model.Column{}
			if err := rows.Scan(&column.ColumnName, &column.ColumnType, &column.ColumnKey, &column.ColumnComment); err != nil {
				log.Fatal(err)
			}
			columns = append(columns, &column)
		}
		t.Columns = columns
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据数据库模型装换成业务模型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetBusData(Tables []*model.Table) []*model.BussinessParent {
	bussinessList := make([]*model.BussinessParent, 0)
	for _, t := range Tables {
		bussiness := new(model.BussinessParent)
		bussiness.TableName = t.TableName
		bussiness.ModelStuctName = glib.ConvertName(t.TableName) + "Model"
		bussiness.DomainStructName = glib.ConvertName(t.TableName)
		bussiness.TableComment = t.TableComment
		bussiness.DBName = DatabaseConfig.DatabaseName
		bussiness.PkgDomain = "DOMAIN_QQQQQQ"
		bussiness.PkgModel = "MODEL_WWWWWW"
		bussiness.ServiceName = "Service_EEEEEE"
		bussiness.TitleName = strings.Replace(bussiness.TableComment, "表", "", -1) + "模型"
		bussiness.FirstLowerModel = glib.FirstToLower(bussiness.ModelStuctName)
		bussiness.FirstLowerModelListName = bussiness.FirstLowerModel + "List"
		bussiness.FirstLowerDomainName = glib.FirstToLower(bussiness.DomainStructName)
		bussiness.FirstLowerDomainListName = bussiness.FirstLowerDomainName + "List"
		clids := make([]*model.BusinessChild, 0)
		for i, c := range t.Columns {

			child := new(model.BusinessChild)
			child.ColumnKey = c.ColumnKey
			child.ColumnName = c.ColumnName
			child.ColumnType = c.ColumnType
			child.ColumnComment = c.ColumnComment
			child.ModelName = glib.ConvertName(c.ColumnName)

			modelType := fmt.Sprint(glib.GoTypeByMysqlType(c.ColumnType))
			if modelType == "struct" {
				child.ModelType = "time.Time"
			} else {
				child.ModelType = modelType
			}
			if c.ColumnKey == "PRI" || (c.ColumnKey == "KEY" && i == 0) {
				bussiness.TableKey = c.ColumnName
				bussiness.TableKeyType = child.ModelType
				bussiness.TableKeyModelName = child.ModelName
			}
			child.ModelTag = fmt.Sprintf("%s", c.ColumnComment)
			child.DomainName = child.ModelName
			child.DomainTag = child.ModelTag
			child.DomainType = child.ModelType

			isCondtion := false
			conditionType := ""
			if strings.Index(child.ColumnName, "id") != -1 ||
				strings.Index(child.ColumnName, "Id") != -1 ||
				strings.Index(child.ColumnName, "ID") != -1 ||
				strings.Index(child.ColumnName, "iD") != -1 {
				isCondtion = true
				conditionType = model.CONDTION_EQUAL
			}
			if (strings.Index(child.ColumnName, "date") != -1 && child.ModelType == "time.Time") || (strings.Index(child.ColumnName, "time") != -1 && child.ModelType == "time.Time") {
				isCondtion = true
				conditionType = model.CONDITION_UEQUAL
			}
			if strings.Index(child.ColumnName, "is") == 0 ||
				strings.Index(child.ColumnName, "Is") == 0 ||
				strings.Index(child.ColumnName, "IS") == 0 ||
				strings.Index(child.ColumnName, "Is") == 0 {
				isCondtion = true
				conditionType = model.CONDITION_BOOL
				child.DomainType = "bool"
				child.ModelType = "bool"
			}
			child.IsCondition = isCondtion
			child.ConditionType = conditionType
			clids = append(clids, child)
		}
		bussiness.Childs = clids
		bussinessList = append(bussinessList, bussiness)
	}
	return bussinessList
}

func WriteFile(writeModel *model.WriteModel) {
	//写入文件
	file6, error := os.Create(writeModel.Path);
	if error != nil {
		fmt.Println(error);
	}
	for _, v := range writeModel.Content {
		file6.Write([]byte(v));
	}
	file6.Close();
}
