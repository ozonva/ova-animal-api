package config

type Settings struct {
	GrpcAddr string
	HttpAddr string
	Network  string
	Db       Db
}

type Db struct {
	Login    string
	Password string
	Host     string
	Port     int
	Name     string
}
