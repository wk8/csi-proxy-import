package csi_proxy_import

import (
	corev1 "k8s.io/api/core/v1"

	v1alpha1client "github.com/wk8/csi-proxy/client/filesystem/v1alpha1"
	v1alpha1pb "github.com/wk8/csi-proxy/api/filesystem/v1alpha1"
)

func main() {
	v1alpha1client.Wkpo()
	v1alpha1pb.Wkpo()
	corev1.Wkpo()
}
