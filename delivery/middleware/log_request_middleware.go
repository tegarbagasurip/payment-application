package middleware

import (
	"os"
	"payment-application/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RequestLog struct {
	StartTime  time.Time
	EndTime    time.Duration
	StatusCode int
	ClientIP   string
	Method     string
	Path       string
	UserAgent  string
}

func LogRequestMiddleware(log *logrus.Logger) gin.HandlerFunc {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("Error Get Config", err.Error())
	}

	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	if err != nil {
		log.Fatalln("Error Get Config", err.Error())
	}

	return func(ctx *gin.Context) {
		endTime := time.Since(time.Now())
		requestLog := RequestLog{
			StartTime:  time.Now(),
			EndTime:    endTime,
			StatusCode: ctx.Writer.Status(),
			ClientIP:   ctx.ClientIP(),
			Method:     ctx.Request.Method,
			Path:       ctx.Request.URL.Path,
			UserAgent:  ctx.Request.UserAgent(),
		}

		switch {
		case ctx.Writer.Status() >= 500:
			log.Error(requestLog)
		case ctx.Writer.Status() >= 400:
			log.Warn(requestLog)
		default:
			log.Info(requestLog)
		}

	}
}
