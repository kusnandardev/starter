package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// Server :
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting :
var ServerSetting = &Server{}

// Database :
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Port        string
	Name        string
	TablePrefix string
}

// DatabaseSetting :
var DatabaseSetting = &Database{}

// App :
type App struct {
	JwtSecret string
	PageSize  int
	PrefixURL string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	Issuer string
}

// AppSetting interface pointer
var AppSetting = &App{}

// SMTP :
type SMTP struct {
	SMTPServer      string
	SMTPPort        int
	SMTPUser        string
	SMTPPasswd      string
	SMTPIdentity    string
	SMTPSenderEmail string
}

// SMTPSetting :
var SMTPSetting = &SMTP{}

// MongoDB :
type MongoDB struct {
	Type     string
	User     string
	Password string
	Host     string
}

// MongoDBSetting :
var MongoDBSetting = &MongoDB{}

// RedisDB :
type RedisDB struct {
	Host string
	Port int
}

// RedisDBSetting :
var RedisDBSetting = &RedisDB{}

var cfg *ini.File

// Setup Load config from ini file
func Setup() {
	now := time.Now()
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("smtp", SMTPSetting)
	mapTo("mongodb", MongoDBSetting)
	mapTo("redisdb", RedisDBSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	// fmt.Println("Config Setup is Ready...")
	timeSpent := time.Since(now)
	log.Printf("Config setting is ready in %v", timeSpent)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
