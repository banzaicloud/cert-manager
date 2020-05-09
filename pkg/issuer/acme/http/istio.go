package http

import (
	"context"
	//"fmt"
	//
	//certmanagerv1alpha2 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	//logf "github.com/jetstack/cert-manager/pkg/logs"
	//extv1beta1 "k8s.io/api/extensions/v1beta1"
	//"k8s.io/api/networking/v1beta1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/apimachinery/pkg/types"
	//ctrl "sigs.k8s.io/controller-runtime"

	cmacme "github.com/jetstack/cert-manager/pkg/apis/acme/v1alpha2"
	logf "github.com/jetstack/cert-manager/pkg/logs"
	v1alpha3 "istio.io/api/networking/v1alpha3"
	clientv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Solver) ensureIstio(ctx context.Context, ch *cmacme.Challenge, svcName string) (i *struct{}, err error) {
	log := logf.FromContext(ctx).WithName("ensureIstio")
	gateway, err := s.IstioClient.NetworkingV1alpha3().Gateways(ch.Namespace).Create(ctx, &clientv1alpha3.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName:    "cm-acme-http-solver-",
			Namespace:       ch.Namespace,
			Labels:          nil,
			Annotations:     nil,
			OwnerReferences: []metav1.OwnerReference{*metav1.NewControllerRef(ch, challengeGvk)},
		},
		Spec: v1alpha3.Gateway{
			Servers: []*v1alpha3.Server{
				{
					Port: &v1alpha3.Port{
						Number:   12345,
						Protocol: "HTTP",
						Name:     "http",
					},
					Hosts: []string{ch.Spec.DNSName},
				},
			},
			Selector: map[string]string{
				"pitty": "putty",
			},
		},
	}, metav1.CreateOptions{})

	if err != nil {
		log.Error(err, "Sob")
		return nil, err
	}
	log.Info("Yay!", "gateway", gateway)
	return nil, nil
}
