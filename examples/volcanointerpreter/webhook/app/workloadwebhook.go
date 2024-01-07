/*
Copyright 2021 The Karmada Authors.

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

package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	jobv1alpha1 "github.com/karmada-io/karmada/examples/volcanointerpreter/apis/job/batch/v1alpha1"
	configv1alpha1 "github.com/karmada-io/karmada/pkg/apis/config/v1alpha1"
	"github.com/karmada-io/karmada/pkg/webhook/interpreter"
	"k8s.io/klog/v2"
)

// Check if our volcanoInterpreter implements necessary interface
var _ interpreter.Handler = &volcanoInterpreter{}
var _ interpreter.DecoderInjector = &volcanoInterpreter{}

// volcanoInterpreter explore resource with request operation.
type volcanoInterpreter struct {
	decoder *interpreter.Decoder
}

// Handle implements interpreter.Handler interface.
// It yields a response to an ExploreRequest.
func (e *volcanoInterpreter) Handle(_ context.Context, req interpreter.Request) interpreter.Response {
	Job := &jobv1alpha1.Job{}
	fmt.Println("走到这了")
	err := e.decoder.Decode(req, Job)
	if err != nil {
		return interpreter.Errored(http.StatusBadRequest, err)
	}
	klog.Infof("Explore Job(%s/%s) for request: %s", Job.GetNamespace(), Job.GetName(), req.Operation)

	switch req.Operation {
	case configv1alpha1.InterpreterOperationInterpretReplica:
		return e.responseWithExploreReplica(Job)
	case configv1alpha1.InterpreterOperationAggregateStatus:
		return e.responseWithExploreAggregateStatus(Job, req)
	default:
		return interpreter.Errored(http.StatusBadRequest, fmt.Errorf("wrong request operation type: %s", req.Operation))
	}
}

// InjectDecoder implements interpreter.DecoderInjector interface.
func (e *volcanoInterpreter) InjectDecoder(d *interpreter.Decoder) {
	e.decoder = d
}

func (e *volcanoInterpreter) responseWithExploreReplica(Job *jobv1alpha1.Job) interpreter.Response {
	res := interpreter.Succeeded("")
	tasksLength := int32(len(Job.Spec.Tasks))
	res.Replicas = &tasksLength
	return res
}

func (e *volcanoInterpreter) responseWithExploreAggregateStatus(Job *jobv1alpha1.Job, req interpreter.Request) interpreter.Response {
	wantedJob := Job.DeepCopy()
	var readyReplicas int32
	for _, item := range req.AggregatedStatus {
		if item.Status == nil {
			continue
		}
		status := &jobv1alpha1.JobStatus{}
		if err := json.Unmarshal(item.Status.Raw, status); err != nil {
			return interpreter.Errored(http.StatusInternalServerError, err)
		}
		readyReplicas += status.Running
	}
	wantedJob.Status.Running = readyReplicas
	marshaledBytes, err := json.Marshal(wantedJob)
	if err != nil {
		return interpreter.Errored(http.StatusInternalServerError, err)
	}
	return interpreter.PatchResponseFromRaw(req.Object.Raw, marshaledBytes)
}
