package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/jetlwx/comm"
	"github.com/jetlwx/kubePodTerminal/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	pods := models.AllPodslist()
	c.Data["SPACE"] = pods

}

func (c *MainController) Sub() {
	c.TplName = "index.html"
	var cmd string
	pod := c.GetString("pod")
	subtype := c.GetString("pb")
	logkeyword := c.GetString("keyword")

	log.Println("pod=", pod)
	log.Println("pb=", subtype)
	if len(pod) == 0 {
		return
	}

	port := models.Random()
	log.Println("port=", port)
	if port == 0 {
		return
	}

	switch subtype {
	case "dockerlogs":
		cmd = "gotty -p " + strconv.Itoa(port) + " --reconnect  --once kubectl logs  " + pod
		log.Println("cmd=", cmd)
	case "webtty":
		cmd = "gotty -p " + strconv.Itoa(port) + " --reconnect  -w    -c leyoujia:jjshome520 --once kubectl exec -it " + pod + "  /bin/bash"
		log.Println("cmd=", cmd)

	case "threadnum":
		cmd = "gotty -p " + strconv.Itoa(port) + " --reconnect --once kubectl exec  " + pod + "  --  /home/logs/.tools.sh javaThreadnum"
		log.Println("cmd=", cmd)

	case "jstack":
		cmd = "gotty -p " + strconv.Itoa(port) + " --reconnect --once kubectl exec  " + pod + "  -- /home/logs/.tools.sh printJstack "
	case "catjavalogs":
		if logkeyword == "" {
			return
		}
		cmd = "gotty -p " + strconv.Itoa(port) + " --reconnect --once kubectl exec   " + pod + "  -- /home/logs/.tools.sh" + "  showLog  " + logkeyword

	case "tailjavalogs":
		if logkeyword == "" {
			return
		}
		lpath, err := comm.ExecOSCmdForBash("kubectl exec  " + pod + "  /home/logs/.tools.sh showLogpath " + logkeyword)
		if err != nil || string(lpath) == "" {
			return
		}

		cmd = "gotty -p " + strconv.Itoa(port) + " --reconnect    --once kubectl exec -it  " + pod + "  -- tail -f -n 1000  " + string(lpath)
	default:
		return
	}

	go comm.ExecOSCmdNoReturn(cmd)
	if subtype == "webtty" || subtype == "tailjavalogs" {
		go killTailProcess(port, pod, subtype)
	}

	server := beego.AppConfig.String("server")

	url := "http://" + server + ":" + strconv.Itoa(port) + "/"
	log.Println("url=", url)
	//	time.Sleep(1 * time.Second)
	c.Redirect(url, 302)

}

func killTailProcess(port int, pod string, subtype string) {
	time.Sleep(10 * time.Second)
	cmd := "ss -antlp | grep " + strconv.Itoa(port) + " | grep -v grep "
	log.Println("cmd=", cmd)

	if port == 0 {
		return
	}
	//  check the web client is disconnect
	for {
		listen, err := comm.ExecOSCmdForBash(cmd)
		time.Sleep(1 * time.Second)
		log.Println("tail client status:", string(listen))

		if err != nil {
			log.Println("error at check webclient is disconnect:", err)
			break
		}

		if len(listen) == 0 {
			log.Println("The web client disconnect.....")
			break
		}

		time.Sleep(5 * time.Second)
	}
	// if disconnect
	var cmdpid string
	switch subtype {
	case "webtty":
		cmdpid = "kubectl exec " + pod + " -- /home/logs/.tools.sh showbashPID"

	case "tailjavalogs":
		cmdpid = "kubectl exec " + pod + " -- /home/logs/.tools.sh showtailPID"

	}

	tpid, err := comm.ExecOSCmdForBash(cmdpid)
	if err != nil {
		log.Println("error at get tail cmd pid:", err)
		return
	}
	time.Sleep(1 * time.Second)

	ttpid := string(tpid)
	log.Println("tailCmdPid=", ttpid)

	if ttpid == "" {
		log.Println("tail pid is null")
		return
	}
	time.Sleep(2 * time.Second)
	killTailcmd := "kubectl exec  " + pod + " --  kill -9 " + ttpid
	comm.ExecOSCmdNoReturn(killTailcmd)
}
