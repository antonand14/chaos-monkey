package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type ChaosAgent struct {
	clientset *kubernetes.Clientset
}

func NewChaosAgent(kubeconfig string) (*ChaosAgent, error) {
	var config *rest.Config
	var err error

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &ChaosAgent{clientset: clientset}, nil
}

// PodAssassin kills random pods matching criteria
func (ca *ChaosAgent) PodAssassin(namespace string, labelSelector string, percentage int) error {
	ctx := context.Background()

	pods, err := ca.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return err
	}

	if len(pods.Items) == 0 {
		return fmt.Errorf("no pods found matching selector")
	}

	numToKill := (len(pods.Items) * percentage) / 100
	if numToKill == 0 {
		numToKill = 1
	}

	log.Printf("Found %d pods, will kill %d (%d%%)", len(pods.Items), numToKill, percentage)

	// Shuffle and kill random pods
	rand.Shuffle(len(pods.Items), func(i, j int) {
		pods.Items[i], pods.Items[j] = pods.Items[j], pods.Items[i]
	})

	for i := 0; i < numToKill && i < len(pods.Items); i++ {
		pod := pods.Items[i]
		log.Printf("Killing pod: %s/%s", pod.Namespace, pod.Name)

		err := ca.clientset.CoreV1().Pods(pod.Namespace).Delete(ctx, pod.Name, metav1.DeleteOptions{})
		if err != nil {
			log.Printf("Error killing pod %s: %v", pod.Name, err)
		} else {
			log.Printf("Successfully killed pod: %s", pod.Name)
		}

		time.Sleep(1 * time.Second) // Stagger deletions
	}

	return nil
}

// MeasureRecovery checks how long it takes for pods to recover
func (ca *ChaosAgent) MeasureRecovery(namespace string, labelSelector string, expectedCount int) (time.Duration, error) {
	ctx := context.Background()
	startTime := time.Now()

	for {
		pods, err := ca.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
			LabelSelector: labelSelector,
		})
		if err != nil {
			return 0, err
		}

		readyCount := 0
		for _, pod := range pods.Items {
			for _, condition := range pod.Status.Conditions {
				if condition.Type == "Ready" && condition.Status == "True" {
					readyCount++
					break
				}
			}
		}

		if readyCount >= expectedCount {
			recoveryTime := time.Since(startTime)
			log.Printf("Recovery complete: %d/%d pods ready in %v", readyCount, expectedCount, recoveryTime)
			return recoveryTime, nil
		}

		if time.Since(startTime) > 5*time.Minute {
			return 0, fmt.Errorf("recovery timeout after 5 minutes")
		}

		time.Sleep(2 * time.Second)
	}
}

func main() {
	log.Println("Chaos Agent starting...")

	// TODO: Accept configuration from API/environment
	// TODO: Implement other chaos types (network, resource, etc.)
	// TODO: Add safety checks and dry-run mode

	log.Println("Chaos Agent ready. Waiting for experiments...")
	select {} // Keep running
}
