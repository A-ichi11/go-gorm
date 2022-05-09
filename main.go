package main

import (
	"errors"
	"fmt"

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

	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	// getOne(db)
	find(db)
	// insert(db)
	// inserts(db)
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
	user := User{
		Name: "太郎",
	}
	result := db.Create(&user)
	fmt.Println("result:", result)
	fmt.Println("error:", result.Error)
	fmt.Println("rowAffeted", result.RowsAffected)
}

// 複数作成
func inserts(db *gorm.DB) {
	users := []User{{Name: "花子"}, {Name: "龍太郎"}, {Name: "太一"}}
	result := db.Create(&users)
	fmt.Println("result:", result)
	fmt.Println("error:", result.Error)
	fmt.Println("rowAffeted", result.RowsAffected)
}

// 単体取得
func getOne(db *gorm.DB) {
	user := User{}

	// SELECT * FROM users ORDER BY id LIMIT 1;
	result := db.First(&user)
	fmt.Println("user:", user)
	fmt.Println("result:", result)
	fmt.Println("error:", result.Error)
	fmt.Println("rowAffeted", result.RowsAffected)

	// check error ErrRecordNotFound
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	// SELECT * FROM users LIMIT 1;
	db.Take(&user)
	fmt.Println("user:", user)

	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user)
	fmt.Println("user:", user)

}

// 複数取得
func find(db *gorm.DB) {
	users := []User{}
	result := db.Find(&users)
	fmt.Println("user:", users)
	fmt.Println("result:", result)
	fmt.Println("error:", result.Error)
	fmt.Println("rowAffeted", result.RowsAffected)
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
func delete(db *gorm.DB) {
	db.Where("id = 1").Delete(&User{})
	// DELETE FROM users WHERE id = 1;

	db.Delete(&User{}, 1)
	// DELETE FROM users WHERE id = 1;

	users := []User{}
	db.Delete(&users, []int{1, 2, 3})
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
