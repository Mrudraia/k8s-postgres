package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/k8s-postgres/db"

	_ "github.com/lib/pq"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var ns string
	flag.StringVar(&ns, "namespace", "", "namespace")
	db := db.Open()

	rows, err := db.Query("SELECT * FROM test")
	fmt.Println(rows)
	if err != nil {
		log.Fatal(err)
	}
	kubeconfig := filepath.Join("kubeconfig.txt")
	fmt.Println(kubeconfig)
	log.Println("Using kubeconfig file: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})

	if err != nil {
		log.Fatalln("Failed to get pods:", err)
	}

	services, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})

	if err != nil {
		log.Fatalln("Failed to get services:", err)
	}

	http.HandleFunc("/pods", func(w http.ResponseWriter, r *http.Request) {
		for i, pod := range pods.Items {
			fmt.Fprintf(w, "[%d] %s\n", i, pod.GetName())
		}
	})

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		for i, ser := range services.Items {
			fmt.Fprintf(w, "[%d] %s\n", i, ser.GetName())
		}
	})

	http.ListenAndServe(":9090", nil)

}
