package telegram

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//func SendMessage(ctx context.Context, botToken string, m interface{}) error {
//	const METHOD = "sendMessage"
//	var URL = fmt.Sprintf("https://api.telegram.org/bot%s/%s", botToken, METHOD)
//
//	jsonStr, err := json.Marshal(m)
//	if err != nil {
//		return err
//	}
//
//	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
//	req.Header.Set("Content-Type", "application/json")
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	if code := resp.StatusCode; code != 200 {
//		type tgResp struct {
//			OK          bool   `json:"ok"`
//			ErrorCode   int    `json:"error_code,omitempty"`
//			Description string `json:"description,omitempty"`
//		}
//
//		body, _ := ioutil.ReadAll(resp.Body)
//
//		var TgResp tgResp
//		err = json.Unmarshal(body, &TgResp)
//		if err != nil {
//			return err
//		}
//		return fmt.Errorf("%d | %s", code, TgResp.Description)
//	}
//
//	return nil
//}

func (bot *BOT) SendMessageWithFile(ctx context.Context, chatId int64, path string) error {
	file := tgbotapi.NewInputMediaDocument(tgbotapi.FilePath(path))

	msg := tgbotapi.NewMediaGroup(chatId, []interface{}{
		file,
	})

	_, err := bot.Bot.SendMediaGroup(msg)
	if err != nil {
		return fmt.Errorf("ошибка отправки сообщения: %s", err)
	}

	return nil
}

func (bot *BOT) SendMessageWithMediaPhoto(ctx context.Context, chatId int64, path string) error {
	file := tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(path))

	msg := tgbotapi.NewMediaGroup(chatId, []interface{}{
		file,
	})

	_, err := bot.Bot.SendMediaGroup(msg)
	if err != nil {
		return fmt.Errorf("ошибка отправки сообщения: %s", err)
	}

	return nil
}

func (bot *BOT) SendMessageWithText(ctx context.Context, chatId int64, text string) error {

	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "HTML"

	_, err := bot.Bot.Send(msg)
	if err != nil {
		return fmt.Errorf("ошибка отправки сообщения: %s", err)
	}

	return nil
}
