package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	pb "github.com/CleanCoders9229/POC-Microservices/Services/proto/UserSchema"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	ginPortDefault  = ":8000"
	grpcAddrDefault = "localhost:50051"
)

type User struct {
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Fullname string `form:"fullname" json:"fullname" xml:"fullname"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func userRegister(c *gin.Context, conn *grpc.ClientConn) {
	var user User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// gRPC Req and Res
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Profile{
		Fullname:    user.Fullname,
		Password:    user.Password,
		Email:       user.Email,
		IsActivated: false,
		CreatedDate: false,
	}
	manager := pb.NewRegistrationClient(conn)

	// Firebase Auth - Create New User
	res, err := manager.CreateNewUser(ctx, req)

	if err != nil {
		log.Printf("Response Error from Firebase: %v", err)
		c.JSON(400, gin.H{})
	}
	log.Printf("Server response: %s", res.String())

	// Gin Response
	c.JSON(200, gin.H{
		"status":   "posted",
		"email":    user.Email,
		"fullname": user.Fullname,
		"password": user.Password,
	})

}

func main() {
	// Flags Parser
	grpcAddr := flag.String("grpcAddr", grpcAddrDefault, "gRPC Address and Port.")
	ginPort := flag.String("ginPort", ginPortDefault, "Gin Server Port.")
	flag.Parse()

	// gRPC
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("gRPC connection error: %v", err)
	}

	log.Printf("gRPC Open to: %v", *grpcAddr)
	defer conn.Close()

	// GIN Server
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/user/register", func(c *gin.Context) {
		userRegister(c, conn)
	})
	router.Run(*ginPort)
}
