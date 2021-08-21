package database_test

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type Book struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT;not null;"`
	Title       string `gorm:"not null"`
	PublisherID uint   `gorm:"not null"`
	Publisher   Publisher
	Authors     []Author `gorm:"many2many:author_books"`
}

type Author struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT;not null;"`
	Name  string `gorm:"not null"`
	Books []Book `gorm:"many2many:author_books"`
}

type Publisher struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT;not null;"`
	Name  string `gorm:"not null"`
	Books []Book
}
type Post struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `gorm:"size:255;not null"`
	Content     string `gorm:"type:text not null"`
	Status      string `gorm:"size:255;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt sql.NullTime
	Tags        []Tag `gorm:"many2many:post_tags"` // new!!
}

type Tag struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:255;not null;unique"`
	Posts     []Post `gorm:"many2many:post_tags"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Vtuber struct {
	gorm.Model
	Name       string
	VtuberTags []VtuberTag `gorm:"ForeingKey:VtuberTagID"`
}
type VtuberTag struct {
	// gorm.Model
	VtuberID uint
	Tag      string
}

func migrate(db *gorm.DB) error {
	// 外部キーも設定すべきですが無視します
	if err := db.AutoMigrate(&Book{}).
		AutoMigrate(&Author{}).
		AutoMigrate(&Publisher{}).
		Error; err != nil {
		return err
	}
	return nil
}
func TestMain(t *testing.T) {
	// config, err := database.NewLocalDB("user", "password", "sample")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// db, err := config.Connect()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer db.Close()

	// vtubers := []Vtuber{}
	// if err := db.Where("name = ?", "name1").Preload("VtuberTags").Find(&vtubers).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println(vtubers)

	// vtuberTags := []VtuberTag{}
	// if err := db.Where("tag = ?", "tag2").Find(&vtuberTags).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// for _, v := range vtuberTags {
	// 	fmt.Println(v.VtuberID)
	// }
	// fmt.Println(vtuberTags)
}

const (
	publisherName = "test-publisher"
	authorName1   = "test-author-1"
	authorName2   = "test-author-2"
	BookTitle1    = "test-book-1"
	BookTitle2    = "test-book-2"
)

func seeds(db *gorm.DB) error {
	if !db.First(&Publisher{Name: publisherName}).RecordNotFound() {
		return xerrors.Unwrap(gorm.ErrRecordNotFound)
	}

	publisher := Publisher{Name: publisherName}
	if err := db.Create(&publisher).Error; err != nil {
		return err
	}

	author1 := Author{Name: authorName1}
	author2 := Author{Name: authorName2}
	if err := db.Create(&author1).Create(&author2).Error; err != nil {
		return err
	}

	book1 := Book{Title: BookTitle1, PublisherID: publisher.ID}
	book2 := Book{Title: BookTitle2, PublisherID: publisher.ID}
	if err := db.Model(&author1).Association("Books").Append(&book1).Append(&book2).Error; err != nil {
		return err
	}
	if err := db.Model(&author2).Association("Books").Append(&book2).Error; err != nil {
		return err
	}
	return nil
}
func getBook(db *gorm.DB) ([]Book, error) {
	var books []Book
	if err := db.Where("ID = ?", 1).Preload("Publisher").Preload("Authors").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
func getAuthor(db *gorm.DB) ([]Author, error) {
	var authors []Author
	if err := db.Preload("Books").Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}
func getPublisher(db *gorm.DB) ([]Publisher, error) {
	var publishers []Publisher
	if err := db.Preload("Books").Find(&publishers).Error; err != nil {
		return nil, err
	}
	return publishers, nil
}
