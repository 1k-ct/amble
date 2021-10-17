package main

// func TestContextGolangContext(t *testing.T) {
// 	gin.SetMode(gin.ReleaseMode)

//  [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//   - using env:   export GIN_MODE=release
//   - using code:  gin.SetMode(gin.ReleaseMode)
//  gin.SetMode をrelease mode を選択しないとgithub actions のテストが
//  worningにないりテストが失敗する

// 	c, _ := gin.CreateTestContext(httptest.NewRecorder())
// 	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"foo\":\"bar\", \"bar\":\"foo\"}"))
// 	assert.NoError(t, c.Err())
// 	assert.Nil(t, c.Done())
// 	ti, ok := c.Deadline()
// 	fmt.Println(ti, ok)
// 	assert.Equal(t, ti, time.Time{})
// 	assert.False(t, ok)
// 	assert.Equal(t, c.Value(0), c.Request)
// 	assert.Nil(t, c.Value("foo"))

// 	c.Set("foo", "bar")
// 	assert.Equal(t, "bar", c.Value("foo"))
// 	assert.Nil(t, c.Value(1))
// }
