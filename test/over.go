package test

import (
	"fmt"
	"errors"

	"github.com/shopspring/decimal"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/dbx"
)

var db *dbx.Database

func init() {
	settings := dbx.Settings{
		DriverName: "mysql",
		Host:       "127.0.0.1:3306",
		User:       "root",
		Password:   "root",
		Database:   "resk",
		Options: map[string]string{
			"parseTime": "true",
		},
	}
	var err error
	db, err = dbx.Open(settings)
	if err != nil {
		fmt.Println(err)
	}
	db.RegisterTable(&GoodsSigned{}, "goods")
	db.RegisterTable(&GoodsUnsigned{}, "goods_unsigned")
}

// 事务锁方案
func UpdateForLock(g Goods) {
	// 通过db.tx函数构建事务锁代码块
	err := db.Tx(func(runner *dbx.TxRunner) error {
		// 第一步：锁定需要修改的资源，也就是需要修改的数据行
		// 编写事务锁查询语句，使用for update子句来锁定资源
		query := "select * from goods " +
			"where envelope_no=? " +
			"for update"
		out := &GoodsSigned{}
		ok, err := runner.Get(out, query, g.EnvelopeNo)
		if !ok || err != nil {
			return err
		}
		// 第二步：计算剩余金额和剩余数量
		subAmount := decimal.NewFromFloat(0.01)
		remainAmount := out.RemainAmount.Sub(subAmount)
		remainQuantity := out.RemainQuantity - 1
		// 第三步：执行更新
		update := "update goods " +
			"set remain_amount=?,remain_quantity=? " +
			"where envelope_no=?"
		_, row, err := db.Execute(update, remainAmount, remainQuantity, g.EnvelopeNo)
		if err !=nil{
			return err
		}
		if row < 1{
			return errors.New("库存扣减失败")
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

// 数据库无符号类型+直接更新方案
func UpdateForUnsigned(g Goods) {
	update := "update goods_unsigned " +
		"set remain_amount=remain_amount-?, " +
		"remain_quantity=remain_quantity-? " +
		"where envelope_no=?"
	_, row, err := db.Execute(update, 0.01, 1, g.EnvelopeNo)
	if err != nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}

// 乐观锁方案
func UpdateForOptimistic(g Goods) {
	update := "update goods " +
		"set remain_amount=remain_amount-?, " +
		"remain_quantity=remain_quantity-? " +
		"where envelope_no=? " +
		"and remain_amount>=? " +
		"and remain_quantity>=?"
	_, row, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}

// 乐观锁+无符号字段双保险方案
func UpdateForOptimisticAndUnsigned(g Goods) {
	update := "update goods_unsigned " +
		"set remain_amount=remain_amount-?, " +
		" remain_quantity=remain_quantity-? " +
		" where envelope_no=? " +
		" and remain_amount>=? " +
		" and remain_quantity>=? "
	_, row, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		fmt.Println(err)
	}
	if row < 1 {
		fmt.Println("扣减失败")
	}
}
