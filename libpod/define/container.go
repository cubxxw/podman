package define

import (
	"fmt"
)

// Valid restart policy types.
const (
	// RestartPolicyNone indicates that no restart policy has been requested
	// by a container.
	RestartPolicyNone = ""
	// RestartPolicyNo is identical in function to RestartPolicyNone.
	RestartPolicyNo = "no"
	// RestartPolicyAlways unconditionally restarts the container.
	RestartPolicyAlways = "always"
	// RestartPolicyOnFailure restarts the container on non-0 exit code,
	// with an optional maximum number of retries.
	RestartPolicyOnFailure = "on-failure"
	// RestartPolicyUnlessStopped unconditionally restarts unless stopped
	// by the user. It is identical to Always except with respect to
	// handling of system restart, which Podman does not yet support.
	RestartPolicyUnlessStopped = "unless-stopped"
)

// RestartPolicyMap maps between restart-policy valid values to restart policy types
var RestartPolicyMap = map[string]string{
	"none":                     RestartPolicyNone,
	RestartPolicyNo:            RestartPolicyNo,
	RestartPolicyAlways:        RestartPolicyAlways,
	RestartPolicyOnFailure:     RestartPolicyOnFailure,
	RestartPolicyUnlessStopped: RestartPolicyUnlessStopped,
}

// Validate that the given string is a valid restart policy.
func ValidateRestartPolicy(policy string) error {
	switch policy {
	case RestartPolicyNone, RestartPolicyNo, RestartPolicyOnFailure, RestartPolicyAlways, RestartPolicyUnlessStopped:
		return nil
	default:
		return fmt.Errorf("%q is not a valid restart policy: %w", policy, ErrInvalidArg)
	}
}

// InitContainerTypes
const (
	// AlwaysInitContainer is an init container that runs on each
	// pod start (including restart)
	AlwaysInitContainer = "always"
	// OneShotInitContainer is a container that only runs as init once
	// and is then deleted.
	OneShotInitContainer = "once"
	// ContainerInitPath is the default path of the mounted container init.
	ContainerInitPath = "/run/podman-init"
)

// Kubernetes Kinds
const (
	// A Pod kube yaml spec
	K8sKindPod = "pod"
	// A Deployment kube yaml spec
	K8sKindDeployment = "deployment"
	// A DaemonSet kube yaml spec
	K8sKindDaemonSet = "daemonset"
	// a Job kube yaml spec
	K8sKindJob = "job"
)
