package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"strconv"
	"agent/handler"
)

//type Message struct {
//	Id   int64  `json:"id"`
//	Name string `json:"name"`
//  Data string `json:"data"`
//
//}



func main() {
	StartServer(8443,"ca.pem","key.pem","cert.pem")
}

func StartServer(port int,ca string,key string,cert string){
  /*mux := http.NewServeMux()
  mux.HandleFunc("/cmd", Cleaner)
  address := ":8000"
  log.Println("Starting server on address", address)
  err:=http.ListenAndServe(address, mux)

  //http.HandleFunc("/", Cleaner)
  //err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}*/
	// Set up a /hello resource handler
	http.HandleFunc("/cmd", handler.CmdHandler)
	http.HandleFunc("/hello", handler.HelloHandler)

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	// Create a Server instance to listen on port 8443 with the TLS config
	server := &http.Server{
		Addr:      ":"+strconv.Itoa(port),
		TLSConfig: tlsConfig,
	}
	// Listen to HTTPS connections with the server certificate and wait
	log.Fatal(server.ListenAndServeTLS(cert,key))

}
