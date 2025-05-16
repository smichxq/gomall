package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// 迁移时指定表名
func (p Product) TableName() string {
	return "product"
}

func Create(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}

func (query ProductQuery) QeryById(productId uint32) (product Product, err error) {
	// 使用context方便后续扩展(链路追踪)
	// Model指定查询表
	// First第二个参数通过主键查询
	err = query.db.WithContext(query.ctx).Model(&Product{}).First(&product, productId).Error
	return
}

func (query ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	// 将放q到两个占位符方便查询
	err = query.db.WithContext(query.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?",

		"%"+q+"%", "%"+q+"%").Error

	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

type CachedProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CachedProductQuery) GetById(productId uint32) (product Product, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cacheKey)

	err = func() error {
		err1 := cachedResult.Err()
		if err1 != nil {
			return err1
		}
		cachedResultByte, err2 := cachedResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &product)
		if err3 != nil {
			return err3
		}
		return nil
	}()
	if err != nil {
		product, err = c.productQuery.QeryById(productId)
		if err != nil {
			return Product{}, err
		}
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = c.cacheClient.Set(c.productQuery.ctx, cacheKey, encoded, time.Hour)
	}
	return
}

func NewCachedProductQuery(pq ProductQuery, cacheClient *redis.Client) CachedProductQuery {
	return CachedProductQuery{productQuery: pq, cacheClient: cacheClient, prefix: "cloudwego_shop"}
}

func GetProductById(db *gorm.DB, ctx context.Context, productId int) (product Product, err error) {
	err = db.WithContext(ctx).Model(&Product{}).Where(&Product{Base: Base{ID: productId}}).First(&product).Error
	return product, err
}

func SearchProduct(db *gorm.DB, ctx context.Context, q string) (product []*Product, err error) {
	err = db.WithContext(ctx).Model(&Product{}).Find(&product, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return product, err
}
