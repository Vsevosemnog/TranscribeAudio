package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BOT struct {
	Bot           *tgbotapi.BotAPI
	DefaultChatID int64
}

type SendingMessage struct {
	ChatID                   int64       `json:"chat_id"`
	Text                     string      `json:"text"`
	ParseMode                string      `json:"parse_mode,omitempty"`
	Entities                 []Entity    `json:"entities,omitempty"` //TODO
	DisableWebPagePreview    bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId         int         `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

type SendingPhoto struct {
	ChatID    int64  `json:"chat_id"`
	Photo     string `json:"photo"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode,omitempty"`
	//Entities                 []Entity    `json:"entities,omitempty"` //TODO
	//DisableWebPagePreview    bool        `json:"disable_web_page_preview,omitempty"`
	//DisableNotification      bool        `json:"disable_notification,omitempty"`
	//ReplyToMessageId         int         `json:"reply_to_message_id,omitempty"`
	//AllowSendingWithoutReply bool        `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

type SendingDocument struct {
	ChatID    int64  `json:"chat_id"`
	Document  string `json:"document"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string       `json:"text"`
	Url                          string       `json:"url"`
	LoginUrl                     LoginUrl     `json:"login_url,omitempty"`
	CallbackData                 string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool         `json:"pay,omitempty"`
}

type LoginUrl struct {
	//TODO
}

type CallbackGame struct {
	//TODO
}

type Entity struct {
	//TODO
}
