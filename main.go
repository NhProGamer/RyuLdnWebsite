package main

import (
	"RyuLdnWebsite/config"
	"RyuLdnWebsite/routes"
	"RyuLdnWebsite/services"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"strings"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.LoadEnvVariables()

	services.InitRedis()

	r := gin.Default()
	r.Static("/static", "static")
	r.StaticFile("/", "public/index.html")

	proxies, err := ParseCIDRs(config.GetEnv("ALLOWED_PROXIES"))
	if err != nil {
		log.Fatal(err)
	}
	err = r.SetTrustedProxies(proxies)
	if err != nil {
		log.Fatal(err)
	}

	routes.InitRoutes(r)

	// Démarrer le serveur
	host := config.GetEnv("HOST")
	port := config.GetEnv("PORT")
	if host == "" {
		log.Fatal("HOST is not set")
	}
	if port == "" {
		log.Fatal("PORT is not set")
	}
	if err := r.Run(host + ":" + port); err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}
}

func ParseCIDRs(input string) ([]string, error) {
	var validCIDRs []string
	cidrList := strings.Split(input, ",")

	for _, cidr := range cidrList {
		cidr = strings.TrimSpace(cidr)
		_, _, err := net.ParseCIDR(cidr)
		if err == nil {
			validCIDRs = append(validCIDRs, cidr)
		}
	}

	return validCIDRs, nil
}
