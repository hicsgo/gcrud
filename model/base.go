package model

/* ================================================================================
 * 模型中心
 * author: jcheng
 * ================================================================================ */
const (
	CONDTION_EQUAL   = "equal"
	CONDITION_UEQUAL = "unequal"
	CONDITION_BOOL   = "bool"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库配置项
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type DatabaseConfig struct {
	DatabaseName string //数据库名称
	Host         string //host
	Port         string //端口
	UserName     string //用户名
	PassWord     string //密码
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 业务领域模型及数据模型及SQL值
 * 业务领域模型中特殊的需要自行处理
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type BussinessParent struct {
	DBName                   string           //db的名称 eg:tapai_db
	TableName                string           //表名称 eg:user
	TitleName                string           //模块描述名称
	TableComment             string           //表说明eg:用户表
	TableKey                 string           //表中索引
	TableKeyType             string           //表中索引类型
	TableKeyModelName        string           //表中索引模型名称
	ModelStuctName           string           //模型结构名称
	DomainStructName         string           //领域结构名称
	ServiceName              string           //服务名称
	PkgModel                 string           //model包别名
	PkgDomain                string           //domain包别名
	FirstLowerModel          string           //模型首字母小写
	FirstLowerModelListName  string           //模型集合首字母小写
	FirstLowerDomainName     string           //领域模型首字母小写
	FirstLowerDomainListName string           //领域模型集合首字母小写
	Childs                   []*BusinessChild //子集
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 业务领域模型及数据模型及SQL值
 * 业务领域模型中特殊的需要自行处理
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type BusinessChild struct {
	ColumnName    string //列名称 eg:user_name
	ColumnKey     string //列键属性 eg:pri
	ColumnType    string //列类型 eg:varchar(255)
	ColumnComment string //列说明 eg:1成功2失败
	ModelName     string //模型列名称: User

	ModelType  string //模型列类型:string
	ModelTag   string //模型列Tag //1成功2失败
	DomainName string //领域列名称:User

	DomainType    string //领域列类型:string
	DomainTag     string //领域列Tag //1成功2失败
	IsCondition   bool   //是否是查询字段
	ConditionType string //查询字段类型(等于、bool、大于小于)
	IsPri         bool   //是否是主键
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库表模型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type Table struct {
	TableName    string
	TableComment string
	Columns      []*Column
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库列模型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type Column struct {
	ColumnName    string //列名称 eg:user_name
	ColumnKey     string //列键属性 eg:pri
	ColumnType    string //列类型 eg:varchar(255)
	ColumnComment string //列说明 eg:1成功2失败
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 写入模型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type WriteModel struct {
	Path     string //写入文件路径
	FileName string //文件名称
	Content  []string //内容
}
