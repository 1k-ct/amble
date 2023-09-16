package ffile

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// test
//     ├─build1
//     │      IMG_0060.jpeg
//     │      IMG_0061.jpeg
//     │      IMG_0062.jpeg
//     │      IMG_0082.jpeg
//     │      IMG_0083.jpeg
//     │      IMG_0084.jpeg
//     │      IMG_0100.jpeg
//     │      IMG_0101.jpeg
//     │      IMG_0102.jpeg
//     │      IMG_0103.jpeg
//     │      IMG_0104.jpeg
//     │      IMG_0105.jpeg
//     │
//     └─build2  // ここから下が自動で作成される
//         ├─folder1
//         │      IMG_0060.jpeg
//         │      IMG_0061.jpeg
//         │      IMG_0062.jpeg
//         │      IMG_0082.jpeg
//         │      IMG_0083.jpeg
//         │      IMG_0084.jpeg
//         │
//         └─folder2
//                 IMG_0100.jpeg
//                 IMG_0101.jpeg
//                 IMG_0102.jpeg
//                 IMG_0103.jpeg
//                 IMG_0104.jpeg
//                 IMG_0105.jpeg

func TestMain(t *testing.T) {
	// ファイルが格納されているディレクトリと新しいフォルダのベースディレクトリを指定
	sourceDir := "./test/build1"   // ファイルが格納されているディレクトリのパス
	destBaseDir := "./test/build2" // 新しいフォルダを作成するベースディレクトリのパス
	batchSize := 6                 // 1つのフォルダに含めるファイルの数

	// ファイルのリストを取得
	files, err := getFiles(sourceDir)
	if err != nil {
		fmt.Printf("ファイルのリストを取得できませんでした: %v\n", err)
		return
	}
	for _, file := range files {
		fmt.Println(file)
	}

	// フォルダを作成するためのカウンタ
	folderCount := 1

	// ファイルをバッチごとに処理
	for i := 0; i < len(files); i += batchSize {
		// 新しいフォルダを作成
		folderName := fmt.Sprintf("folder%d", folderCount)
		folderPath := filepath.Join(destBaseDir, folderName)
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			fmt.Printf("フォルダの作成に失敗しました: %v\n", err)
			return
		}

		// ファイルをコピー
		for j := i; j < i+batchSize && j < len(files); j++ {
			sourceFile := files[j]
			destFile := filepath.Join(folderPath, filepath.Base(sourceFile))
			if err := copyFile(sourceFile, destFile); err != nil {
				fmt.Printf("ファイルのコピーに失敗しました: %v\n", err)
				return
			}
		}

		folderCount++
	}

	fmt.Println("ファイルの分割が完了しました。")
}
func getFiles(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// ファイルをコピー
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
func ListFilesInFolder(t *testing.T) {
	folderPath := "./test/build1" // フォルダのパスを指定してください

	files, err := listFilesInFolder(folderPath)
	if err != nil {
		fmt.Printf("ファイルリストの取得に失敗しました: %v\n", err)
		return
	}

	fmt.Println("フォルダ内のファイルリスト:")
	for _, file := range files {
		fmt.Println(file)
	}
}

func listFilesInFolder(folderPath string) ([]string, error) {
	var files []string

	// フォルダ内のファイル一覧を取得
	fileInfoList, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	// ファイル情報を名前でソート
	sort.Slice(fileInfoList, func(i, j int) bool {
		return fileInfoList[i].Name() < fileInfoList[j].Name()
	})

	// フォルダ内のファイルのパスをリストに追加
	for _, fileInfo := range fileInfoList {
		if !fileInfo.IsDir() {
			files = append(files, filepath.Join(folderPath, fileInfo.Name()))
		}
	}

	return files, nil
}
