package router

import (
	"context"
	"fmt"
	"net/http"
	"sso/config"
	"strings"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"

	// "google.golang.org/api/pubsub/v1"

	"golang.org/x/oauth2" // oauth2 패키지 추가
	"golang.org/x/oauth2/google"
)

var log = logrus.New()

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/v1/google/watcher", watchMail)
	router.GET("/api/v1/google/auth", handleMain)
	router.GET("/api/v1/google/auth/success", handleCallback)

	return router
}

// LoadServiceAccountJSON loads the service account JSON key file and returns an HTTP client
func LoadServiceAccountJSON(ctx context.Context, filePath, userEmail string) (*http.Client, error) {
	// JSON 키 파일 로드
	jsonKey, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// JWT Config 생성
	config, err := google.JWTConfigFromJSON(jsonKey, gmail.GmailReadonlyScope)
	if err != nil {
		return nil, err
	}

	// 특정 사용자의 이메일 계정으로 동작하도록 설정
	config.Subject = userEmail
	client := config.Client(ctx)
	return client, nil
}

func tokenFromHeader(c *gin.Context) (*oauth2.Token, error) {
	// Authorization 헤더 가져오기
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	// Bearer 토큰 분리
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("authorization header format must be 'Bearer {token}'")
	}

	return &oauth2.Token{AccessToken: parts[1]}, nil
}

func watchMail(c *gin.Context) {
	ctx := context.Background()
	token, err := tokenFromHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	jsonFilePath := "config/gmail-service-account.json"
	userEmail := "asddqe111@bbrick.com"

	// 사용자 OAuth2 토큰을 사용하여 HTTP 클라이언트 생성
	oauth2Client := config.OAuthConfig.Client(ctx, token)

	// Gmail 서비스 생성 (사용자 OAuth2 클라이언트를 사용)
	gmailService, err := gmail.NewService(ctx, option.WithHTTPClient(oauth2Client))
	if err != nil {
		log.Printf("Unable to create Gmail service: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// 서비스 계정 클라이언트 생성
	serviceAccountClient, err := LoadServiceAccountJSON(ctx, jsonFilePath, userEmail)
	if err != nil {
		log.Printf("Failed to load service account JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// WatchRequest 설정
	watchRequest := &gmail.WatchRequest{
		LabelIds:  []string{"INBOX"},
		TopicName: "projects/goyoai-portal/topics/email-109",
	}

	// 서비스 계정 클라이언트를 사용하여 Watch 요청을 보냅니다.
	watchResponse, err := gmail.NewService(ctx, option.WithHTTPClient(serviceAccountClient)).Users.Watch("me", watchRequest).Do()
	if err != nil {
		log.Printf("Unable to create watch: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	log.Printf("Subscription created, expiration: %v", watchResponse.Expiration)
	c.JSON(http.StatusOK, gin.H{"message": "Watch created successfully", "expiration": watchResponse.Expiration})
}

func handleMain(c *gin.Context) {
	url := config.OAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func handleCallback(c *gin.Context) {
	ctx := context.Background()
	code := c.Query("code")
	token, err := config.OAuthConfig.Exchange(ctx, code)
	if err != nil {
		log.Errorf("Unable to retrieve token from web: %v", err)
		return
	}

	log.Println(token.AccessToken)

	client := config.OAuthConfig.Client(ctx, token)
	gmailService, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Errorf("Unable to create Gmail service: %v", err)
		return
	}

	// // Pub/Sub 서비스 초기화
	// pubsubService, err := pubsub.NewService(ctx, option.WithHTTPClient(client))
	// if err != nil {
	// 	log.Errorf("Unable to create PubSub client: %v", err)
	// }

	// // Pub/Sub 토픽 생성
	// topicName := "projects/goyoai-portal/topics/email-109"
	// topic := &pubsub.Topic{Name: topicName}
	// _, err = pubsubService.Projects.Topics.Create(topicName, topic).Do()
	// if err != nil {
	// 	log.Errorf("Unable to create topic: %v", err)
	// }

	user := "me"
	r, err := gmailService.Users.Messages.List(user).Do()
	if err != nil {
		log.Errorf("Unable to retrieve messages: %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, r.Messages)
}
