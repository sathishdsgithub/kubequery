/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package apps

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/fake"
)

func getDeployment(t *testing.T) *v1.Deployment {
	data, err := ioutil.ReadFile("deployment_test.json")
	assert.Nil(t, err)

	d := &v1.Deployment{}
	err = json.Unmarshal(data, d)
	assert.Nil(t, err)

	return d
}

func TestDeploymentsGenerate(t *testing.T) {
	k8s.SetClient(fake.NewSimpleClientset(getDeployment(t)), types.UID("blah"))
	ds, err := DeploymentsGenerate(context.TODO(), table.QueryContext{})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]string{
		{
			"available_replicas":   "0",
			"cluster_uid":          "blah",
			"creation_timestamp":   "-62135596800",
			"host_ipc":             "0",
			"host_network":         "0",
			"host_pid":             "0",
			"min_ready_seconds":    "0",
			"observed_generation":  "0",
			"paused":               "0",
			"ready_replicas":       "0",
			"replicas":             "0",
			"strategy":             "{}",
			"unavailable_replicas": "0",
			"updated_replicas":     "0",
		},
	}, ds)
}

func TestDeploymentContainersGenerate(t *testing.T) {
	k8s.SetClient(fake.NewSimpleClientset(getDeployment(t)), types.UID("blah"))
	ds, err := DeploymentContainersGenerate(context.TODO(), table.QueryContext{})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]string{}, ds)
}

func TestDeploymentVolumesGenerate(t *testing.T) {
	k8s.SetClient(fake.NewSimpleClientset(getDeployment(t)), types.UID("blah"))
	ds, err := DeploymentVolumesGenerate(context.TODO(), table.QueryContext{})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]string{}, ds)
}
