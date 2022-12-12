package opration

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createDeployment() {

	createdeployment := &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "app/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: deployemntName,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					labelkey: labelvalue,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						labelkey: labelvalue,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "pod",
							Image: image,
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{
									ContainerPort: port,
								},
							},
						},
					},
				},
			},
		},
	}

	deployment, err := clientset.AppsV1().Deployments("default").Create(context.Background(), createdeployment, metav1.CreateOptions{})
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("deployment", deployment.Name, "created")
}

func readDeployment() {
	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	format := "%v\t%v\n"

	fmt.Fprintf(writer, format, "NAME", "READY")
	for _, deployment := range deployments.Items {
		status := fmt.Sprintf("%v/%v", deployment.Status.AvailableReplicas, deployment.Status.Replicas)

		fmt.Fprintf(writer, format, deployment.Name, status)
	}

	writer.Flush()
}

func updateDeployment() {
	createdeployment := &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "app/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: deployemntName,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &updatereplicase,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					labelkey: labelvalue,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						labelkey: labelvalue,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "pod",
							Image: image,
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{
									ContainerPort: port,
								},
							},
						},
					},
				},
			},
		},
	}

	deployment, err := clientset.AppsV1().Deployments("default").Update(context.Background(), createdeployment, metav1.UpdateOptions{})
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("deployment", deployment.Name, "updated")
}

func deleteDeployment() {

	err := clientset.AppsV1().Deployments("default").Delete(context.Background(), deployemntName, metav1.DeleteOptions{})
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("deployment", "name", "deleted")
}
