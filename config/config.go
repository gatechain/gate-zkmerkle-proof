package config

type Config struct {
	MysqlDataSource string
	UserDataFile    string
	DbSuffix        string
	TreeDB          struct {
		Driver string
		Option struct {
			Addr string
		}
	}
	Redis struct {
		Host     string
		Type     string
		Password string
	}
	ZkKeyName string
}
