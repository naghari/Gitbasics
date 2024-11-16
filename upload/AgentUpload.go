package upload
/*
*    Use this to Upload Files to a Server
*    Need to encrypt the file
*   Need to use mutual auth TLS.
*/
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"os"
	"path/filepath"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string,aeskey string) (*http.Response, error) {
	//file, err := os.Open(path)
	//if err != nil {
	//	return nil, err
	//}
	//defer file.Close()
  /* Load File content as dat*/
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		 log.Fatal(err)
	}
	/*  Encrypt the file with AES encryption*/
	body1:= encrypt(aeskey, dat)
	/* creating empty buffer */
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	//_, err = io.Copy(part, file)
	_, err = io.WriteString(part, body1)

  /* Adding fields to form data*/
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	//req, err := http.NewRequest("POST", uri, body)
	//req.Header.Set("Content-Type", writer.FormDataContentType())
	//return req, err
   //use TLS
	 // Read the key pair to create certificate
 	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
 	if err != nil {
 		log.Fatal(err)
 	}

 	// Create a CA certificate pool and add cert.pem to it
 	caCert, err := ioutil.ReadFile("ca.pem")
 	if err != nil {
 		log.Fatal(err)
 	}
 	caCertPool := x509.NewCertPool()
 	caCertPool.AppendCertsFromPEM(caCert)

 	// Create a HTTPS client and supply the created CA pool and certificate
 	client := &http.Client{
 		Transport: &http.Transport{
 			TLSClientConfig: &tls.Config{
 				RootCAs: caCertPool,
 				Certificates: []tls.Certificate{cert},
 			},
 		},
 	}
///

   r, err := client.Post(uri, writer.FormDataContentType(),body)
  //req.Header.Set("Content-Type", writer.FormDataContentType())
   //return req, err

 	// Request /hello via the created HTTPS client over port 8443 via GET
 	//r, err := client.Get("https://localhost:8443/hello")
 	if err != nil {
 		log.Fatal(err)
 	}
	fmt.Println(err)
 return r,err
 	// Read the response body
 /*	defer r.Body.Close()
 	body, err2 := ioutil.ReadAll(r.Body)
 	if err2 != nil {
 		log.Fatal(err2)
 	}

 	// Print the response body to stdout
 	fmt.Printf("%s\n", body)

 return
 */

}


func SendFile(aeskey string) {
	path, _ := os.Getwd()
	path += "/test.txt"
	extraParams := map[string]string{
		"title":       "EJ File",
		"author":      "Atm 0001",
		"description": "A document with all the Go programming language secrets",
	}
  resp,err	:= newfileUploadRequest("https://localhost:9443/upload", extraParams, "file", "./test.txt",aeskey)
	/*request, err := newfileUploadRequest("http://127.0.0.1:8080/upload", extraParams, "file", "./test.txt")*/
	//if err != nil {
	//	log.Fatal(err)
	//}
	//client := &http.Client{}
	//resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
	}
  resp.Body.Close()
	//	fmt.Println(resp.StatusCode)
	//	fmt.Println(resp.Header)
  //  fmt.Println(body)
	}
}
