package chatjson

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
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
	VideoID    string   `json:"video_id"`
	Authorname []string `json:"authorname"`
}

func FitchUsersName(fileName, videoID string) (*ChatAttendee, error) {
	chatAttendee := &ChatAttendee{}
	chatDataJsons, err := BufioScannerJsonFile(fileName)
	if err != nil {
		return nil, err
	}
	for _, v := range chatDataJsons.ChatDataJsonInEmojis {
		authorName := v.Addchatitemaction.Item.Livechattextmessagerenderer.Authorname.Simpletext
		if !IsOverlappingJsonValue(authorName, chatAttendee.Authorname) {
			chatAttendee.Authorname = append(chatAttendee.Authorname, authorName)
		}
	}
	chatAttendee.VideoID = videoID
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

//
//
// type AttendanceRate struct {
// 	videoID string
// }
// type Member struct {
// 	Name string
// }
func WriteFileJsonAttendanceRate(filename, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}

// ParseFileNameWithoutExt
// "dir/dir/file.go" => "file"
// "dir/file"        => "file"
func ParseFileNameWithoutExt(path string) string {
	basefile := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	return basefile
}
