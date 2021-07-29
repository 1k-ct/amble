package persistence

// func Test_actionsPersistence_Like(t *testing.T) {
// 	type args struct {
// 		staticID string
// 	}
// 	tests := []struct {
// 		name    string
// 		ap      *actionsPersistence
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "いいね＋１テスト",
// 			ap:   &actionsPersistence{},
// 			args: args{
// 				staticID: "f36d8a5b-1438-47b6-8b63-6d0099b61340",
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ap := &actionsPersistence{}
// 			if err := ap.Like(tt.args.staticID); (err != nil) != tt.wantErr {
// 				t.Errorf("actionsPersistence.Like() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_actionsPersistence_Retweet(t *testing.T) {
// 	type args struct {
// 		staticID string
// 	}
// 	tests := []struct {
// 		name    string
// 		ap      *actionsPersistence
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "リツイート＋１テスト",
// 			ap:   &actionsPersistence{},
// 			args: args{
// 				staticID: "3b19ddcc-b75b-4706-a3cc-d60a19763b53",
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ap := &actionsPersistence{}
// 			if err := ap.Retweet(tt.args.staticID); (err != nil) != tt.wantErr {
// 				t.Errorf("actionsPersistence.Retweet() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_actionsPersistence_Reply(t *testing.T) {
// 	type args struct {
// 		staticID string
// 	}
// 	tests := []struct {
// 		name    string
// 		ap      *actionsPersistence
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "",
// 			ap:   &actionsPersistence{},
// 			args: args{
// 				staticID: "3b19ddcc-b75b-4706-a3cc-d60a19763b53",
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ap := &actionsPersistence{}
// 			if err := ap.Reply(tt.args.staticID); (err != nil) != tt.wantErr {
// 				t.Errorf("actionsPersistence.Reply() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
