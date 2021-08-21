package chatjson_test

// var fileName = "./chat-data/Okokxcgr_II.json"

// func TestBufioScannerJsonFile(t *testing.T) {
// 	chatData, err := chatjson.BufioScannerJsonFile(fileName)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	for _, chat := range chatData.ChatDataJsonInEmojis {
// 		c := chat.Addchatitemaction.Item.Livechattextmessagerenderer.Message.Runs
// 		for _, v := range c {
// 			fmt.Print(v.Text)
// 		}
// 		fmt.Println()
// 	}
// }
// func TestFitchUsersName(t *testing.T) {

// 	videoID := "Okokxcgr_II"
// 	chatAttendee, err := chatjson.FitchUsersName(fileName, videoID)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if err := chatjson.WriteJsonFile(chatAttendee, "Okokxcgr_II-names"); err != nil {
// 		t.Error(err)
// 	}
// 	// fmt.Println(chatAttendee.Authorname)
// 	// fmt.Println(len(chatAttendee.Authorname))
// 	if chatAttendee.VideoID != videoID {
// 		t.Errorf("get: %v != want: %v", chatAttendee.VideoID, videoID)
// 	}
// }
// func TestIsOverlappingJsonValue(t *testing.T) {
// 	values := []string{"1", "2", "3"}
// 	value := "5"
// 	if chatjson.IsOverlappingJsonValue(value, values) {
// 		fmt.Println("find")
// 	}
// }
// func TestWriteFileJsonAttendanceRate(t *testing.T) {
// 	chatAttendee, err := chatjson.FitchUsersName(fileName, "Okokxcgr_II")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	data, err := json.Marshal(chatAttendee)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	d := string(data)
// 	d = d + "\n"
// 	if err := chatjson.WriteFileJsonAttendanceRate("add.json", d); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestGetFileNameWithoutExr(t *testing.T) {
// 	path := "path/to/file.name.hoge.name"

// 	filename := chatjson.ParseFileNameWithoutExt(path)
// 	fmt.Println(filename)
// }

// func Test_getFileNameWithoutExt(t *testing.T) {
// 	type args struct {
// 		path string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{
// 			name: "base file path",
// 			args: args{
// 				path: fileName,
// 			},
// 			want: "Okokxcgr_II",
// 		},
// 		{
// 			name: "base file path",
// 			args: args{path: "dir/dir/file.go"},
// 			want: "file",
// 		},
// 		{
// 			name: fileName,
// 			args: args{path: "dir/file"},
// 			want: "file",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := chatjson.ParseFileNameWithoutExt(tt.args.path); got != tt.want {
// 				t.Errorf("getFileNameWithoutExt() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
