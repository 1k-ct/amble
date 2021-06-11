package chatjson

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

type ChatDataJsonInEmoji struct {
	Addchatitemaction struct {
		Item struct {
			Livechattextmessagerenderer struct {
				Message struct {
					Runs []struct {
						Text  string `json:"text,omitempty"`
						Emoji struct {
							Emojiid     string   `json:"emojiId"`
							Shortcuts   []string `json:"shortcuts"`
							Searchterms []string `json:"searchTerms"`
							Image       struct {
								Thumbnails []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
								Accessibility struct {
									Accessibilitydata struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
							} `json:"image"`
							Iscustomemoji bool `json:"isCustomEmoji"`
						} `json:"emoji,omitempty"`
					} `json:"runs"`
				} `json:"message"`
				Authorname struct {
					Simpletext string `json:"simpleText"`
				} `json:"authorName"`
				Authorphoto struct {
					Thumbnails []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"thumbnails"`
				} `json:"authorPhoto"`
				Contextmenuendpoint struct {
					Commandmetadata struct {
						Webcommandmetadata struct {
							Ignorenavigation bool `json:"ignoreNavigation"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					Livechatitemcontextmenuendpoint struct {
						Params string `json:"params"`
					} `json:"liveChatItemContextMenuEndpoint"`
				} `json:"contextMenuEndpoint"`
				ID            string `json:"id"`
				Timestampusec string `json:"timestampUsec"`
				Authorbadges  []struct {
					Livechatauthorbadgerenderer struct {
						Customthumbnail struct {
							Thumbnails []struct {
								URL string `json:"url"`
							} `json:"thumbnails"`
						} `json:"customThumbnail"`
						Tooltip       string `json:"tooltip"`
						Accessibility struct {
							Accessibilitydata struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
					} `json:"liveChatAuthorBadgeRenderer"`
				} `json:"authorBadges"`
				Authorexternalchannelid  string `json:"authorExternalChannelId"`
				Contextmenuaccessibility struct {
					Accessibilitydata struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"contextMenuAccessibility"`
				Timestamptext struct {
					Simpletext string `json:"simpleText"`
				} `json:"timestampText"`
			} `json:"liveChatTextMessageRenderer"`
		} `json:"item"`
		Clientid string `json:"clientId"`
	} `json:"addChatItemAction"`
}

type ChatDataJsons struct {
	ChatDataJsonInEmojis []ChatDataJsonInEmoji
}

func BufioScannerJsonFile(fileName string) (*ChatDataJsons, error) {
	chatDataJson := &ChatDataJsons{}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		jsonBytes := ([]byte)(text)
		data := new(ChatDataJsonInEmoji)

		if err := json.Unmarshal(jsonBytes, data); err != nil {
			return nil, err
		}
		chatDataJson.ChatDataJsonInEmojis = append(chatDataJson.ChatDataJsonInEmojis, *data)
	}
	return chatDataJson, nil
}

type ChatAttendee struct {
	Authorname []string `json:"authorname"`
}

func FitchUsersName(fileName string) (*ChatAttendee, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	chatAttendee := &ChatAttendee{}

	for scanner.Scan() {
		text := scanner.Text()
		jsonBytes := ([]byte)(text)
		data := new(ChatDataJsonInEmoji)

		if err := json.Unmarshal(jsonBytes, data); err != nil {
			return nil, err
		}
		authorName := data.Addchatitemaction.Item.Livechattextmessagerenderer.Authorname.Simpletext
		if !IsOverlappingJsonValue(authorName, chatAttendee.Authorname) {
			chatAttendee.Authorname = append(chatAttendee.Authorname, authorName)
		}
	}
	return chatAttendee, nil
}

// func BufioScannerJsonFile(){}
// values = 1, 2, 3
// value  = 2
func IsOverlappingJsonValue(value string, values []string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func WriteJsonFile(obj interface{}, outputFileName string) error {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(outputFileName+".json", jsonData, 0644); err != nil {
		return err
	}
	return nil
}
