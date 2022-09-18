# R4Rooms - Room Booking Web Application

#### Frameworks Used;

Gin-Gonic : The whole project is done completely with gin-gonic framework which is a most popular framework in golang with more than 60k stars in github.

Install Gin-Gonic

```
go get -u github.com/gin-gonic/gin
```

Import in your code

```
import "github.com/gin-gonic/gin"
```

My main file

```
package main
import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/gingonic"
	"github.com/VJ-Vijay77/r4room/pkg/routes"
)
func main() {
	Engine := gingonic.GinInit()
	database.GetDb()
	routes.InitRoutes(Engine)
	Engine.Run(":8080")
}

```

#### Database Used;

PostgreSQL: The project is done completely with postgresql with the help of an ORM Tool called **Gorm**. Gorm is a simple framework which enables to do queries in a much simpler form. But using goring is little bit slow actually , the base sql packages are the fastest like database/sql,sqlx,sqlc etc.

To install Gorm

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

To connect to PostgreSQL

```
import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

#### Features;

##### User Side

* User can search,filter,sort and book rooms
* User can prebook rooms
* User can cancel a booking and get the half price refund on wallet
* Payment Integration is done with Razorpay and COD
* Order history and checkout details
* Password and User profile management
* Profile picture updation facility
* User can apply categorised coupons when booking

#####  Admin Side

* User management like user creation,updation,deletion and blocking
* Add rooms ,pictures and manage rooms like delete,edit and add rooms
* Adding coupons and categorised offers
* Adding Room Category
* Sales report generation
