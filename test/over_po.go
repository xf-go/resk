package test

import (
	"time"

	"github.com/shopspring/decimal"
)

// 红包商品表，默认有符号类型字段
type GoodsSigned struct {
	Goods
}

// 红包商品表，无符号类型字段
type GoodsUnsigned struct {
	Goods
}

// 红包商品
type Goods struct {
	RemainAmount   decimal.Decimal `db:"remain_amount"`        //红包剩余金额额
	RemainQuantity int             `db:"remain_quantity"`      //红包剩余数量
	EnvelopeNo     string          `db:"envelope_no,uni"`      //红包编号,红包唯一标识
	CreatedAt      time.Time       `db:"created_at,omitempty"` //创建时间
	UpdatedAt      time.Time       `db:"updated_at,omitempty"` //更新时间
}
