package main

import (
	"RyuLdnWebsite/config"
	"RyuLdnWebsite/ldnhealthcheck"
	"RyuLdnWebsite/routes"
	"RyuLdnWebsite/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.LoadEnvVariables()

	services.InitRedis()

	ldnHost := config.GetEnv("LDN_HOST")
	ldnPort := config.GetEnv("LDN_PORT")
	ldnHeathcheckTime := config.GetEnv("LDN_HEALTHCHECK_TIME")

	if ldnHost == "" {
		log.Fatal("LDN_HOST is not set")
	}
	if ldnPort == "" {
		log.Fatal("LDN_PORT is not set")
	}
	numLdnHostPort, err := strconv.Atoi(ldnPort)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	numLdnHeathcheckTime, err := strconv.Atoi(ldnHeathcheckTime)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	ldnhealthcheck.SceduleTask(time.Duration(numLdnHeathcheckTime)*time.Second, ldnHost, numLdnHostPort, 30*time.Second)

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
