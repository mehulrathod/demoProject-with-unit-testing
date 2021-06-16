package Routers

import (
	"github.com/gin-gonic/gin"
	helper "vipulsirdemo/apiHelpers"
	v1con "vipulsirdemo/controllers/api/v1"
	"vipulsirdemo/models"
	v1repo "vipulsirdemo/repository/api/v1"
	service "vipulsirdemo/services/api/v1"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	//Giving access to public folder
	r.Static("/public", "public")

	// for header access
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	uc := user()
	hc := hobby()
	pc := product()
	cc := category()
	mc := menu()
	omc := orderMenu()
	oc := order()

	v1 := r.Group("/api/v1")
	v1.POST("signup", uc.SignUp)
	v1.POST("login", uc.Login)
	v1.POST("hobby", hc.AddHobby)
	v1.POST("product", pc.AddProduct)

	//Authorized routes
	v1.Use(helper.AuthHandler())
	v1.POST("edit-profile", uc.EditProfile)
	//restaurant demo routers
	v1.POST("add-category", cc.AddCategory)
	v1.POST("edit-category", cc.EditCategory)
	v1.GET("categories", cc.GetAllCategories)
	v1.POST("add-menu", mc.AddMenu)
	v1.POST("edit-menu", mc.EditMenu)
	v1.GET("menu", mc.GetAllMenu)
	v1.POST("add-order-menu", omc.AddOrderMenu)
	v1.POST("order-menu", omc.GetOrderMenu)
	v1.POST("add-order", oc.AddOrder)
	v1.POST("user-orders", oc.GetOrders)

	return r
}

func user() *v1con.UserController {
	repo := v1repo.UserRepo{
		DB: models.GetDB(),
	}
	us := service.UserService{
		User:     models.User{},
		UserRepo: repo,
	}
	uc := v1con.UserController{
		UserService: us,
	}
	return &uc
}

func hobby() *v1con.HobbiesController {
	repo := v1repo.HobbyRepo{
		DB: models.GetDB(),
	}
	hs := service.HobbiesService{
		Hobby:     models.Hobbies{},
		HobbyRepo: repo,
	}
	hc := v1con.HobbiesController{
		HobbiesService: hs,
	}
	return &hc
}

func product() *v1con.ProductController {
	repo := v1repo.ProductRepo{
		DB: models.GetDB(),
	}
	ps := service.ProductService{
		Product:     models.Products{},
		ProductRepo: repo,
	}
	pc := v1con.ProductController{
		ProductService: ps,
	}
	return &pc
}

func category() *v1con.CategoryController {
	repo := v1repo.CategoryRepo{
		DB: models.GetDB(),
	}
	cs := service.CategoryService{
		Category:     models.Category{},
		CategoryRepo: repo,
	}
	cc := v1con.CategoryController{
		CategoryService: cs,
	}
	return &cc
}

func menu() *v1con.MenuController {
	repo := v1repo.MenuRepo{
		DB: models.GetDB(),
	}
	ms := service.MenuService{
		Menu:     models.Menu{},
		MenuRepo: repo,
	}
	mc := v1con.MenuController{
		MenuService: ms,
	}
	return &mc
}

func order() *v1con.OrderController {
	repo := v1repo.OrderRepo{
		DB: models.GetDB(),
	}
	ors := service.OrderService{
		Order:     models.Order{},
		OrderRepo: repo,
	}
	oc := v1con.OrderController{
		OrderService: ors,
	}
	return &oc
}

func orderMenu() *v1con.OrderMenuController {
	repo := v1repo.OrderMenuRepo{
		DB: models.GetDB(),
	}
	oms := service.OrderMenuService{
		OrderMenu:     models.OrderMenu{},
		OrderMenuRepo: repo,
	}
	omc := v1con.OrderMenuController{
		OrderMenuService: oms,
	}
	return &omc
}
