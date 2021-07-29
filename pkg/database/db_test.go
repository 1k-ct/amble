package database_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/1k-ct/twitter-dem/pkg/database"
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
	config, err := database.NewLocalDB("user", "password", "sample")
	if err != nil {
		t.Fatal(err)
	}
	db, err := config.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	// if err := seeds(db); err != nil {
	// 	t.Fatal(err)
	// }
	// books, err := getBook(db)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println(books)
	// authors, err := getAuthor(db)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println(authors)
	// publishers, err := getPublisher(db)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println(publishers)
	// if err := db.AutoMigrate(&Post{}, &Tag{}).Error; err != nil {
	// 	t.Fatal(err)
	// }

	// if !db.First(&Post{}).RecordNotFound() {
	// 	t.Fatal(err)
	// }

	// tag := &Tag{
	// 	Name:      "tag-name1",
	// 	Posts:     []Post{},
	// 	CreatedAt: time.Time{},
	// 	UpdatedAt: time.Time{},
	// }

	// post := Post{
	// 	ID:      4,
	// 	Title:   "title3",
	// 	Content: "content3",
	// 	PublishedAt: sql.NullTime{
	// 		Time:  time.Now(),
	// 		Valid: true,
	// 	},
	// 	Status: "3",
	// }

	// if err := db.Create(&post).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// tag := Tag{
	// 	Name: "tag2",
	// }
	// if err := db.Create(&tag).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// postTag := Post{
	// 	ID:      0,
	// 	Title:   "title6",
	// 	Content: "6",
	// 	Status:  "6",
	// }
	// if err := db.Model(&tag).Association("Posts").Append(&postTag).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// tags := []Tag{}

	// if err := db.Preload("Posts").Find(&tags).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// for _, v := range tags {
	// 	fmt.Println(v.ID, v.Name, v.Posts)
	// }

	// posts := []Post{}
	// if err := db.Preload("Tags").Find(&posts).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// // fmt.Println(posts)
	// for _, v := range posts {
	// 	fmt.Println(v.ID, v.Title, v.Content, v.Status, v.PublishedAt.Time, v.PublishedAt.Valid, v.Tags)
	// }

	// if err := db.DropTableIfExists(&Vtuber{}, &VtuberTag{}).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// if err := db.AutoMigrate(&Vtuber{}).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// if err := db.AutoMigrate(&VtuberTag{}).AddForeignKey("vtuber_id", "vtubers(id)", "CASCADE", "CASCADE").Error; err != nil {
	// 	t.Fatal(err)
	// }
	// vtuber := &Vtuber{
	// 	Name: "name1",
	// 	VtuberTags: []VtuberTag{{
	// 		Tag: "tag1",
	// 	}, {
	// 		Tag: "tag2",
	// 	}},
	// }
	// if err := db.Create(vtuber).Error; err != nil {
	// 	t.Fatal(err)
	// }
	vtubers := []Vtuber{}
	if err := db.Where("name = ?", "name1").Preload("VtuberTags").Find(&vtubers).Error; err != nil {
		t.Fatal(err)
	}
	fmt.Println(vtubers)

	vtuberTags := []VtuberTag{}
	if err := db.Where("tag = ?", "tag2").Find(&vtuberTags).Error; err != nil {
		t.Fatal(err)
	}
	for _, v := range vtuberTags {
		fmt.Println(v.VtuberID)
	}
	fmt.Println(vtuberTags)
	// vtuberTag := VtuberTag{Tag: "tag1"}
	// vtubers := Vtuber{
	// 	Name:       "name2",
	// 	VtuberTags: []VtuberTag{{Tag: "tag2"}},
	// }
	// if err := db.Create(&vtuberTag).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// if err := db.Model(&vtuberTag).Association("VtuberTags").Append(&vtubers).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// if err := db.Find(&vtubers).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println(vtubers)

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
