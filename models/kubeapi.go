package models

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	jc "github.com/jetlwx/comm"
	k8s "k8s.io/client-go/pkg/api/v1"
)

var (
	kubeAPI = beego.AppConfig.String("kubeAPI")
)

// //get kube api namespaces list
// func PodsNamespace() (l []string) {
// 	space := k8s.NamespaceList{}
// 	url := kubeAPI + "/api/v1/namespaces"
// 	res, code, err := jc.GetJsonFromUrl(url)
// 	if code != 200 || err != nil {
// 		log.Println("http code: ", code, " at get kubeapi:", err)
// 		return l
// 	}

// 	e := json.Unmarshal(res, &space)
// 	if e != nil {
// 		log.Println("error at json.Unmarshal:", e)
// 		return l
// 	}

// 	for _, v := range space.Items {
// 		log.Println(v.ObjectMeta.Name)
// 		if v.ObjectMeta.Name != "" {
// 			l = append(l, v.ObjectMeta.Name)
// 		}
// 	}

// 	//log.Println("space=", l)
// 	return l
// }

//podlist
func Podlist(namespace string) (l []string) {
	pods := k8s.PodList{}
	url := kubeAPI + "api/v1/namespaces/" + namespace + "/pods"
	res, code, err := jc.GetJsonFromUrl(url)
	if code != 200 || err != nil {
		log.Println("http code: ", code, " at get kubeapi:", err)
		return l
	}

	e := json.Unmarshal(res, &pods)
	if e != nil {
		log.Println("error at json.Unmarshal:", e)
		return l
	}

	for _, v := range pods.Items {
		log.Println(v.ObjectMeta.Name)
		if v.ObjectMeta.Name != "" {
			l = append(l, v.ObjectMeta.Name)
		}
	}

	log.Println("pods=", l)
	return l
}

func AllPodslist() (l []string) {
	pods := k8s.PodList{}
	url := kubeAPI + "/api/v1/pods"
	res, code, err := jc.GetJsonFromUrl(url)
	if code != 200 || err != nil {
		log.Println("http code: ", code, " at get kubeapi:", err)
		return l
	}

	e := json.Unmarshal(res, &pods)
	if e != nil {
		log.Println("error at json.Unmarshal:", e)
		return l
	}

	for _, v := range pods.Items {
		log.Println(v.ObjectMeta.Name)
		if v.ObjectMeta.Name != "" {
			l = append(l, v.ObjectMeta.Name)
		}
	}

	//	log.Println("pods=", l)
	return l
}
