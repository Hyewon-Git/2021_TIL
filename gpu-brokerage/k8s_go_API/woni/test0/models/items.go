package models

import (
	"fmt"
	_ "log"

	_ "gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	//gorm.Model
	ID     int    `json:"id"`
	PNAME  string `json:"pname"`
	PPRICE int    `json:"pprice"`
}

//addItem 삽입
func AddItem(item Product) {

	result := db.Create(&item) //생성할 데이터의 포인터 넘기
	//CheckErr(result.Error)

	fmt.Println("insert count: ", result.RowsAffected)

}

//getAllItem 전체조회
func AllItem() (result []Product) {
	/*
		var products []Product
		db.Find(&products, "PNAME = ?", "apple")
		return products
	*/

	//복수 row를 가진 SQL쿼리  Select

	var item Product
	rows, err := db.Model(&item).Rows()
	CheckErr(err)

	for rows.Next() {
		item := Product{}
		err := rows.Scan(
			&item.ID,
			&item.PNAME,
			&item.PPRICE,
		)
		CheckErr(err)
		result = append(result, item)
	}

	return
}

//SearchItem 특정상품검색
func SearchItem(itemName string) (result []Product) {
	// 특정 상품명 p_name으로 DB조회
	var item Product
	rows, err := db.Model(&item).Where("p_name = ?", itemName).Rows()
	CheckErr(err)
	//해당데이터 rows에서 result로 값넣고  에러확인!
	for rows.Next() {
		item := Product{}
		err := rows.Scan(&item.ID, &item.PNAME, &item.PPRICE)
		CheckErr(err)
		result = append(result, item)
	}

	return
}
