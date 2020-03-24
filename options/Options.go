package options

const (
	MYSQL        = "mysql"
	CUSTOM_COLUM = "customColum"
	CAML_COLUM   = "camlColum"
)

type Options struct {
	Host        string // 数据库地址
	Port        int    // 端口号
	Database    string // 数据名称
	UserName    string // 账号
	Password    string // 密码
	DriverName  string // 驱动名称
	ColumAssign string // 列命映射方式
}

type option func(options *Options)

func ColumAssign(columAssign string) option {
	return func(options *Options) {
		options.ColumAssign = columAssign
	}
}
func DriverName(driverName string) option {
	return func(options *Options) {
		options.DriverName = driverName
	}
}

func Host(host string) option {
	return func(options *Options) {
		options.Host = host
	}
}

func Port(port int) option {
	return func(options *Options) {
		options.Port = port
	}
}

func Database(database string) option {
	return func(options *Options) {
		options.Database = database
	}
}

func UserName(username string) option {
	return func(options *Options) {
		options.UserName = username
	}
}

func Password(password string) option {
	return func(options *Options) {
		options.Password = password
	}
}

func NewOptions(ops ...option) *Options {
	options := &Options{
		Host:        "jdbc://localhost",
		Port:        3306,
		Database:    "localhost",
		UserName:    "root",
		Password:    "root",
		DriverName:  MYSQL,
		ColumAssign: CAML_COLUM,
	}

	for _, setting := range ops {
		setting(options)
	}
	return options
}
