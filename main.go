/*
Copyright © 2020 Michael Rubin <mhr@neverthere.org>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	//	appsv1 "k8s.io/api/apps/v1"

	apiv1 "k8s.io/api/core/v1"

	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	//	"k8s.io/client-go/util/retry"

	// Adding the golangci-lint import here as a way to insure
	// that it is installed automatically with a standard version
	// when we either build the product or run tests.
	_ "github.com/golangci/golangci-lint/pkg/commands"
)

const numArgs = 2

func main() {
	var kubeconfig *string

	home := homedir.HomeDir()
	if home != "" {
		kubeconfig =
			flag.String("kubeconfig",
				filepath.Join(home, ".kube", "config"),
				"(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	if len(flag.Args()) != numArgs {
		fmt.Println("Please specify <eployment> <replicas>")
		os.Exit(1)
	}

	target := flag.Arg(0)
	replicaStr := flag.Arg(1)

	replicas, err := strconv.Atoi(replicaStr)
	if err != nil {
		fmt.Println("Replicas must be an integer")
		os.Exit(1)
	}

	fmt.Printf("Setting Deplyment %s to %d replicas.", target, replicas)

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// get and update replicas
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deployment, _ := deploymentsClient.Get(context.TODO(), target, metav1.GetOptions{})
	// TODO
	//	deploymentsClient.Update(context.TODO(), deployment, metav1.UpdateOptions{})
	fmt.Printf("%#v", deployment)
}
