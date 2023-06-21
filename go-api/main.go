package main

import (
	contactpb "cloudrun-grpc/go-contact/proto"
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpcMetadata "google.golang.org/grpc/metadata"
)

const contactServerHost = "go-contact-nouqicixbq-as.a.run.app"

type ContactServer struct {
	contactpb.UnimplementedContactServer
}

func main() {
	r := gin.Default()
	r.GET("/contact/:phone", getContactHandler)
	r.GET("/contact-auth/:phone", getContactAuthHandler)
	r.Run(":8081")
}

func getContactHandler(c *gin.Context) {
	contactReply, err := getContact(c.Param("phone"), context.Background())
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, err)
	} else {
		log.Printf("%v", contactReply)
		c.JSON(http.StatusOK, contactReply)
	}
}

func getContactAuthHandler(c *gin.Context) {

	ctx, err := getAuthContext()
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	contactReply, err := getContact(c.Param("phone"), ctx)
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, err)
	} else {
		log.Printf("%v", contactReply)
		c.JSON(http.StatusOK, contactReply)
	}
}

func getAuthContext() (context.Context, error) {
	ctx := context.Background()
	tokenSource, err := idtoken.NewTokenSource(ctx, "https://"+contactServerHost)
	if err != nil {
		return nil, err
	}
	token, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}
	return grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token.AccessToken), nil
}

func getContact(phone string, ctx context.Context) (*contactpb.ContactReply, error) {
	addr := contactServerHost + ":443"

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	// Initialize client connections outside handler in your implementation
	g, err := grpc.Dial(addr, grpc.WithAuthority(contactServerHost), grpc.WithTransportCredentials(cred))
	if err != nil {
		return nil, err
	}
	defer g.Close()

	cc := contactpb.NewContactClient(g)

	contactRequest := &contactpb.ContactRequest{
		PhoneNumber: phone,
	}
	contactReply, err := cc.GetContact(ctx, contactRequest)
	if err != nil {
		return nil, err
	}
	return contactReply, nil
}
