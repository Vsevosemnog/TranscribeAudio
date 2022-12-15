package models

import (
	"transcribeaudio/internal/telegram"

	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/bkmz/go-rabbitmq-connector"
	"github.com/go-co-op/gocron"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"gorm.io/gorm"
)

type AppConfigure struct {
	DB          *gorm.DB
	YandexCloud struct {
		S3Client *s3.Client
		SDK      *ycsdk.SDK
	}
	AWS struct {
		KeyID              string
		KeySecret          string
		S3Client           *s3.Client
		MediaConvertClient *mediaconvert.Client
	}
	AWSConfig aws.Config
	BOT       telegram.BOT
	ZOOM      struct {
		Key    string
		Secret string
	}
	Scheduler *gocron.Scheduler
	AMQP      struct {
		Channel *rabbitmq.Channel
	}
}

type T struct {
	DownloadToken string `json:"download_token"`
	Event         string `json:"event"`
	EventTs       int64  `json:"event_ts"`
	Payload       struct {
		AccountId string `json:"account_id"`
		Object    struct {
			AccountId      string      `json:"account_id"`
			Duration       int         `json:"duration"`
			HostEmail      string      `json:"host_email"`
			HostId         string      `json:"host_id"`
			Id             int64       `json:"id"`
			OnPrem         interface{} `json:"on_prem"`
			Password       string      `json:"password"`
			RecordingCount int         `json:"recording_count"`
			RecordingFiles []struct {
				DownloadUrl    string    `json:"download_url"`
				FileExtension  string    `json:"file_extension"`
				FileSize       int       `json:"file_size"`
				FileType       string    `json:"file_type"`
				Id             string    `json:"id"`
				MeetingId      string    `json:"meeting_id"`
				PlayUrl        string    `json:"play_url"`
				RecordingEnd   time.Time `json:"recording_end"`
				RecordingStart time.Time `json:"recording_start"`
				RecordingType  string    `json:"recording_type"`
				Status         string    `json:"status"`
			} `json:"recording_files"`
			ShareUrl  string    `json:"share_url"`
			StartTime time.Time `json:"start_time"`
			Timezone  string    `json:"timezone"`
			Topic     string    `json:"topic"`
			TotalSize int       `json:"total_size"`
			Type      int       `json:"type"`
			Uuid      string    `json:"uuid"`
		} `json:"object"`
	} `json:"payload"`
	RequestId string `json:"request_id"`
}
