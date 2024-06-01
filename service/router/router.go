package router

import (
	"net/http"
	"strconv"

	"wahyu/e-commerce/core"
	"wahyu/e-commerce/service/entities"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin *gin.Engine
}

type RouterContract interface {
	NewRouter() http.Handler
}

func RouterConstructor(gin *gin.Engine) RouterContract {
	return &Router{
		gin: gin,
	}
}

func (r *Router) NewRouter() http.Handler {

	r.gin.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": true,
			"msg":    "👌",
		})
	})

	r.gin.GET("/products", func(c *gin.Context) {
		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().
			Joins("LEFT JOIN shop ON product.shop_id = shop.id").
			Select(
				"product.id",
				"product.name",
				"product.description",
				"product.price",
				"product.image",
				"product.shop_id",
				"shop.name AS shopName",
			)

		var data []entities.Product

		q.Find(&data)
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	r.gin.GET("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) 

		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().
		Joins("LEFT JOIN shop ON product.shop_id = shop.id").
		Where("Product.id = ?", &id).
		Select(
			"product.id",
			"product.name",
			"product.description",
			"product.price",
			"product.image",
			"product.shop_id",
			"shop.name AS shopName",
		)

		var data []entities.Product

		q.Find(&data)
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	r.gin.POST("/products", func(c *gin.Context) {
		input := &entities.ProductInsert{}
		if err := c.ShouldBindJSON(&input); err != nil {

			e := err.Error()
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		app := core.NewApp()
		db := app.Mysql

		db.Debug().Create(&input)

		c.JSON(200, gin.H{
			"status": true,
			"result": &input,
		})
	})

	//r.gin.POST("/products/batch", func(c *gin.Context) {
	// 	input := &entities.ProductBatch{}
	// 	if err := c.ShouldBindJSON(&input); err != nil {
	// 		e := err.Error()
	// 		c.JSON(500, gin.H{
	// 			"status":     false,
	// 			"errMessage": &e,
	// 		})
	// 		return
	// 	}

	// 	app := core.NewApp()
	// 	db := app.Mysql

	// 	db.Debug().Create(&input.Data)

	// 	c.JSON(200, gin.H{
	// 		"status": true,
	// 		"result": &input.Data,
	// 	})
	// })

	r.gin.PUT("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		input := &entities.ProductInsert{}
		if err := c.ShouldBindJSON(&input); err != nil {
			e := err.Error()
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		app := core.NewApp()
		db := app.Mysql

		var entity entities.Product
		if err := db.Where("id = ?", id).First(&entity).Error; err != nil {
			e := err.Error()
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		
		updateFields := make(map[string]interface{})
		if input.Name != nil {
			updateFields["name"] = input.Name
		}
		if input.Description != nil {
			updateFields["description"] = input.Description
		}
		if input.Price != nil {
			updateFields["price"] = input.Price
		}
		if input.Image != nil {
			updateFields["image"] = input.Image
		}
		if input.ShopId != nil {
			updateFields["shop_id"] = input.ShopId
		}

		if err := db.Model(&entity).Updates(updateFields).Error; err != nil {
			e := err.Error()
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		c.JSON(200, gin.H{
			"status": true,
			"result": &entity,
		})
	})

	r.gin.DELETE("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) 

		app := core.NewApp()
		db := app.Mysql

		db.Delete(&entities.Product{}, id)

		c.JSON(200, gin.H{
			"status": true,
		})
	})

	return r.gin
}
