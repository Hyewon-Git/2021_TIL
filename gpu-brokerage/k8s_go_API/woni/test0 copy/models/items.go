package models

import (
	"fmt"
	_ "log"

	_ "github.com/go-sql-driver/mysql"
)

type ItemStruct struct {
	ID     int    `json:"id"`
	PNAME  string `json:"pname"`
	PPRICE int    `json:"pprice"`
}

//addItem 삽입
func AddItem(item ItemStruct) {

	//    result, err := db.Exec("INSERT INTO test1 VALUES (?, ?)", 11, "Jack")
	//	  로 한줄에 적어도되고 아래처럼  sql 문이랑 입력인자문을 따로적어도됨!
	stmt, err := db.Prepare("INSERT product SET p_name=?, p_price=?")
	CheckErr(err)
	res, err := stmt.Exec(item.PNAME, item.PPRICE)
	CheckErr(err)

	nRow, err := res.RowsAffected()
	fmt.Println("insert count: ", nRow)

}

//getAllItem 전체조회
func AllItem() (result []ItemStruct) {
	//복수 row를 가진 SQL쿼리  Select
	rows, err := db.Query("SELECT id, p_name,p_price FROM product")
	CheckErr(err)
	defer rows.Close() //지연하여 닫기

	//DB에서 가져와서 보여주는 result
	for rows.Next() {
		item := ItemStruct{}
		err := rows.Scan(
			&item.ID,
			&item.PNAME,
			&item.PPRICE,
		)
		CheckErr(err)
		result = append(result, item)
	}
	//defer Close()
	return
}

//SearchItem 특정상품검색
func SearchItem(itemName string) (result []ItemStruct) {
	// 특정 상품명 p_name으로 DB조회
	rows, err := db.Query("SELECT id, p_name,p_price FROM product where p_name=?", itemName)
	fmt.Println("err이전 ")
	CheckErr(err)
	defer rows.Close()

	//해당데이터 rows에서 result로 값넣고  에러확인!
	for rows.Next() {
		item := ItemStruct{}
		err := rows.Scan(&item.ID, &item.PNAME, &item.PPRICE)
		CheckErr(err)
		result = append(result, item)
	}
	return
}
