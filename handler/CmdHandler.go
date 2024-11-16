package handler
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"agent/model"
	"agent/upload"
)

// curl localhost:8000 -d '{"name":"Hello"}'
func CmdHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg model.Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
  if(msg.Name=="upload"){
		fmt.Println("Cmd from Server is Upload and can be processed")
		//"00112233445566778899aabbccddeeff"
    upload.SendFile(msg.Data)
	}else{
		fmt.Println("Cmd from Server is "+msg.Name)
	}

	// call a function to do network
	//like callinv a fileupload
  msg.Data="Received."
	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
