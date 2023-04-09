package config

type AppSettings struct {
	Name    string
	Version string
	Debug   bool
}
type DatabaseSettings struct {
	DbType   string
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type LogSettings struct {
	Level  string
	Format string
	File   string
}

type Settings struct {
	App      AppSettings
	Database DatabaseSettings
	Log      LogSettings
}
