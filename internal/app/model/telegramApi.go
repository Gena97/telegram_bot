package model

import "time"

// Update - структура для получения всех типов обновлений
type Update struct {
	UpdateID                int                          `json:"update_id"`
	Message                 *Message                     `json:"message,omitempty"`
	EditedMessage           *Message                     `json:"edited_message,omitempty"`
	ChannelPost             *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage         *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	Poll                    *Poll                        `json:"poll,omitempty"`
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
}

type Message struct {
	MessageID                    int                           `json:"message_id"`
	MessageThreadID              int                           `json:"message_thread_id,omitempty"`
	From                         *User                         `json:"from,omitempty"`
	SenderChat                   *Chat                         `json:"sender_chat,omitempty"`
	SenderBoostCount             int                           `json:"sender_boost_count,omitempty"`
	SenderBusinessBot            *User                         `json:"sender_business_bot,omitempty"`
	BusinessConnectionID         string                        `json:"business_connection_id,omitempty"`
	Date                         int                           `json:"date"`
	Chat                         Chat                          `json:"chat"`
	ForwardOrigin                *MessageOrigin                `json:"forward_origin,omitempty"`
	IsTopicMessage               bool                          `json:"is_topic_message,omitempty"`
	IsAutomaticForward           bool                          `json:"is_automatic_forward,omitempty"`
	ReplyToMessage               *Message                      `json:"reply_to_message,omitempty"`
	ExternalReply                *ExternalReplyInfo            `json:"external_reply,omitempty"`
	Quote                        *TextQuote                    `json:"quote,omitempty"`
	ReplyToStory                 *Story                        `json:"reply_to_story,omitempty"`
	ViaBot                       *User                         `json:"via_bot,omitempty"`
	EditDate                     int                           `json:"edit_date,omitempty"`
	HasProtectedContent          bool                          `json:"has_protected_content,omitempty"`
	IsFromOffline                bool                          `json:"is_from_offline,omitempty"`
	MediaGroupID                 string                        `json:"media_group_id,omitempty"`
	AuthorSignature              string                        `json:"author_signature,omitempty"`
	Text                         string                        `json:"text,omitempty"`
	Entities                     []MessageEntity               `json:"entities,omitempty"`
	LinkPreviewOptions           *LinkPreviewOptions           `json:"link_preview_options,omitempty"`
	EffectID                     string                        `json:"effect_id,omitempty"`
	Animation                    *Animation                    `json:"animation,omitempty"`
	Audio                        *Audio                        `json:"audio,omitempty"`
	Document                     *Document                     `json:"document,omitempty"`
	PaidMedia                    *PaidMediaInfo                `json:"paid_media,omitempty"`
	Photo                        []PhotoSize                   `json:"photo,omitempty"`
	Sticker                      *Sticker                      `json:"sticker,omitempty"`
	Story                        *Story                        `json:"story,omitempty"`
	Video                        *Video                        `json:"video,omitempty"`
	VideoNote                    *VideoNote                    `json:"video_note,omitempty"`
	Voice                        *Voice                        `json:"voice,omitempty"`
	Caption                      string                        `json:"caption,omitempty"`
	CaptionEntities              []MessageEntity               `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia        bool                          `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler              bool                          `json:"has_media_spoiler,omitempty"`
	Contact                      *Contact                      `json:"contact,omitempty"`
	Dice                         *Dice                         `json:"dice,omitempty"`
	Game                         *Game                         `json:"game,omitempty"`
	Poll                         *Poll                         `json:"poll,omitempty"`
	Venue                        *Venue                        `json:"venue,omitempty"`
	Location                     *Location                     `json:"location,omitempty"`
	NewChatMembers               []User                        `json:"new_chat_members,omitempty"`
	LeftChatMember               *User                         `json:"left_chat_member,omitempty"`
	NewChatTitle                 string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                 []PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto              bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated             bool                          `json:"group_chat_created,omitempty"`
	SupergroupChatCreated        bool                          `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated           bool                          `json:"channel_chat_created,omitempty"`
	MigrateToChatID              int64                         `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID            int64                         `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                *Message                      `json:"pinned_message,omitempty"`
	Invoice                      *Invoice                      `json:"invoice,omitempty"`
	SuccessfulPayment            *SuccessfulPayment            `json:"successful_payment,omitempty"`
	RefundedPayment              *RefundedPayment              `json:"refunded_payment,omitempty"`
	UsersShared                  *UsersShared                  `json:"users_shared,omitempty"`
	ChatShared                   *ChatShared                   `json:"chat_shared,omitempty"`
	ConnectedWebsite             string                        `json:"connected_website,omitempty"`
	WriteAccessAllowed           *WriteAccessAllowed           `json:"write_access_allowed,omitempty"`
	PassportData                 *PassportData                 `json:"passport_data,omitempty"`
	ProximityAlertTriggered      *ProximityAlertTriggered      `json:"proximity_alert_triggered,omitempty"`
	BoostAdded                   *ChatBoostAdded               `json:"boost_added,omitempty"`
	ChatBackgroundSet            *ChatBackground               `json:"chat_background_set,omitempty"`
	ForumTopicCreated            *ForumTopicCreated            `json:"forum_topic_created,omitempty"`
	ForumTopicEdited             *ForumTopicEdited             `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed             *ForumTopicClosed             `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened           *ForumTopicReopened           `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden      *GeneralForumTopicHidden      `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden    *GeneralForumTopicUnhidden    `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated              *GiveawayCreated              `json:"giveaway_created,omitempty"`
	Giveaway                     *Giveaway                     `json:"giveaway,omitempty"`
	GiveawayWinners              *GiveawayWinners              `json:"giveaway_winners,omitempty"`
	GiveawayCompleted            *GiveawayCompleted            `json:"giveaway_completed,omitempty"`
	VideoChatScheduled           *VideoChatScheduled           `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted             *VideoChatStarted             `json:"video_chat_started,omitempty"`
	VideoChatEnded               *VideoChatEnded               `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	WebAppData                   *WebAppData                   `json:"web_app_data,omitempty"`
	ReplyMarkup                  *InlineKeyboardMarkup         `json:"reply_markup,omitempty"`
}

type User struct {
	ID                      int64  `json:"id"`                                    // Уникальный идентификатор пользователя или бота
	IsBot                   bool   `json:"is_bot"`                                // Истинно, если это бот
	FirstName               string `json:"first_name"`                            // Имя пользователя или бота
	LastName                string `json:"last_name,omitempty"`                   // Фамилия пользователя или бота (опционально)
	Username                string `json:"username,omitempty"`                    // Имя пользователя или бота (опционально)
	LanguageCode            string `json:"language_code,omitempty"`               // Языковой код пользователя (опционально)
	IsPremium               bool   `json:"is_premium,omitempty"`                  // Истинно, если это Telegram Premium пользователь (опционально)
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`    // Истинно, если пользователь добавил бота в меню вложений (опционально)
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`             // Истинно, если бот может быть приглашен в группы (опционально)
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"` // Истинно, если у бота отключен режим конфиденциальности (опционально)
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`     // Истинно, если бот поддерживает inline-запросы (опционально)
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`     // Истинно, если бот может быть подключен к Telegram Business аккаунту (опционально)
	HasMainWebApp           bool   `json:"has_main_web_app,omitempty"`            // Истинно, если у бота есть основное Web App приложение (опционально)
}

// BusinessConnection описывает соединение бота с бизнес-аккаунтом.
type BusinessConnection struct {
	ID         string    `json:"id"`                   // Уникальный идентификатор соединения
	User       User      `json:"user"`                 // Бизнес-аккаунт, который создал соединение
	UserChatID int64     `json:"user_chat_id"`         // Идентификатор чата с пользователем, создавшим соединение
	Date       time.Time `json:"date"`                 // Дата установления соединения
	CanReply   bool      `json:"can_reply,omitempty"`  // Может ли бот действовать от имени бизнес-аккаунта
	IsEnabled  bool      `json:"is_enabled,omitempty"` // Активно ли соединение
	// Если есть дополнительные необязательные поля
	RemovedChatBoost bool `json:"removed_chat_boost,omitempty"` // Пример необязательного поля
}

// BusinessMessagesDeleted представляет объект, получаемый при удалении сообщений из подключенной бизнес-учетной записи.
type BusinessMessagesDeleted struct {
	BusinessConnectionID string  `json:"business_connection_id"` // Уникальный идентификатор бизнес-соединения
	Chat                 Chat    `json:"chat"`                   // Информация о чате в бизнес-учетной записи
	MessageIDs           []int64 `json:"message_ids"`            // Список идентификаторов удаленных сообщений в чате бизнес-учетной записи
}

// MessageReactionUpdated представляет изменение реакции на сообщение, выполненное пользователем.
type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`                   // Чат, содержащий сообщение, на которое пользователь отреагировал
	MessageID   int64          `json:"message_id"`             // Уникальный идентификатор сообщения в чате
	User        *User          `json:"user,omitempty"`         // Пользователь, изменивший реакцию (опционально, если пользователь не анонимный)
	ActorChat   *Chat          `json:"actor_chat,omitempty"`   // Чат от имени которого была изменена реакция (опционально, если пользователь анонимный)
	Date        int64          `json:"date"`                   // Дата изменения в Unix time
	OldReaction []ReactionType `json:"old_reaction,omitempty"` // Предыдущий список типов реакций, установленных пользователем (опционально)
	NewReaction []ReactionType `json:"new_reaction,omitempty"` // Новый список типов реакций, установленных пользователем (опционально)
}

// MessageReactionCountUpdated представляет изменения реакций на сообщение с анонимными реакциями.
type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`       // Чат, содержащий сообщение
	MessageID int64           `json:"message_id"` // Уникальный идентификатор сообщения в чате
	Date      int64           `json:"date"`       // Дата изменения в Unix time
	Reactions []ReactionCount `json:"reactions"`  // Список реакций, присутствующих на сообщении
}

// InlineQuery представляет входящий инлайн-запрос.
type InlineQuery struct {
	ID       string    `json:"id"`                  // Уникальный идентификатор для этого запроса
	From     User      `json:"from"`                // Отправитель
	Query    string    `json:"query"`               // Текст запроса (до 256 символов)
	Offset   string    `json:"offset"`              // Смещение результатов, которые будут возвращены, может контролироваться ботом
	ChatType string    `json:"chat_type,omitempty"` // Тип чата, из которого был отправлен инлайн-запрос (опционально)
	Location *Location `json:"location,omitempty"`  // Местоположение отправителя (опционально)
}

// InlineQueryResultsButton представляет кнопку, которая будет показана над результатами инлайн-запроса.
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`                      // Текст на кнопке
	WebApp         *WebAppInfo `json:"web_app,omitempty"`         // Описание веб-приложения, которое будет запущено при нажатии кнопки (опционально)
	StartParameter string      `json:"start_parameter,omitempty"` // Параметр глубокого связывания для /start сообщения, отправленного боту при нажатии кнопки (опционально)
}

// AnswerInlineQueryParams представляет параметры метода answerInlineQuery.
type AnswerInlineQueryParams struct {
	InlineQueryID string                    `json:"inline_query_id"`       // Уникальный идентификатор для ответа на запрос
	Results       []InlineQueryResult       `json:"results"`               // Массив результатов для инлайн-запроса
	CacheTime     *int                      `json:"cache_time,omitempty"`  // Максимальное время в секундах, в течение которого результат может кэшироваться на сервере (опционально)
	IsPersonal    *bool                     `json:"is_personal,omitempty"` // true, если результаты могут кэшироваться только для пользователя, отправившего запрос (опционально)
	NextOffset    *string                   `json:"next_offset,omitempty"` // Смещение, которое клиент должен отправить в следующем запросе (опционально)
	Button        *InlineQueryResultsButton `json:"button,omitempty"`      // Кнопка, показываемая над инлайн-результатами (опционально)
}

type InlineQueryResult struct {
	Type                string                `json:"type"`                            // Type of the result (e.g., venue, contact, game, etc.)
	ID                  string                `json:"id"`                              // Unique identifier for this result, 1-64 bytes
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`          // Inline keyboard attached to the message
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"` // Content of the message to be sent instead of the result

	// Common fields for different result types
	Title                 string          `json:"title,omitempty"`                    // Title for the result (if applicable)
	Description           string          `json:"description,omitempty"`              // Short description of the result (if applicable)
	Caption               string          `json:"caption,omitempty"`                  // Caption to be sent (if applicable)
	ParseMode             string          `json:"parse_mode,omitempty"`               // Mode for parsing entities in the caption (if applicable)
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`         // Special entities in the caption (if applicable)
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"` // If the caption must be shown above the message media (if applicable)
	ThumbnailURL          string          `json:"thumbnail_url,omitempty"`            // URL of the thumbnail for the result (if applicable)
	ThumbnailWidth        int             `json:"thumbnail_width,omitempty"`          // Thumbnail width (if applicable)
	ThumbnailHeight       int             `json:"thumbnail_height,omitempty"`         // Thumbnail height (if applicable)

	// Specific fields for certain types
	Latitude       float64 `json:"latitude,omitempty"`         // Latitude of the venue (if applicable)
	Longitude      float64 `json:"longitude,omitempty"`        // Longitude of the venue (if applicable)
	PhoneNumber    string  `json:"phone_number,omitempty"`     // Contact's phone number (if applicable)
	LastName       string  `json:"last_name,omitempty"`        // Contact's last name (if applicable)
	VCard          string  `json:"vcard,omitempty"`            // Additional data about the contact (if applicable)
	GameShortName  string  `json:"game_short_name,omitempty"`  // Short name of the game (if applicable)
	PhotoFileID    string  `json:"photo_file_id,omitempty"`    // Valid file identifier of the photo (if applicable)
	GIFFileID      string  `json:"gif_file_id,omitempty"`      // Valid file identifier for the GIF file (if applicable)
	Mpeg4FileID    string  `json:"mpeg4_file_id,omitempty"`    // Valid file identifier for the MPEG4 file (if applicable)
	StickerFileID  string  `json:"sticker_file_id,omitempty"`  // Valid file identifier of the sticker (if applicable)
	DocumentFileID string  `json:"document_file_id,omitempty"` // Valid file identifier for the file (if applicable)
	VideoFileID    string  `json:"video_file_id,omitempty"`    // Valid file identifier for the video file (if applicable)
	VoiceFileID    string  `json:"voice_file_id,omitempty"`    // Valid file identifier for the voice message (if applicable)
	AudioFileID    string  `json:"audio_file_id,omitempty"`    // Valid file identifier for the audio file (if applicable)

	FoursquareID    string `json:"foursquare_id,omitempty"`     // Foursquare identifier of the venue (if applicable)
	FoursquareType  string `json:"foursquare_type,omitempty"`   // Foursquare type of the venue (if applicable)
	GooglePlaceID   string `json:"google_place_id,omitempty"`   // Google Places identifier of the venue (if applicable)
	GooglePlaceType string `json:"google_place_type,omitempty"` // Google Places type of the venue (if applicable)
}
type InputMessageContent struct {
	// Common fields for all message content types
	ParseMode string          `json:"parse_mode,omitempty"` // Optional. Mode for parsing entities in the message text
	Entities  []MessageEntity `json:"entities,omitempty"`   // Optional. List of special entities that appear in message text

	// InputTextMessageContent
	Text               string              `json:"message_text,omitempty"`         // Text of the message to be sent, 1-4096 characters
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Link preview generation options

	// InputLocationMessageContent
	Latitude             float64 `json:"latitude,omitempty"`               // Latitude of the location in degrees
	Longitude            float64 `json:"longitude,omitempty"`              // Longitude of the location in degrees
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`    // Optional. Radius of uncertainty for the location in meters
	LivePeriod           int     `json:"live_period,omitempty"`            // Optional. Period in seconds for live updates
	Heading              int     `json:"heading,omitempty"`                // Optional. Direction for live locations
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"` // Optional. Max distance for proximity alerts

	// InputVenueMessageContent
	VenueTitle      string `json:"title,omitempty"`             // Name of the venue
	VenueAddress    string `json:"address,omitempty"`           // Address of the venue
	FoursquareID    string `json:"foursquare_id,omitempty"`     // Optional. Foursquare ID of the venue
	FoursquareType  string `json:"foursquare_type,omitempty"`   // Optional. Foursquare type of the venue
	GooglePlaceID   string `json:"google_place_id,omitempty"`   // Optional. Google Places ID of the venue
	GooglePlaceType string `json:"google_place_type,omitempty"` // Optional. Google Places type of the venue

	// InputContactMessageContent
	PhoneNumber string `json:"phone_number,omitempty"` // Contact's phone number
	FirstName   string `json:"first_name,omitempty"`   // Contact's first name
	LastName    string `json:"last_name,omitempty"`    // Optional. Contact's last name
	VCard       string `json:"vcard,omitempty"`        // Optional. Additional data about the contact in vCard format

	// InputInvoiceMessageContent
	InvoiceTitle              string         `json:"title,omitempty"`                         // Product name, 1-32 characters
	InvoiceDescription        string         `json:"description,omitempty"`                   // Product description, 1-255 characters
	Payload                   string         `json:"payload,omitempty"`                       // Bot-defined invoice payload
	ProviderToken             string         `json:"provider_token,omitempty"`                // Optional. Payment provider token
	Currency                  string         `json:"currency,omitempty"`                      // Three-letter ISO 4217 currency code
	Prices                    []LabeledPrice `json:"prices,omitempty"`                        // Price breakdown
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`                // Optional. Maximum accepted tip amount
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`         // Optional. Suggested tip amounts
	ProviderData              string         `json:"provider_data,omitempty"`                 // Optional. Data shared with the payment provider
	PhotoURL                  string         `json:"photo_url,omitempty"`                     // Optional. URL of the product photo
	PhotoSize                 int            `json:"photo_size,omitempty"`                    // Optional. Photo size in bytes
	PhotoWidth                int            `json:"photo_width,omitempty"`                   // Optional. Photo width
	PhotoHeight               int            `json:"photo_height,omitempty"`                  // Optional. Photo height
	NeedName                  bool           `json:"need_name,omitempty"`                     // Optional. Requires user's full name
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`             // Optional. Requires user's phone number
	NeedEmail                 bool           `json:"need_email,omitempty"`                    // Optional. Requires user's email address
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`         // Optional. Requires user's shipping address
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"` // Optional. Sends user's phone number to the provider
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`        // Optional. Sends user's email address to the provider
	IsFlexible                bool           `json:"is_flexible,omitempty"`                   // Optional. Final price depends on the shipping method
}

// LabeledPrice represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label  string `json:"label"`  // Portion label
	Amount int    `json:"amount"` // Price of the product in the smallest units of the currency (integer)
}

// ChosenInlineResult представляет результат инлайн-запроса, выбранный пользователем и отправленный их партнеру по чату.
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`                   // Уникальный идентификатор для выбранного результата
	From            User      `json:"from"`                        // Пользователь, который выбрал результат
	Location        *Location `json:"location,omitempty"`          // Местоположение отправителя (опционально)
	InlineMessageID *string   `json:"inline_message_id,omitempty"` // Идентификатор отправленного инлайн-сообщения (опционально)
	Query           string    `json:"query"`                       // Запрос, использованный для получения результата
}

// CallbackQuery представляет входящий запрос обратного вызова от кнопки обратного вызова в инлайн-клавиатуре.
type CallbackQuery struct {
	ID              string                    `json:"id"`                          // Уникальный идентификатор для этого запроса
	From            User                      `json:"from"`                        // Отправитель
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`           // Сообщение, отправленное ботом с кнопкой обратного вызова (опционально)
	InlineMessageID *string                   `json:"inline_message_id,omitempty"` // Идентификатор сообщения, отправленного через бота в инлайн-режиме (опционально)
	ChatInstance    string                    `json:"chat_instance"`               // Глобальный идентификатор, уникально соответствующий чату
	Data            *string                   `json:"data,omitempty"`              // Данные, связанные с кнопкой обратного вызова (опционально)
	GameShortName   *string                   `json:"game_short_name,omitempty"`   // Короткое имя игры для возврата (опционально)
}

// ShippingQuery содержит информацию о входящем запросе на доставку.
type ShippingQuery struct {
	ID              string          `json:"id"`               // Уникальный идентификатор запроса
	From            User            `json:"from"`             // Пользователь, отправивший запрос
	InvoicePayload  string          `json:"invoice_payload"`  // Нагрузочный пакет счета, указанный ботом
	ShippingAddress ShippingAddress `json:"shipping_address"` // Указанный пользователем адрес доставки
}

// PreCheckoutQuery содержит информацию о входящем запросе на предоплату.
type PreCheckoutQuery struct {
	ID               string     `json:"id"`                           // Уникальный идентификатор запроса
	From             User       `json:"from"`                         // Пользователь, отправивший запрос
	Currency         string     `json:"currency"`                     // Трехбуквенный код валюты ISO 4217
	TotalAmount      int        `json:"total_amount"`                 // Общая цена в наименьших единицах валюты
	InvoicePayload   string     `json:"invoice_payload"`              // Нагрузочный пакет счета, указанный ботом
	ShippingOptionID *string    `json:"shipping_option_id,omitempty"` // Идентификатор выбранного пользователем варианта доставки (опционально)
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`         // Информация о заказе, предоставленная пользователем (опционально)
}

// PaidMediaPurchased содержит информацию о покупке оплаченного медиа.
type PaidMediaPurchased struct {
	From             User   `json:"from"`               // Пользователь, который приобрел медиа
	PaidMediaPayload string `json:"paid_media_payload"` // Нагрузочный пакет оплаченного медиа, указанный ботом
}

// Poll содержит информацию о голосовании.
type Poll struct {
	ID                    string          `json:"id"`                             // Уникальный идентификатор голосования
	Question              string          `json:"question"`                       // Вопрос голосования (1-300 символов)
	QuestionEntities      []MessageEntity `json:"question_entities,omitempty"`    // Специальные сущности, которые появляются в вопросе (опционально)
	Options               []PollOption    `json:"options"`                        // Список вариантов голосования
	TotalVoterCount       int             `json:"total_voter_count"`              // Общее количество пользователей, проголосовавших в голосовании
	IsClosed              bool            `json:"is_closed"`                      // true, если голосование закрыто
	IsAnonymous           bool            `json:"is_anonymous"`                   // true, если голосование анонимно
	Type                  string          `json:"type"`                           // Тип голосования (например, "обычное" или "викторина")
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`        // true, если голосование позволяет несколько ответов
	CorrectOptionID       *int            `json:"correct_option_id,omitempty"`    // 0-индексированный идентификатор правильного ответа (опционально)
	Explanation           *string         `json:"explanation,omitempty"`          // Текст, показываемый при выборе неправильного ответа (опционально)
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"` // Специальные сущности в объяснении (опционально)
	OpenPeriod            *int            `json:"open_period,omitempty"`          // Время в секундах, в течение которого голосование будет активно (опционально)
	CloseDate             *int            `json:"close_date,omitempty"`           // Время закрытия голосования (Unix timestamp) (опционально)
}

// PollAnswer представляет ответ пользователя в неанонимном голосовании.
type PollAnswer struct {
	PollID    string `json:"poll_id"`              // Уникальный идентификатор голосования
	VoterChat *Chat  `json:"voter_chat,omitempty"` // Чат, в котором изменился ответ на голосование (опционально)
	User      *User  `json:"user,omitempty"`       // Пользователь, изменивший ответ на голосование (опционально)
	OptionIDs []int  `json:"option_ids"`           // 0-индексированные идентификаторы выбранных вариантов ответов
}

// ChatMemberUpdated представляет изменения в статусе участника чата.
type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`                        // Чат, к которому принадлежит пользователь
	From                    User            `json:"from"`                        // Исполнитель действия, приведшего к изменению
	Date                    int             `json:"date"`                        // Дата изменения (Unix time)
	OldChatMember           ChatMember      `json:"old_chat_member"`             // Предыдущая информация о члене чата
	NewChatMember           ChatMember      `json:"new_chat_member"`             // Новая информация о члене чата
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`       // Ссылка для приглашения, использованная для присоединения (опционально)
	ViaJoinRequest          bool            `json:"via_join_request"`            // true, если пользователь присоединился через прямой запрос на присоединение
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link"` // true, если пользователь присоединился через ссылку для приглашения в папке чатов
}

// ChatInviteLink represents an invite link for a chat.
type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`                          // The invite link.
	Creator                 User    `json:"creator"`                              // Creator of the link.
	CreatesJoinRequest      bool    `json:"creates_join_request"`                 // True if users need approval to join via the link.
	IsPrimary               bool    `json:"is_primary"`                           // True if the link is primary.
	IsRevoked               bool    `json:"is_revoked"`                           // True if the link is revoked.
	Name                    *string `json:"name,omitempty"`                       // Optional. Invite link name.
	ExpireDate              *int    `json:"expire_date,omitempty"`                // Optional. Expiration date (Unix timestamp).
	MemberLimit             *int    `json:"member_limit,omitempty"`               // Optional. Max number of members allowed to join via this link (1-99999).
	PendingJoinRequestCount *int    `json:"pending_join_request_count,omitempty"` // Optional. Number of pending join requests.
	SubscriptionPeriod      *int    `json:"subscription_period,omitempty"`        // Optional. Duration of the subscription in seconds.
	SubscriptionPrice       *int    `json:"subscription_price,omitempty"`         // Optional. Price in Telegram Stars for subscription.
}
type PaidMedia struct {
	Type     string      `json:"type"`               // Type of the paid media: "preview", "photo", or "video".
	Width    *int        `json:"width,omitempty"`    // Optional. Media width as defined by the sender (only for preview).
	Height   *int        `json:"height,omitempty"`   // Optional. Media height as defined by the sender (only for preview).
	Duration *int        `json:"duration,omitempty"` // Optional. Duration of the media in seconds (only for preview).
	Photo    []PhotoSize `json:"photo,omitempty"`    // Optional. Photo array (only for photo).
	Video    *Video      `json:"video,omitempty"`    // Optional. Video object (only for video).
}

// ChatJoinRequest представляет запрос на присоединение к чату.
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`                  // Чат, в который был отправлен запрос
	From       User            `json:"from"`                  // Пользователь, отправивший запрос
	UserChatID int64           `json:"user_chat_id"`          // Идентификатор личного чата с пользователем, отправившим запрос
	Date       int             `json:"date"`                  // Дата отправки запроса (Unix time)
	Bio        *string         `json:"bio,omitempty"`         // Биография пользователя (опционально)
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"` // Ссылка для приглашения, использованная для отправки запроса (опционально)
}

// ChatBoostUpdated представляет увеличение, добавленное в чат или измененное.
type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`  // Чат, который был увеличен
	Boost ChatBoost `json:"boost"` // Информация о повышении чата
}

// ChatBoostRemoved представляет увеличение, удаленное из чата.
type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`        // Чат, который был увеличен
	BoostID    string          `json:"boost_id"`    // Уникальный идентификатор повышения
	RemoveDate int             `json:"remove_date"` // Время удаления (Unix timestamp)
	Source     ChatBoostSource `json:"source"`      // Источник удаленного повышения
}

// Chat представляет объект чата.
type Chat struct {
	ID        int64   `json:"id"`                   // Уникальный идентификатор для этого чата
	Type      string  `json:"type"`                 // Тип чата ("private", "group", "supergroup" или "channel")
	Title     *string `json:"title,omitempty"`      // Заголовок (опционально, для супергрупп, каналов и групповых чатов)
	Username  *string `json:"username,omitempty"`   // Имя пользователя (опционально, для частных чатов, супергрупп и каналов)
	FirstName *string `json:"first_name,omitempty"` // Имя другой стороны в частном чате (опционально)
	LastName  *string `json:"last_name,omitempty"`  // Фамилия другой стороны в частном чате (опционально)
	IsForum   *bool   `json:"is_forum,omitempty"`   // true, если супергруппа является форумом (опционально)
}

// ExternalReplyInfo содержит информацию о сообщении, на которое идет ответ.
type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`                         // Происхождение сообщения, на которое отвечают
	Chat               *Chat               `json:"chat,omitempty"`                 // Чат, к которому принадлежит оригинальное сообщение (опционально)
	MessageID          *int                `json:"message_id,omitempty"`           // Уникальный идентификатор сообщения в оригинальном чате (опционально)
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Опции для генерации превью ссылки
	Animation          *Animation          `json:"animation,omitempty"`            // Информация о сообщении-анимации (опционально)
	Audio              *Audio              `json:"audio,omitempty"`                // Информация о звуковом файле (опционально)
	Document           *Document           `json:"document,omitempty"`             // Информация о общем файле (опционально)
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`           // Информация о платном медиа (опционально)
	Photo              []PhotoSize         `json:"photo,omitempty"`                // Доступные размеры фотографии (опционально)
	Sticker            *Sticker            `json:"sticker,omitempty"`              // Информация о стикере (опционально)
	Story              *Story              `json:"story,omitempty"`                // Пересланная история (опционально)
	Video              *Video              `json:"video,omitempty"`                // Информация о видео (опционально)
	VideoNote          *VideoNote          `json:"video_note,omitempty"`           // Информация о видеосообщении (опционально)
	Voice              *Voice              `json:"voice,omitempty"`                // Информация о голосовом сообщении (опционально)
	HasMediaSpoiler    *bool               `json:"has_media_spoiler,omitempty"`    // True, если медиа сообщения покрыто анимацией спойлера (опционально)
	Contact            *Contact            `json:"contact,omitempty"`              // Информация о переданном контакте (опционально)
	Dice               *Dice               `json:"dice,omitempty"`                 // Информация о кубике с случайным значением (опционально)
	Game               *Game               `json:"game,omitempty"`                 // Информация о игре (опционально)
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`             // Информация о запланированном розыгрыше (опционально)
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`     // Завершенный розыгрыш с публичными победителями (опционально)
	Invoice            *Invoice            `json:"invoice,omitempty"`              // Информация о счете для оплаты (опционально)
	Location           *Location           `json:"location,omitempty"`             // Информация о переданной локации (опционально)
	Poll               *Poll               `json:"poll,omitempty"`                 // Информация о нативном опросе (опционально)
	Venue              *Venue              `json:"venue,omitempty"`                // Информация о месте (опционально)
}

// TextQuote содержит информацию о цитируемой части сообщения.
type TextQuote struct {
	Text     string          `json:"text"`                // Текст цитируемой части сообщения
	Entities []MessageEntity `json:"entities,omitempty"`  // Специальные сущности, которые появляются в цитате
	Position int             `json:"position"`            // Приблизительная позиция цитаты в оригинальном сообщении
	IsManual *bool           `json:"is_manual,omitempty"` // True, если цитата выбрана вручную
}

// Story представляет объект истории.
type Story struct {
	Chat Chat `json:"chat"` // Чат, который опубликовал историю
	ID   int  `json:"id"`   // Уникальный идентификатор для истории в чате
}

// MessageOrigin describes the origin of a message.
type MessageOrigin struct {
	Type       string                   `json:"type"`                  // Type of the message origin
	User       *MessageOriginUser       `json:"user,omitempty"`        // Origin from a known user
	HiddenUser *MessageOriginHiddenUser `json:"hidden_user,omitempty"` // Origin from an unknown user
	Chat       *MessageOriginChat       `json:"chat,omitempty"`        // Origin from a chat
	Channel    *MessageOriginChannel    `json:"channel,omitempty"`     // Origin from a channel
}

// MessageOriginUser represents a message originally sent by a known user.
type MessageOriginUser struct {
	Type       string `json:"type"`        // Type of the message origin, always "user"
	Date       int    `json:"date"`        // Date the message was sent originally in Unix time
	SenderUser User   `json:"sender_user"` // User that sent the message originally
}

// MessageOriginHiddenUser represents a message originally sent by an unknown user.
type MessageOriginHiddenUser struct {
	Type           string `json:"type"`             // Type of the message origin, always "hidden_user"
	Date           int    `json:"date"`             // Date the message was sent originally in Unix time
	SenderUserName string `json:"sender_user_name"` // Name of the user that sent the message originally
}

// MessageOriginChat represents a message originally sent on behalf of a chat.
type MessageOriginChat struct {
	Type            string  `json:"type"`                       // Type of the message origin, always "chat"
	Date            int     `json:"date"`                       // Date the message was sent originally in Unix time
	SenderChat      Chat    `json:"sender_chat"`                // Chat that sent the message originally
	AuthorSignature *string `json:"author_signature,omitempty"` // Signature of the original message author (optional)
}

// MessageOriginChannel represents a message originally sent to a channel chat.
type MessageOriginChannel struct {
	Type            string  `json:"type"`                       // Type of the message origin, always "channel"
	Date            int     `json:"date"`                       // Date the message was sent originally in Unix time
	Chat            Chat    `json:"chat"`                       // Channel chat to which the message was originally sent
	MessageID       int     `json:"message_id"`                 // Unique message identifier inside the chat
	AuthorSignature *string `json:"author_signature,omitempty"` // Signature of the original post author (optional)
}

// MessageEntity represents one special entity in a text message.
type MessageEntity struct {
	Type          string  `json:"type"`                      // Type of the entity (e.g., mention, hashtag, URL, etc.)
	Offset        int     `json:"offset"`                    // Offset in UTF-16 code units to the start of the entity
	Length        int     `json:"length"`                    // Length of the entity in UTF-16 code units
	URL           *string `json:"url,omitempty"`             // Optional URL for "text_link"
	User          *User   `json:"user,omitempty"`            // Optional user for "text_mention"
	Language      *string `json:"language,omitempty"`        // Optional programming language for "pre"
	CustomEmojiID *string `json:"custom_emoji_id,omitempty"` // Optional custom emoji identifier
}

// LinkPreviewOptions describes the options used for link preview generation.
type LinkPreviewOptions struct {
	IsDisabled       *bool   `json:"is_disabled,omitempty"`        // Optional, true if the link preview is disabled
	URL              *string `json:"url,omitempty"`                // Optional URL for the link preview
	PreferSmallMedia *bool   `json:"prefer_small_media,omitempty"` // Optional, true if small media is preferred
	PreferLargeMedia *bool   `json:"prefer_large_media,omitempty"` // Optional, true if large media is preferred
	ShowAboveText    *bool   `json:"show_above_text,omitempty"`    // Optional, true if the preview is shown above the text
}

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileID       string     `json:"file_id"`             // Identifier for this file
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier for this file
	Width        int        `json:"width"`               // Video width as defined by the sender
	Height       int        `json:"height"`              // Video height as defined by the sender
	Duration     int        `json:"duration"`            // Duration of the video in seconds
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional thumbnail for the animation
	FileName     *string    `json:"file_name,omitempty"` // Optional original filename
	MIMEType     *string    `json:"mime_type,omitempty"` // Optional MIME type of the file
	FileSize     *int64     `json:"file_size,omitempty"` // Optional file size in bytes
}

// Audio represents an audio file to be treated as music by Telegram clients.
type Audio struct {
	FileID       string     `json:"file_id"`             // Identifier for this file
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier for this file
	Duration     int        `json:"duration"`            // Duration of the audio in seconds
	Performer    *string    `json:"performer,omitempty"` // Optional performer of the audio
	Title        *string    `json:"title,omitempty"`     // Optional title of the audio
	FileName     *string    `json:"file_name,omitempty"` // Optional original filename
	MIMEType     *string    `json:"mime_type,omitempty"` // Optional MIME type of the file
	FileSize     *int64     `json:"file_size,omitempty"` // Optional file size in bytes
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional thumbnail of the album cover
}

// Document represents a general file.
type Document struct {
	FileID       string     `json:"file_id"`             // Identifier for this file
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier for this file
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional document thumbnail
	FileName     *string    `json:"file_name,omitempty"` // Optional original filename
	MIMEType     *string    `json:"mime_type,omitempty"` // Optional MIME type of the file
	FileSize     *int64     `json:"file_size,omitempty"` // Optional file size in bytes
}

// PaidMediaInfo describes the paid media added to a message.
type PaidMediaInfo struct {
	StarCount int         `json:"star_count"` // Number of Telegram Stars needed to access the media
	PaidMedia []PaidMedia `json:"paid_media"` // Information about the paid media
}

// PhotoSize represents one size of a photo or a file/sticker thumbnail.
type PhotoSize struct {
	FileID       string `json:"file_id"`             // Identifier for this file
	FileUniqueID string `json:"file_unique_id"`      // Unique identifier for this file
	Width        int    `json:"width"`               // Photo width
	Height       int    `json:"height"`              // Photo height
	FileSize     *int64 `json:"file_size,omitempty"` // Optional file size in bytes
}

// Sticker represents a sticker.
type Sticker struct {
	FileID           string        `json:"file_id"`                     // Identifier for this file
	FileUniqueID     string        `json:"file_unique_id"`              // Unique identifier for this file
	Type             string        `json:"type"`                        // Type of the sticker
	Width            int           `json:"width"`                       // Sticker width
	Height           int           `json:"height"`                      // Sticker height
	IsAnimated       bool          `json:"is_animated"`                 // True if the sticker is animated
	IsVideo          bool          `json:"is_video"`                    // True if the sticker is a video sticker
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`         // Optional thumbnail of the sticker
	Emoji            *string       `json:"emoji,omitempty"`             // Optional associated emoji
	SetName          *string       `json:"set_name,omitempty"`          // Optional name of the sticker set
	PremiumAnimation *File         `json:"premium_animation,omitempty"` // For premium regular stickers
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`     // For mask stickers
	CustomEmojiID    *string       `json:"custom_emoji_id,omitempty"`   // For custom emoji stickers
	NeedsRepainting  *bool         `json:"needs_repainting,omitempty"`  // True if the sticker must be repainted
	FileSize         *int64        `json:"file_size,omitempty"`         // Optional file size in bytes
}

type MaskPosition struct {
	Point  string  `json:"point"`   // The part of the face relative to which the mask should be placed ("forehead", "eyes", "mouth", or "chin").
	XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size.
	YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size.
	Scale  float64 `json:"scale"`   // Mask scaling coefficient (e.g., 2.0 means double size).
}

// Video represents a video file.
type Video struct {
	FileID       string     `json:"file_id"`             // Identifier for this file
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier for this file
	Width        int        `json:"width"`               // Video width
	Height       int        `json:"height"`              // Video height
	Duration     int        `json:"duration"`            // Duration of the video in seconds
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional video thumbnail
	FileName     *string    `json:"file_name,omitempty"` // Optional original filename
	MIMEType     *string    `json:"mime_type,omitempty"` // Optional MIME type of the file
	FileSize     *int64     `json:"file_size,omitempty"` // Optional file size in bytes
}

// VideoNote represents a video message.
type VideoNote struct {
	FileID       string     `json:"file_id"`             // Identifier for this file
	FileUniqueID string     `json:"file_unique_id"`      // Unique identifier for this file
	Length       int        `json:"length"`              // Diameter of the video message
	Duration     int        `json:"duration"`            // Duration of the video in seconds
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional video thumbnail
	FileSize     *int64     `json:"file_size,omitempty"` // Optional file size in bytes
}

// Voice represents a voice note.
type Voice struct {
	FileID       string  `json:"file_id"`             // Identifier for this file
	FileUniqueID string  `json:"file_unique_id"`      // Unique identifier for this file
	Duration     int     `json:"duration"`            // Duration of the audio in seconds
	MIMEType     *string `json:"mime_type,omitempty"` // Optional MIME type of the file
	FileSize     *int64  `json:"file_size,omitempty"` // Optional file size in bytes
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string  `json:"phone_number"`        // Contact's phone number
	FirstName   string  `json:"first_name"`          // Contact's first name
	LastName    *string `json:"last_name,omitempty"` // Optional contact's last name
	UserID      *int64  `json:"user_id,omitempty"`   // Optional contact's user identifier in Telegram
	VCard       *string `json:"vcard,omitempty"`     // Optional additional data about the contact in vCard format
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"` // Emoji on which the dice throw animation is based
	Value int    `json:"value"` // Value of the dice (1-6 for certain emojis)
}

// Game represents a game.
type Game struct {
	Title        string           `json:"title"`                   // Title of the game
	Description  string           `json:"description"`             // Description of the game
	Photo        []PhotoSize      `json:"photo"`                   // Photo displayed in the game message
	Text         *string          `json:"text,omitempty"`          // Optional brief description of the game or high scores
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"` // Optional special entities in text
	Animation    *Animation       `json:"animation,omitempty"`     // Optional animation displayed in the game message
}

// Venue represents a venue.
type Venue struct {
	Location        Location `json:"location"`                    // Venue location; can't be a live location
	Title           string   `json:"title"`                       // Name of the venue
	Address         string   `json:"address"`                     // Address of the venue
	FoursquareID    *string  `json:"foursquare_id,omitempty"`     // Optional Foursquare identifier
	FoursquareType  *string  `json:"foursquare_type,omitempty"`   // Optional Foursquare type
	GooglePlaceID   *string  `json:"google_place_id,omitempty"`   // Optional Google Places identifier
	GooglePlaceType *string  `json:"google_place_type,omitempty"` // Optional Google Places type
}

// Location represents a point on the map.
type Location struct {
	Latitude             float64  `json:"latitude"`                         // Latitude as defined by the sender
	Longitude            float64  `json:"longitude"`                        // Longitude as defined by the sender
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`    // Optional uncertainty radius in meters
	LivePeriod           *int     `json:"live_period,omitempty"`            // Optional time for location updates in seconds
	Heading              *int     `json:"heading,omitempty"`                // Optional direction in degrees
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"` // Optional max distance for proximity alerts
}

// Invoice contains basic information about an invoice.
type Invoice struct {
	Title          string `json:"title"`           // Product name
	Description    string `json:"description"`     // Product description
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter
	Currency       string `json:"currency"`        // Currency code
	TotalAmount    int    `json:"total_amount"`    // Total price in the smallest currency units
}

// SuccessfulPayment contains information about a successful payment.
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`                     // Currency code
	TotalAmount             int        `json:"total_amount"`                 // Total price in the smallest units
	InvoicePayload          string     `json:"invoice_payload"`              // Invoice payload
	ShippingOptionID        *string    `json:"shipping_option_id,omitempty"` // Optional shipping option ID
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`         // Optional order information
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`   // Telegram payment identifier
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`   // Provider payment identifier
}

// UsersShared contains information about users whose identifiers were shared with the bot.
type UsersShared struct {
	RequestID int          `json:"request_id"` // Identifier of the request
	Users     []SharedUser `json:"users"`      // Information about shared users
}

// ChatShared contains information about a chat that was shared with the bot.
type ChatShared struct {
	RequestID int         `json:"request_id"`         // Identifier of the request
	ChatID    int64       `json:"chat_id"`            // Identifier of the shared chat
	Title     *string     `json:"title,omitempty"`    // Optional title of the chat
	Username  *string     `json:"username,omitempty"` // Optional username of the chat
	Photo     []PhotoSize `json:"photo,omitempty"`    // Optional photo of the chat
}

// RefundedPayment contains basic information about a refunded payment.
type RefundedPayment struct {
	Currency                string  `json:"currency"`                             // Currency code, always "XTR"
	TotalAmount             int     `json:"total_amount"`                         // Total refunded price in the smallest units
	InvoicePayload          string  `json:"invoice_payload"`                      // Invoice payload
	TelegramPaymentChargeID string  `json:"telegram_payment_charge_id"`           // Telegram payment identifier
	ProviderPaymentChargeID *string `json:"provider_payment_charge_id,omitempty"` // Optional provider payment identifier
}

// WriteAccessAllowed represents a service message about a user allowing a bot to write messages.
type WriteAccessAllowed struct {
	FromRequest        *bool   `json:"from_request,omitempty"`         // Optional. True if granted from Web App request
	WebAppName         *string `json:"web_app_name,omitempty"`         // Optional. Name of the Web App, if applicable
	FromAttachmentMenu *bool   `json:"from_attachment_menu,omitempty"` // Optional. True if granted from attachment menu
}

// PassportData describes Telegram Passport data shared with the bot.
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`        // Information about shared documents
	Credentials EncryptedCredentials       `json:"credentials"` // Encrypted credentials for decryption
}

// ProximityAlertTriggered represents a service message about a triggered proximity alert.
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"` // User that triggered the alert
	Watcher  User `json:"watcher"`  // User that set the alert
	Distance int  `json:"distance"` // Distance between the users
}

// ChatBoostAdded represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"` // Number of boosts added by the user
}

// ChatBackground represents a chat background.
type ChatBackground struct {
	Type BackgroundType `json:"type"` // Type of the background
}

// ForumTopicCreated represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	Name              string  `json:"name"`                           // Name of the topic
	IconColor         int     `json:"icon_color"`                     // Color of the topic icon in RGB
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"` // Optional custom emoji ID for the topic icon
}

// ForumTopicClosed represents a service message about a forum topic closed in the chat.
type ForumTopicClosed struct {
	// Currently holds no information
}

// ForumTopicEdited represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	Name              *string `json:"name,omitempty"`                 // Optional new name of the topic
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"` // Optional new emoji ID, empty if icon was removed
}

// ForumTopicReopened represents a service message about a reopened forum topic.
type ForumTopicReopened struct {
	// Currently holds no information
}

// GeneralForumTopicHidden represents a service message about a General forum topic hidden.
type GeneralForumTopicHidden struct {
	// Currently holds no information
}

// GeneralForumTopicUnhidden represents a service message about a General forum topic unhidden.
type GeneralForumTopicUnhidden struct {
	// Currently holds no information
}

// GiveawayCreated represents a service message about the creation of a scheduled giveaway.
type GiveawayCreated struct {
	PrizeStarCount *int `json:"prize_star_count,omitempty"` // Optional. Number of Telegram Stars for giveaway winners
}

// Giveaway represents a message about a scheduled giveaway.
type Giveaway struct {
	Chats                         []Chat   `json:"chats"`                                      // List of chats the user must join to participate
	WinnersSelectionDate          int      `json:"winners_selection_date"`                     // Unix timestamp for when winners will be selected
	WinnerCount                   int      `json:"winner_count"`                               // Number of winners to be selected
	OnlyNewMembers                *bool    `json:"only_new_members,omitempty"`                 // Optional. True if only new members can win
	HasPublicWinners              *bool    `json:"has_public_winners,omitempty"`               // Optional. True if winners list is public
	PrizeDescription              *string  `json:"prize_description,omitempty"`                // Optional. Description of additional prize
	CountryCodes                  []string `json:"country_codes,omitempty"`                    // Optional. List of eligible countries
	PrizeStarCount                *int     `json:"prize_star_count,omitempty"`                 // Optional. Stars for Telegram Star giveaways
	PremiumSubscriptionMonthCount *int     `json:"premium_subscription_month_count,omitempty"` // Optional. Months for Telegram Premium
}

// GiveawayWinners represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	Chat                          Chat    `json:"chat"`                                       // The chat that created the giveaway
	GiveawayMessageID             int     `json:"giveaway_message_id"`                        // Identifier of the giveaway message
	WinnersSelectionDate          int     `json:"winners_selection_date"`                     // Unix timestamp for when winners were selected
	WinnerCount                   int     `json:"winner_count"`                               // Total number of winners
	Winners                       []User  `json:"winners"`                                    // List of up to 100 winners
	AdditionalChatCount           *int    `json:"additional_chat_count,omitempty"`            // Optional. Other chats to join
	PrizeStarCount                *int    `json:"prize_star_count,omitempty"`                 // Optional. Stars for Telegram Star giveaways
	PremiumSubscriptionMonthCount *int    `json:"premium_subscription_month_count,omitempty"` // Optional. Months for Premium
	UnclaimedPrizeCount           *int    `json:"unclaimed_prize_count,omitempty"`            // Optional. Undistributed prizes
	OnlyNewMembers                *bool   `json:"only_new_members,omitempty"`                 // Optional. True if only new members could win
	WasRefunded                   *bool   `json:"was_refunded,omitempty"`                     // Optional. True if canceled due to refund
	PrizeDescription              *string `json:"prize_description,omitempty"`                // Optional. Description of additional prize
}

// GiveawayCompleted represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`                    // Number of winners in the giveaway
	UnclaimedPrizeCount *int     `json:"unclaimed_prize_count,omitempty"` // Optional. Undistributed prizes
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`      // Optional. Message with completed giveaway
	IsStarGiveaway      *bool    `json:"is_star_giveaway,omitempty"`      // Optional. True if it's a Telegram Star giveaway
}

// VideoChatScheduled represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	StartDate int `json:"start_date"` // Unix timestamp for when the video chat is supposed to start
}

// VideoChatStarted represents a service message about a video chat started in the chat.
type VideoChatStarted struct {
	// Currently holds no information
}

// VideoChatEnded represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	Duration int `json:"duration"` // Video chat duration in seconds
}

// VideoChatParticipantsInvited represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"` // New members that were invited to the video chat
}

// WebAppData describes data sent from a Web App to the bot.
type WebAppData struct {
	Data       string `json:"data"`        // The data sent from the Web App
	ButtonText string `json:"button_text"` // Text of the web_app keyboard button that opened the Web App
}

// InlineKeyboardMarkup represents an inline keyboard that appears next to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows
}

// InlineKeyboardButton represents one button of an inline keyboard.
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`                                       // Label text on the button
	URL                          string                       `json:"url,omitempty"`                              // Optional URL to be opened when the button is pressed
	CallbackData                 string                       `json:"callback_data,omitempty"`                    // Optional data sent in a callback query to the bot
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`                          // Optional Web App description
	LoginURL                     *LoginUrl                    `json:"login_url,omitempty"`                        // Optional HTTPS URL for user authorization
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`              // Optional inline query for chat selection
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"` // Optional inline query for current chat
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`  // Optional chat selection with type
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`                    // Optional game description
	Pay                          bool                         `json:"pay,omitempty"`                              // Optional Pay button
}

// Reaction представляет возможные типы реакций.
type ReactionType struct {
	Type          string `json:"type"`            // Тип реакции
	Emoji         string `json:"emoji"`           // Эмодзи реакции
	CustomEmojiID string `json:"custom_emoji_id"` // Идентификатор кастомного эмодзи
}

// ReactionCount представляет реакцию, добавленную к сообщению, вместе с количеством её добавлений.
type ReactionCount struct {
	Type       ReactionType `json:"type"`        // Тип реакции
	TotalCount int          `json:"total_count"` // Количество раз, когда реакция была добавлена
}

// WebAppInfo describes a Web App.
type WebAppInfo struct {
	URL string `json:"url"` // An HTTPS URL of a Web App to be opened
}

// LoginUrl represents parameters for an inline keyboard button used for user authorization.
type LoginUrl struct {
	URL                string `json:"url"`                            // HTTPS URL with user authorization data
	ForwardText        string `json:"forward_text,omitempty"`         // New text of the button in forwarded messages (optional)
	BotUsername        string `json:"bot_username,omitempty"`         // Username of the bot for authorization (optional)
	RequestWriteAccess bool   `json:"request_write_access,omitempty"` // Request permission to send messages (optional)
}

// SwitchInlineQueryChosenChat represents an inline button for switching to inline mode in a chosen chat.
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`               // Default inline query (optional)
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`    // Allow private chats with users (optional)
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`     // Allow private chats with bots (optional)
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`   // Allow group chats (optional)
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"` // Allow channel chats (optional)
}

// GameHighScore represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int  `json:"position"` // Position in the high score table
	User     User `json:"user"`     // User who achieved the score
	Score    int  `json:"score"`    // Score
}

// GetGameHighScoresParams represents parameters for the getGameHighScores method.
type GetGameHighScoresParams struct {
	UserID          int64  `json:"user_id"`                     // Target user ID (required)
	ChatID          int64  `json:"chat_id,omitempty"`           // Unique identifier for the target chat (optional)
	MessageID       int64  `json:"message_id,omitempty"`        // Identifier of the sent message (optional)
	InlineMessageID string `json:"inline_message_id,omitempty"` // Identifier of the inline message (optional)
}

// CallbackGame is a placeholder for game-related information.
type CallbackGame struct {
	// Currently holds no information.
}

// SetGameScoreParams represents parameters for setting a user's game score.
type SetGameScoreParams struct {
	UserID             int64  `json:"user_id"`                        // User identifier (required)
	Score              int    `json:"score"`                          // New score, must be non-negative (required)
	Force              bool   `json:"force,omitempty"`                // Allow score to decrease (optional)
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"` // Prevent message editing (optional)
	ChatID             int64  `json:"chat_id,omitempty"`              // Unique identifier for the target chat (optional)
	MessageID          int64  `json:"message_id,omitempty"`           // Identifier of the sent message (optional)
	InlineMessageID    string `json:"inline_message_id,omitempty"`    // Identifier of the inline message (optional)
}

// SharedUser represents information about a user shared with the bot.
type SharedUser struct {
	UserID    int64       `json:"user_id"`              // Identifier of the shared user
	FirstName string      `json:"first_name,omitempty"` // First name of the user (optional)
	LastName  string      `json:"last_name,omitempty"`  // Last name of the user (optional)
	Username  string      `json:"username,omitempty"`   // Username of the user (optional)
	Photo     []PhotoSize `json:"photo,omitempty"`      // User's photo (optional)
}

// OrderInfo represents information about an order.
type OrderInfo struct {
	Name            string          `json:"name,omitempty"`             // User name (optional)
	PhoneNumber     string          `json:"phone_number,omitempty"`     // User's phone number (optional)
	Email           string          `json:"email,omitempty"`            // User email (optional)
	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"` // User shipping address (optional)
}

// ShippingAddress represents a shipping address for an order.
type ShippingAddress struct {
	Country    string `json:"country"`     // Country name
	State      string `json:"state"`       // State name
	City       string `json:"city"`        // City name
	Street     string `json:"street"`      // Street address
	House      string `json:"house"`       // House number
	PostalCode string `json:"postal_code"` // Postal code
}

// ChatBoost represents information about a chat boost.
type ChatBoost struct {
	BoostID        string          `json:"boost_id"`        // Unique identifier of the boost
	AddDate        int64           `json:"add_date"`        // Point in time (Unix timestamp) when the chat was boosted
	ExpirationDate int64           `json:"expiration_date"` // Point in time (Unix timestamp) when the boost will automatically expire
	Source         ChatBoostSource `json:"source"`          // Source of the added boost
}

// ChatBoostSource represents the source of the added boost.
type ChatBoostSource struct {
	SourceType string `json:"source_type"`       // Type of the source (e.g., "premium")
	Details    string `json:"details,omitempty"` // Additional details about the source (optional)
}

// EncryptedPassportElement describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	Type        string         `json:"type"`                   // Element type
	Data        string         `json:"data,omitempty"`         // Base64-encoded encrypted element data
	PhoneNumber string         `json:"phone_number,omitempty"` // User's verified phone number
	Email       string         `json:"email,omitempty"`        // User's verified email address
	Files       []PassportFile `json:"files,omitempty"`        // Array of encrypted files with documents
	FrontSide   *PassportFile  `json:"front_side,omitempty"`   // Encrypted file with the front side of the document
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"` // Encrypted file with the reverse side of the document
	Selfie      *PassportFile  `json:"selfie,omitempty"`       // Encrypted file with the selfie of the user holding a document
	Translation []PassportFile `json:"translation,omitempty"`  // Array of encrypted files with translated versions of documents
	Hash        string         `json:"hash"`                   // Base64-encoded element hash
}

// EncryptedCredentials describes data required for decrypting and authenticating EncryptedPassportElement.
type EncryptedCredentials struct {
	Data   string `json:"data"`   // Base64-encoded encrypted JSON-serialized data
	Hash   string `json:"hash"`   // Base64-encoded data hash for authentication
	Secret string `json:"secret"` // Base64-encoded secret required for data decryption
}

// PassportFile represents an encrypted file related to a Telegram Passport element.
type PassportFile struct {
	FileID   string `json:"file_id"`   // Unique file identifier
	FileSize int64  `json:"file_size"` // File size
	UniqueID string `json:"unique_id"` // Unique identifier for the file
}

// File represents a file ready to be downloaded.
type File struct {
	FileID       string `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"`      // Unique identifier for this file, consistent over time and for different bots
	FileSize     int64  `json:"file_size,omitempty"` // Optional. File size in bytes
	FilePath     string `json:"file_path,omitempty"` // Optional. File path to download the file
}

// PollOption represents an answer option in a poll.
type PollOption struct {
	Text         string          `json:"text"`                    // Option text, 1-100 characters
	TextEntities []MessageEntity `json:"text_entities,omitempty"` // Special entities in the option text
	VoterCount   int             `json:"voter_count"`             // Number of voters for this option
}

// InaccessibleMessage представляет недоступное сообщение.
type InaccessibleMessage struct {
	Chat      Chat `json:"chat"`       // Чат, к которому принадлежит сообщение
	MessageID int  `json:"message_id"` // Уникальный идентификатор сообщения в чате
	Date      int  `json:"date"`       // Всегда 0, используется для различения обычных и недоступных сообщений
}

// MaybeInaccessibleMessage представляет сообщение, которое может быть недоступным.
type MaybeInaccessibleMessage struct {
	Message         *Message             `json:"message,omitempty"`              // Обычное сообщение
	InaccessibleMsg *InaccessibleMessage `json:"inaccessible_message,omitempty"` // Недоступное сообщение
}

type ChatMember struct {
	Status              string  `json:"status"`                           // Статус участника в чате
	User                User    `json:"user"`                             // Информация о пользователе
	IsAnonymous         *bool   `json:"is_anonymous,omitempty"`           // Истина, если присутствие пользователя скрыто
	CustomTitle         *string `json:"custom_title,omitempty"`           // Пользовательский заголовок (опционально)
	CanBeEdited         *bool   `json:"can_be_edited,omitempty"`          // Истина, если бот может редактировать права администратора
	CanManageChat       *bool   `json:"can_manage_chat,omitempty"`        // Истина, если администратор может управлять чатом
	CanDeleteMessages   *bool   `json:"can_delete_messages,omitempty"`    // Истина, если администратор может удалять сообщения
	CanManageVideoChats *bool   `json:"can_manage_video_chats,omitempty"` // Истина, если администратор может управлять видеозвонками
	CanRestrictMembers  *bool   `json:"can_restrict_members,omitempty"`   // Истина, если администратор может ограничивать участников
	CanPromoteMembers   *bool   `json:"can_promote_members,omitempty"`    // Истина, если администратор может продвигать других
	CanChangeInfo       *bool   `json:"can_change_info,omitempty"`        // Истина, если пользователь может изменять информацию о чате
	CanInviteUsers      *bool   `json:"can_invite_users,omitempty"`       // Истина, если пользователь может приглашать новых участников
	CanPostStories      *bool   `json:"can_post_stories,omitempty"`       // Истина, если администратор может размещать истории
	CanEditStories      *bool   `json:"can_edit_stories,omitempty"`       // Истина, если администратор может редактировать истории
	CanDeleteStories    *bool   `json:"can_delete_stories,omitempty"`     // Истина, если администратор может удалять истории
	CanPostMessages     *bool   `json:"can_post_messages,omitempty"`      // Истина, если администратор может размещать сообщения (для каналов)
	CanEditMessages     *bool   `json:"can_edit_messages,omitempty"`      // Истина, если администратор может редактировать сообщения (для каналов)
	CanPinMessages      *bool   `json:"can_pin_messages,omitempty"`       // Истина, если пользователь может закреплять сообщения
	CanManageTopics     *bool   `json:"can_manage_topics,omitempty"`      // Истина, если пользователь может управлять темами в супер-группах
	UntilDate           *int    `json:"until_date,omitempty"`             // Дата, когда ограничения будут сняты; Unix time
	IsMember            *bool   `json:"is_member,omitempty"`              // Истина, если пользователь является членом чата в момент запроса
}

type BackgroundType struct {
	Type              string          `json:"type"`                           // Type of the background (fill, wallpaper, pattern, chat_theme)
	Fill              *BackgroundFill `json:"fill,omitempty"`                 // Background fill information (optional)
	Document          *Document       `json:"document,omitempty"`             // Document with the wallpaper or pattern (optional)
	DarkThemeDimming  *int            `json:"dark_theme_dimming,omitempty"`   // Dimming of the background in dark themes (0-100, optional)
	IsBlurred         *bool           `json:"is_blurred,omitempty"`           // True if the wallpaper is blurred (optional)
	IsMoving          *bool           `json:"is_moving,omitempty"`            // True if the background moves when the device is tilted (optional)
	Intensity         *int            `json:"intensity,omitempty"`            // Intensity of the pattern (0-100, optional)
	IsInverted        *bool           `json:"is_inverted,omitempty"`          // True if fill applies only to the pattern (optional)
	ThemeName         *string         `json:"theme_name,omitempty"`           // Name of the chat theme (optional)
	IconColor         *int            `json:"icon_color,omitempty"`           // Color of the topic icon in RGB format (optional)
	IconCustomEmojiID *string         `json:"icon_custom_emoji_id,omitempty"` // Custom emoji ID for the topic icon (optional)
}

type BackgroundFill struct {
	Type          string `json:"type"`                     // Type of the background fill (solid, gradient, freeform_gradient)
	Color         *int   `json:"color,omitempty"`          // Color of the background fill in RGB24 format (for solid fill, optional)
	TopColor      *int   `json:"top_color,omitempty"`      // Top color of the gradient in RGB24 format (optional)
	BottomColor   *int   `json:"bottom_color,omitempty"`   // Bottom color of the gradient in RGB24 format (optional)
	RotationAngle *int   `json:"rotation_angle,omitempty"` // Clockwise rotation angle of the gradient fill (0-359, optional)
	Colors        *[]int `json:"colors,omitempty"`         // List of 3 or 4 base colors for freeform gradient in RGB24 format (optional)
}
