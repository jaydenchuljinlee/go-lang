package config

import (
	"log"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type Config struct {
	Google GoogleConfig `mapstructure:"google"`
}

type GoogleConfig struct {
	APIKey       string `mapstructure:"api_key"`
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	LoginURI     string `mapstructure:"login_uri"`
	AuthURI      string `mapstructure:"auth_uri"`
	RedirectURI  string `mapstructure:"redirect_uri"`
}

var AppConfig Config
var OAuthConfig *oauth2.Config

func LoadConfig() {
	viper.SetConfigName("config")   // 설정 파일의 이름 지정
	viper.SetConfigType("yaml")     // 설정 파일 유형은 YAML
	viper.AddConfigPath("./config") // 설정 파일 경로 추가

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	log.Printf("Loaded configuration: %+v\n", AppConfig.Google)
	// OAuth 설정 생성
	OAuthConfig = &oauth2.Config{
		ClientID:     AppConfig.Google.ClientID,
		ClientSecret: AppConfig.Google.ClientSecret,
		RedirectURL:  AppConfig.Google.RedirectURI,
		Endpoint: oauth2.Endpoint{
			AuthURL:  AppConfig.Google.LoginURI,
			TokenURL: AppConfig.Google.AuthURI,
		},
		Scopes: []string{"https://mail.google.com/", "https://www.googleapis.com/auth/pubsub"},
	}

}
