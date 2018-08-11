package webtest

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"os/exec"
	"bytes"
	"os"
	"net"
)

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
	ip, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
		fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
		return
	}
	fmt.Fprintf(w, "IP: %s\n", ip)
	fmt.Fprintf(w, "Port: %s\n", port)
}