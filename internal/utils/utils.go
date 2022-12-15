package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/ai/stt/v2"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func YandexSttRecognition(ctx context.Context, sdk *ycsdk.SDK, audioFileUrl string, resultFilePath string) error {

	op, err := sdk.WrapOperation(YandexSttCreateRequest(ctx, sdk, audioFileUrl))
	if err != nil {
		fmt.Printf("ошибка запуска опрерации транскрибирования: %s", err)
	}
	for {
		err = op.Poll(ctx)
		if err != nil {
			fmt.Printf("ошибка обновления статуса опрерации транскрибирования: %s", err)
		}

		if op.Done() {
			break
		}

		time.Sleep(5 * time.Second)
	}
	resp, err := op.Response()
	if err != nil {
		fmt.Printf("ошибка получения результата опрерации транскрибирования: %s", err)
	}
	sttResponse := resp.(*stt.LongRunningRecognitionResponse)
	// открываем файл для записи результатов
	resultFile, err := os.OpenFile(resultFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("ошибка открытия файла для записи результатов: %s", err)
	}
	defer func() {
		_ = resultFile.Close()
	}()

	for _, chunk := range sttResponse.Chunks {
		result := chunk.GetAlternatives()
		resultString := result[0].Text
		_, err = resultFile.WriteString(fmt.Sprintf("%s\n", resultString))
		if err != nil {
			fmt.Printf("ошибка записи в файл строки [%s]: %s", resultString, err)
		}
	}

	return nil
}

func YandexSttCreateRequest(ctx context.Context, sdk *ycsdk.SDK, audioUri string) (*operation.Operation, error) {

	request := &stt.LongRunningRecognitionRequest{
		Config: &stt.RecognitionConfig{
			Specification: &stt.RecognitionSpec{
				LanguageCode:    "ru-RU",
				ProfanityFilter: true,
				PartialResults:  true,
				AudioEncoding:   stt.RecognitionSpec_MP3,
				//SampleRateHertz: 8000,
			},
		},
		Audio: &stt.RecognitionAudio{
			AudioSource: &stt.RecognitionAudio_Uri{
				Uri: audioUri,
			},
		},
	}

	op, err := sdk.AI().STT().Stt().LongRunningRecognize(ctx, request)
	return op, err
}
