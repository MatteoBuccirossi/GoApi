package main 

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strconv"
	"log"
)

func main() {
	app:= fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Connected")
		return c.SendString("hello world fiber")
	})
	app.Get("/api/:type", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("type"))
		obj:= make(map[string]string)
		obj["type"] = c.Params("type")
		return c.JSON(obj)
	})

	app.Get("/api/fib/:n/:steps", func(c *fiber.Ctx) error{
		n, err:= strconv.Atoi(c.Params("n"))
		steps, err:= strconv.Atoi(c.Params("steps"))
		if err == nil{
		  fib(n, steps)
		  return c.SendString("fibonacci sequence")
		}else{
			
			log.Fatal(err)
			return c.SendString("invalid input")
		}
	})

	app.Listen(":3000")
}

func fib(n int, steps int){
	mem:= n
	mem2:= n
	for j:= 0; j <= steps; j++{
		mem = mem + mem2 
		mem2 = mem - mem2
		fmt.Println(mem)
	}

}