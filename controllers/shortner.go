package controllers

import (
	"github.com/AdrianKubica/shortid"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"shortner/config"
	"shortner/utils"

	"context"
	"time"
)

type Url struct {
	LongUrl string	`json:"longUrl"`
}
type DB struct {
	LongUrl string	`bson:"longUrl"`
	ShortUrl string	`bson:"shortUrl"`
	ShortCode string	`bson:"shortCode"`
	CreatedAt time.Time `bson:"createdAt"`
}
var collection = config.GetCollection("Shortner")
const BASE_URL = "http://localhost:8000/"

func Shortner (c *fiber.Ctx) error {

	data := new(Url)
	db := new(DB)
 	if err := c.BodyParser(data); err != nil {
            return err
        }
	if(!utils.IsUrl(data.LongUrl)){
			return c.JSON(fiber.Map{
			"message":"invalid url",
		})

}
		filter := bson.M{"longUrl": data.LongUrl}
		err := collection.FindOne(context.TODO(), filter).Decode(&db)

		if err != nil {

		shortCode, generr := shortid.Generate()
		if generr != nil {
			return c.JSON(fiber.Map{
				"message":"Could not generate shortid",
			})
		}

		shortUrl := BASE_URL + shortCode

		db.LongUrl = data.LongUrl
		db.ShortUrl = shortUrl
		db.ShortCode = shortCode
		db.CreatedAt = time.Now()

		collection.InsertOne(context.TODO(), db)
		return c.JSON(fiber.Map{
			"message" : "Created new shorturl",
			"data": shortUrl,
		})
	
		}

		return c.JSON(fiber.Map{
			"message" : "An existing shorturl was found",
			"data":	db.ShortUrl,
		})
	}
	


func Redirect(c *fiber.Ctx) error{
	data := new(DB)
	shortCode := c.Params("code")

	err := collection.FindOne(context.TODO(), bson.M{"shortCode": shortCode}).Decode(&data)
	if err != nil {
		return c.JSON(fiber.Map{
			"message" : "Invalid url",
			"error" : err.Error(),
		})
	}
	return c.Redirect(data.LongUrl)
}