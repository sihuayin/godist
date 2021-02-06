package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	DatabaseUser        string
	DatabasePWD         string
	DatabaseHost        string
	DatabasePort        string
	DatabaseDBName      string
	DatabaseDBPRE       string
	DatabaseMaxIdleConn int
	DatabaseMaxOpenConn int
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadDatabase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection(RunMode)
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadDatabase() {
	sec, err := Cfg.GetSection(RunMode)
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	DatabaseUser = sec.Key("DATABASE_USER").MustString("root")
	DatabasePWD = sec.Key("DATABASE_PASS").MustString("123456")
	DatabaseHost = sec.Key("DATABASE_HOST").MustString("localhost")
	DatabasePort = sec.Key("DATABASE_PORT").MustString("8000")
	DatabaseDBName = sec.Key("DATABASE_DB_NAME").MustString("walle")
	DatabaseDBPRE = sec.Key("DATABASE_TABLE_PREFIX").MustString("")
	DatabaseMaxIdleConn = sec.Key("DATABASE_DB_MAX_IDLE_CONN").MustInt(100)
	DatabaseMaxOpenConn = sec.Key("DATABASE_DB_OPEN_CONN").MustInt(20)
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
