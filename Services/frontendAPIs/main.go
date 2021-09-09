package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	pb "github.com/CleanCoders9229/POC-Microservices/Services/proto/UserSchema"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Printf("Error at Hasing password: %v", err)
	}
	return string(bytes)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
		"email":    res.Email,
		"fullname": res.Fullname,
		"password": res.Password,
		"nextURL":  "/SingleLoginPage",
	})

}

func userLogin(c *gin.Context, conn *grpc.ClientConn) {
	type Token struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"username" json:"password" binding:"required"`
	}

	var inToken Token

	if err := c.ShouldBind(&inToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token := &pb.Token{
		Username: inToken.Username,
		Password: HashPassword(inToken.Password),
	}

	manager := pb.NewRegistrationClient(conn)

	res, err := manager.Login(ctx, token)

	if err != nil {
		log.Printf("Service Response Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Printf("Service Response: %v", res.String())
	c.JSON(http.StatusOK, gin.H{
		"status":  "login success",
		"nextURL": "/profile",
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

	router.POST("/user/login", func(c *gin.Context) {
		userLogin(c, conn)
	})

	router.Run(*ginPort)
}
