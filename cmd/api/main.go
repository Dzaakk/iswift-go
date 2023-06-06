package main

import (
	admin "iswift-go-project/internal/admin/injector"
	oauth "iswift-go-project/internal/oauth/injector"
	productCategory "iswift-go-project/internal/product_category/injector"
	profile "iswift-go-project/internal/profile/injector"
	product "iswift-go-project/internal/product/injector"
	register "iswift-go-project/internal/register/injector"
	cart "iswift-go-project/internal/cart/injector"
	mysql "iswift-go-project/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.DB()

	r := gin.Default()

	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	profile.InitializedService(db).Route(&r.RouterGroup)
	admin.InitializedService(db).Route(&r.RouterGroup)
	productCategory.InitializedService(db).Route(&r.RouterGroup)
	product.InitializedService(db).Route(&r.RouterGroup)
	cart.InitializedService(db).Route(&r.RouterGroup)

	r.Run()

}
