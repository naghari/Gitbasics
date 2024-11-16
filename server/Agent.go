package server
import "fmt"
func info(){

fmt.Println("Agent will have a Server and Client")
fmt.Println("Server will receive Command(json) and executes the requests")
fmt.Println("AgentUpload will help in uploading Ejfile to server as multipart/data")
fmt.Println("File will be encrypted before uploading to server")
fmt.Println("Server can send encryption key as part of command. It can be a plain key(as TLS is used) or encrypted using ATM publickey")
fmt.Println("Server and Client communication  should use TLS mutual auth")


fmt.Println("=====================================")
fmt.Println("1. TLS Java Server    DONE")
fmt.Println("2. TLS Java Client    DONE")
fmt.Println("3. TLS Go FileUpload Client    DONE")
fmt.Println("4. TLS Go Server    DONE")
fmt.Println("5. Upgrade TLS Go Server to receive Command  DONE")

fmt.Println("6. Get encrypted Key from Server in Command and Encrypt the file and upload to server.   NOT DONE")
fmt.Println("7. Download media files from server and move it to a given folder.")








}
