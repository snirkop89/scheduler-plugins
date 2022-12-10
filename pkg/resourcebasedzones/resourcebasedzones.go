package resourcebasedzones

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "ZoneResource"

// ZoneResource is a pluging that schedule pods in a zone based on resource
type ZoneResource struct {
	frameworkHandler framework.Handle
}

var _ framework.FilterPlugin = &ZoneResource{}
var _ framework.ScorePlugin = &ZoneResource{}

func New(obj runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	// args, ok := obj.(*config.ZoneResourceArgs)
	// if !ok {
	// 	return nil, fmt.Errorf("want args to be of type ZoneResourceArgs, got %T", obj)
	// }
	return &ZoneResource{frameworkHandler: handle}, nil
}

func (zr *ZoneResource) Name() string {
	return Name
}

func (zr *ZoneResource) Filter(ctx context.Context, state *framework.CycleState, pod *corev1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.V(4).Infof("zone resource: pod: %s is trying to fit on node %s", pod.Name, nodeInfo.Node().Name)

	totalReqCPU := nodeInfo.Requested.ScalarResources["cpu"]
	klog.V(4).Info("currently requested on the node: %d", totalReqCPU)
	klog.V(4).Info("Remaining CPU on the node: %d", nodeInfo.Node().Status.Allocatable.Cpu().Value()-totalReqCPU)

	return nil
}

func (zr *ZoneResource) Score(ctx context.Context, state *framework.CycleState, p *corev1.Pod, nodeName string) (int64, *framework.Status) {
	klog.V(4).Info("Zone resource got score request")
	return 0, nil
}

func (zr *ZoneResource) ScoreExtensions() framework.ScoreExtensions {
	return nil
}
