package webtest

import (
	"os"
	"strings"
	"regexp"
)

var (
	envReg = regexp.MustCompile(`^(NVIDIA|HOUDINI)`)
)

func getCntName() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	cntName := os.Getenv("CONTAINER_NAME")
	if cntName == "" {
		if err == nil {
			cntName = hostname
		}
		cntName = "unkown"
	}
	return cntName
}

func getTaskSlot() string {
	taskSlot := os.Getenv("TASK_SLOT")
	if taskSlot == "" {
		taskSlot = "unkown"
	}
	return taskSlot
}

func getSrvName() string {
	srvName := os.Getenv("SERVICE_NAME")
	if srvName == "" {
		srvName = "unkown"
	}
	return srvName
}

func getHoudiniEnv() string {
	res := []string{}
	for _, e := range os.Environ() {
		if envReg.MatchString(e) {
			res = append(res, e)
		}
	}
	return strings.Join(res, ",")
}


