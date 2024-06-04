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

		var data entities.Product

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

		var entity entities.ProductInsert
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






	r.gin.GET("/users", func(c *gin.Context) {
		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().Select("id,name,created_at")

		var data []entities.User

		q.Find(&data)
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	r.gin.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) 

		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().
		Where("user.id = ?", &id).Select("*")

		var data entities.User

		q.Find(&data)
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	r.gin.POST("/users", func(c *gin.Context) {
		input := &entities.UserInsert{}
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

	r.gin.PUT("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		input := &entities.UserInsert{}
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

		var entity entities.UserInsert
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
		if input.Password != nil {
			updateFields["password"] = input.Password
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

	r.gin.DELETE("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) 

		app := core.NewApp()
		db := app.Mysql

		db.Delete(&entities.User{}, id)

		c.JSON(200, gin.H{
			"status": true,
		})
	})



	




























	r.gin.GET("/shops", func(c *gin.Context) {
		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().
		Joins("LEFT JOIN user ON shop.user_id = user.id").
		Select(
			"shop.id",
			"shop.name",
			"shop.address",
			"shop.picture",
			"shop.user_id",
			"user.name AS userName",
		)

		var data []entities.Shop

		q.Find(&data)
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	r.gin.GET("/shops/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) 

		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().
		Joins("LEFT JOIN user ON shop.user_id = user.id").
		Where("shop.id = ?", &id).
		Select(
			"shop.id",
			"shop.name",
			"shop.address",
			"shop.picture",
			"shop.user_id",
			"user.name AS userName",
		)

		var data entities.Shop

		q.Find(&data)
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	r.gin.POST("/shops", func(c *gin.Context) {
		input := &entities.ShopInsert{}
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

	r.gin.PUT("/shops/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		input := &entities.ShopInsert{}
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

		var entity entities.ShopInsert
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
		if input.Address != nil {
			updateFields["address"] = input.Address
		}
		if input.Picture != nil {
			updateFields["picture"] = input.Picture
		}
		if input.UserId != nil {
			updateFields["user_id"] = input.UserId
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

	r.gin.DELETE("/shops/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) 

		app := core.NewApp()
		db := app.Mysql

		db.Delete(&entities.Shop{}, id)

		c.JSON(200, gin.H{
			"status": true,
		})
	})






























	
	r.gin.GET("/transactions", func(c *gin.Context) {
		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().
		Joins("LEFT JOIN product ON transaction.product_id = product.id").
		Joins("LEFT JOIN user ON transaction.user_id = user.id").
		Select(
			"transaction.id",
			"transaction.address",
			"transaction.product_id",
			"product.name as ProductName",
			"transaction.user_id",
			"user.name AS userName",
			"transaction.created_at",
		)

		var data []entities.Transaction

		q.Find(&data)
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	r.gin.GET("/transactions/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) 

		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().
		Joins("LEFT JOIN product ON transaction.product_id = product.id").
		Joins("LEFT JOIN user ON transaction.user_id = user.id").
		Where("shop.id = ?", &id).
		Select(
			"transaction.id",
			"transaction.address",
			"transaction.product_id",
			"product.name as ProductName",
			"transaction.user_id",
			"user.name AS userName",
			"transaction.created_at",
		)

		var data entities.Transaction

		q.Find(&data)
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	r.gin.POST("/transactions", func(c *gin.Context) {
		input := &entities.TransactionInsert{}
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

	

	r.gin.DELETE("/transactions/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) 

		app := core.NewApp()
		db := app.Mysql

		db.Delete(&entities.Transaction{}, id)

		c.JSON(200, gin.H{
			"status": true,
		})
	})


































	r.gin.GET("/cart", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("user_id"))

		app := core.NewApp()
		db := app.Mysql

		q := db.Debug().

		Joins("LEFT JOIN product ON cart.product_id = product.id").
		Joins("LEFT JOIN user ON cart.user_id = user.id").
				Where("cart.user_id = ?", &id).
		Select(
			"cart.id",
			"cart.amount",
			"cart.product_id",
			"product.name as productName",
			"product.image as productImage",
			"product.price as productPrice" ,
			"cart.user_id",
			"user.name AS userName",
		)

		var data []entities.Cart

		q.Find(&data)
		c.JSON(200, gin.H{
			"status": true,
			"data":   &data,
		})
	})

	r.gin.POST("/cart", func(c *gin.Context) {
		input := &entities.CartInsert{}
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

	r.gin.PUT("/cart/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		input := &entities.CartInsert{}
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

		var entity entities.CartInsert
		if err := db.Where("id = ?", id).First(&entity).Error; err != nil {
			e := err.Error()
			c.JSON(500, gin.H{
				"status":     false,
				"errMessage": &e,
			})
			return
		}

		
		updateFields := make(map[string]interface{})

		if input.UserId != nil {
			updateFields["user_id"] = input.UserId
		}
		if input.UserId != nil {
			updateFields["product_id"] = input.UserId
		}
		if input.UserId != nil {
			updateFields["amount"] = input.UserId
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

	r.gin.DELETE("/cart/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) 

		app := core.NewApp()
		db := app.Mysql

		db.Delete(&entities.Cart{}, id)

		c.JSON(200, gin.H{
			"status": true,
		})
	})

















	return r.gin
}
