package main

import (
	"restapi/database"
	"restapi/database/migrations"
	"restapi/routes"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	app := fiber.New()

	//inisialisasi database
	database.DatabaseInit()

	//inisialisasi migrasi
	migrations.Migration()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{
	// 		"message": "Hello, World!",
	// 	})
	// })

	//inisialisasi Rute
	routes.RouterApp(app)

	app.Post("/login", login)

	// Unauthenticated route
	app.Get("/", accessible)

	// Restricted Routes
	app.Get("/restricted", restricted)

	//JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQSflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")},
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	// end jwt middleware

	app.Listen(":8080")
}

// login jwt
func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	//cek user dan pass
	if user != "admin" || pass != "password" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	//create the claims
	claims := jwt.MapClaims{
		"user":  "Mujab admin",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//generate encoded token and send it as response
	t, err := token.SignedString([]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQSflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

// accessible jwt
func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

// restricted jwt
func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["user"].(string)
	return c.SendString("Welcome" + name)
}
