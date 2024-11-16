package main
import "fmt"
import "agent/server"
func main(){
  doSomeThing()
}
/*
  Add functionality to debug by writing log and sending error log
  to cerntrl server.
*/
func  doSomeThing(){
  fmt.Println("Starting Listener...")
  server.StartServer(8443,"ca.pem","key.pem","cert.pem");
}
