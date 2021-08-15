package product_controller

import (
	"fmt"
	"gingonic/core"
	product_query "gingonic/querys/product"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

//product response
var filePath string = "https://%s.s3-%s.amazonaws.com/%s"

func CreateProduct_c(c *gin.Context){
	
	var productBody core.Product_R

	if err:=c.ShouldBind(&productBody);err != nil{
		fmt.Println(productBody)
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Data invalid",
		})
		return
	}

	//aws upload image
	sess := c.MustGet("sess").(*session.Session)

	uploader := s3manager.NewUploader(sess)

	MyBucket := os.Getenv("BUCKET_NAME")
	//MyRegion := os.Getenv("AWS_REGION")	

	file,header, err := c.Request.FormFile("imageProduct")
	filename := header.Filename

	//upload s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL: aws.String("public-read"),
		Key: aws.String(filename),
		Body: file,
	})	

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Failed to upload the file",
			"uploader":up,			
		})
		return
	}	

	productBody.Image = up.Location


	ok, err := product_query.CreateProduct(productBody)
	
	if !ok {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"The data is invalid",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Something happens so we can't procces de request",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":"created succesfully",
	})
	return
}