package utils

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"
)

func CheckerTLS(s string) string {
	conn, err := tls.Dial("tcp", s+":443", nil)
	if err != nil {
		fmt.Println("Server doesn't support SSL certificate err: " + err.Error())
		os.Exit(1)
	}

	err = conn.VerifyHostname(s)
	if err != nil {
		fmt.Println("Hostname doesn't match with certificate: " + err.Error())
		os.Exit(1)
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	fmt.Printf("Issuer: %s\nExpiry: %v\n", conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))

	return ("Status: Valid Certificate")
}
