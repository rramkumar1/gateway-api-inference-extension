package backend

import (
	"context"

	"inference.networking.x-k8s.io/llm-instance-gateway/api/v1alpha1"
	klog "k8s.io/klog/v2"
)

type FakePodMetricsClient struct {
	Err map[Pod]error
	Res map[Pod]*PodMetrics
}

func (f *FakePodMetricsClient) FetchMetrics(ctx context.Context, pod Pod, existing *PodMetrics) (*PodMetrics, error) {
	if err, ok := f.Err[pod]; ok {
		return nil, err
	}
	klog.V(1).Infof("pod: %+v\n existing: %+v \n new: %+v \n", pod, existing, f.Res[pod])
	return f.Res[pod], nil
}

type FakeDataStore struct {
	Res map[string]*v1alpha1.Model
}

func (fds *FakeDataStore) FetchModelData(modelName string) (returnModel *v1alpha1.Model) {
	return fds.Res[modelName]
}
