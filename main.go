package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	// gorm.Modelをつけると、idとCreatedAtとUpdatedAtとDeletedAtが作られる
	gorm.Model

	Name     string
	Age      int
	IsActive bool
}

func main() {
	db := dbInit()

	// db.AutoMigrate(&Product{})
	// db.AutoMigrate(&User{})

	// 単体取得
	getOne(db)

	// 複数取得
	// find(db)

	// 単体追加
	// insert(db)

	// 複数追加
	// inserts(db)

	// delete(db)
}

func dbInit() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/sample_db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// 単体作成
func insert(db *gorm.DB) {
	// t := true
	user := User{
		Name: "太郎",
		Age:  20,
		// IsActive: &t,
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// 複数作成
func inserts(db *gorm.DB) {
	users := []User{
		{
			Name:     "花子",
			Age:      25,
			IsActive: true,
		},
		{
			Name:     "龍太郎",
			Age:      30,
			IsActive: false,
		},
		{
			Name:     "太一",
			Age:      35,
			IsActive: false,
		},
	}
	result := db.Create(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// 単体取得
func getOne(db *gorm.DB) {
	user1 := User{}

	// SELECT * FROM users ORDER BY id LIMIT 1;
	// 昇順で単体取得
	result1 := db.First(&user1)
	fmt.Println("first:", user1)
	if result1.Error != nil {
		log.Fatal(result1.Error)
	}
	fmt.Println("count:", result1.RowsAffected)

	// check error ErrRecordNotFound
	errors.Is(result1.Error, gorm.ErrRecordNotFound)

	// SELECT * FROM users LIMIT 1;
	// 何も指定せず、単体取得
	user2 := User{}
	result2 := db.Take(&user2)
	fmt.Println("take:", user2)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}

	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	user3 := User{}
	// 降順で単体取得
	result3 := db.Last(&user3)
	fmt.Println("last:", user3)
	if result3.Error != nil {
		log.Fatal(result3.Error)
	}

	// プライマリーキーで取得
	user4 := User{}
	result4 := db.First(&user4, 2)
	if result4.Error != nil {
		log.Fatal(result4.Error)
	}

	user5 := User{}
	result5 := db.First(&user5, "id = ?", 3)
	if result5.Error != nil {
		log.Fatal(result5.Error)
	}

}

// 複数取得
func find(db *gorm.DB) {
	users := []User{}
	result := db.Find(&users)
	fmt.Println("user:", users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

// 全件取得
func scan(db *gorm.DB) {
	users := []User{}
	db.Find(&users)
	fmt.Println("user:", users)
}

// 更新(全てのカラム更新)
func save(db *gorm.DB) {
	user := User{}
	db.First(&user)

	user.Name = "太郎"
	db.Save(&user)
	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
	// Saveは 、SQLを実行するときにすべてのフィールドを更新します。
}

// 単一のカラムを更新する
func update(db *gorm.DB) {
	db.Model(&User{}).Where("id = 1").Update("name", "たかし")
}

// 複数のカラムを更新する
func updates(db *gorm.DB) {
	db.Model(&User{}).Updates(User{Name: "太郎"})

}

// 削除
// gorm.DeletedAt フィールドがモデルに含まれている場合、そのモデルは自動的に論理削除されるようになります。
// 論理削除されたレコードは取得処理時に無視されます
func delete(db *gorm.DB) {
	db.Where("id = 1").Delete(&User{})
	// DELETE FROM users WHERE id = 1;

	// db.Delete(&User{}, 1)
	// DELETE FROM users WHERE id = 1;

	// users := []User{}
	// db.Delete(&users, []int{1, 2, 3})
	// DELETE FROM users WHERE id IN (1,2,3);

}

// Update with conditions
// 単一のカラムを更新する
// db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

// 複数のカラムを更新する
// Updatesは構造体もしくはmap[string]interface{}での更新に対応しています
// 構造体での更新の時はゼロ値の更新はしない
// ゼロ値のフィールドも更新対象に含める場合は、更新にmapを使用するか、 Select で更新するフィールドを指定してください。

// 一括更新
// Modelで主キーを指定していない場合、GORMは一括更新を行います。

// テーブル指定
// db.Table("users")

// 条件指定
// db.Where()

// 特定のカラムのみ取得
// db.Select("name", "age").Find(&users)

// 特定のカラムのみ除外して取得
// db.Omit("name", "age").Find(&users)
