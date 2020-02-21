package bpaas

import (
	"github.com/ysfkel/bpaas/kubernetes/kind"
)

func NewKubernetesService() k8s.IKubernetesService {

	return kind.KubernetesService{}
}
