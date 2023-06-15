package main

import (
	admin "iswift-go-project/internal/admin/injector"
	cart "iswift-go-project/internal/cart/injector"
	classRoom "iswift-go-project/internal/class_room/injector"
	dashboard "iswift-go-project/internal/dashboard/injector"
	discount "iswift-go-project/internal/discount/injector"
	oauth "iswift-go-project/internal/oauth/injector"
	order "iswift-go-project/internal/order/injector"
	product "iswift-go-project/internal/product/injector"
	productCategory "iswift-go-project/internal/product_category/injector"
	profile "iswift-go-project/internal/profile/injector"
	register "iswift-go-project/internal/register/injector"
	webhook "iswift-go-project/internal/webhook/injector"
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
	discount.InitializedService(db).Route(&r.RouterGroup)
	order.InitializedService(db).Route(&r.RouterGroup)
	webhook.InitializedService(db).Route(&r.RouterGroup)
	classRoom.InitializedService(db).Route(&r.RouterGroup)
	dashboard.InitializedService(db).Route(&r.RouterGroup)

	r.Run()

}
