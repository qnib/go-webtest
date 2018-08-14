package webtest

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"os/exec"
	"bytes"
	"os"
	"net"
	"log"
	"io/ioutil"
)

func extractIP(req *http.Request) (string, error) {
	var err error
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	return ip, err
}

func getName(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	fmt.Fprintf(w, "container: %s\n", getCntName())
}

func getGPUs(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	cmd := exec.Command("nvidia-smi", "-L")
	cmdStdout := &bytes.Buffer{}
	cmd.Stdout = cmdStdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "%s >> No GPUs or nvidia-smi is not working correctly: %s\n", getCntName(), err.Error())
		return
	}
	fmt.Fprintf(w, "%s: %s\n", getCntName(), getHoudiniEnv())
	fmt.Fprintf(w, string(cmdStdout.Bytes()))
}

func getTask(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	fmt.Fprintf(w, "%s.%s\n", getSrvName(), getTaskSlot())
}

func getIP(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "unkown"
	}
	fmt.Fprintf(w, "You've hit cnt:%s at path:%s on pod:%s\n", getCntName(), req.URL.Path, podName)
	ip, err := extractIP(req)
	if err != nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
		return
	}
	fmt.Fprintf(w, "IP: %s\n", ip)
}

func callMeBackHelp(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	fmt.Fprintln(w, ">> Please send a request body")
	http.Error(w, `e.g.: curl -X POST -H "Content-Type: application/json" -d  '{"path":"/echo","host":"scooterlabs.com"}' 127.0.0.1:8081/callmeback`, 400)
}

func callMeBack(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	if req.Body == nil {
		callMeBackHelp(w, req, params)
		return
	}
	var cb CallBackRequest
	err := json.NewDecoder(req.Body).Decode(&cb)
	if err != nil {
		fmt.Fprintf(w, "req: %v", req.Body)
		http.Error(w, err.Error(), 400)
		return
	}
	if cb.Host == "" {
		cb.Host, _ = extractIP(req)
	}
	if cb.Port == "" {
		cb.Port = "80"
	}
	if cb.Proto == "" {
		cb.Proto = "http"
	}
	addr := fmt.Sprintf("%s://%s:%s%s", cb.Proto, cb.Host, cb.Port, cb.Path)
	log.Printf("Call back: %s", addr)
	res, err := http.Get(addr)
	if err != nil {
		fmt.Fprintf(w, "Could not call back %s > %s\n", addr, err.Error())
		return
	}
	defer res.Body.Close()
	contents, _ := ioutil.ReadAll(res.Body)
	fmt.Fprintf(w, ">>> Callback %s\n", addr)
	fmt.Fprintf(w, string(contents))
}