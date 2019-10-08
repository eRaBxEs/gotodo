package util

// DatabaseInfo ...
type DatabaseInfo struct {
	User     string `sane:"user"`
	DBName   string `sane:"dbname"`
	Password string `sane:"password"`
}

// ConfigFile ...
type ConfigFile struct {
	Host   string       `sane:"host"`
	Port   string       `sane:"port"`
	DBInfo DatabaseInfo `sane:"dbinfo"`
}
