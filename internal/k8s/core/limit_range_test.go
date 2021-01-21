/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package core

import (
	"context"
	"testing"

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/fake"
)

func TestLimitRangesGenerate(t *testing.T) {
	k8s.SetClient(fake.NewSimpleClientset(&v1.LimitRange{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "lr1",
			Namespace: "n123",
			UID:       types.UID("1234"),
			Labels:    map[string]string{"a": "b"},
		},
		Spec: v1.LimitRangeSpec{
			Limits: []v1.LimitRangeItem{
				{
					Type:                 v1.LimitTypeContainer,
					Max:                  v1.ResourceList{v1.ResourceCPU: resource.MustParse("0")},
					Min:                  v1.ResourceList{v1.ResourceCPU: resource.MustParse("4")},
					Default:              v1.ResourceList{v1.ResourceCPU: resource.MustParse("3")},
					DefaultRequest:       v1.ResourceList{v1.ResourceCPU: resource.MustParse("2")},
					MaxLimitRequestRatio: v1.ResourceList{v1.ResourceCPU: resource.MustParse("1")},
				},
			},
		},
	}), types.UID("hello"))

	js, err := LimitRangesGenerate(context.TODO(), table.QueryContext{})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]string{
		{
			"cluster_uid":             "hello",
			"creation_timestamp":      "0",
			"default":                 "{\"cpu\":\"3\"}",
			"default_request":         "{\"cpu\":\"2\"}",
			"labels":                  "{\"a\":\"b\"}",
			"max":                     "{\"cpu\":\"0\"}",
			"max_limit_request_ratio": "{\"cpu\":\"1\"}",
			"min":                     "{\"cpu\":\"4\"}",
			"name":                    "lr1",
			"namespace":               "n123",
			"type":                    "Container",
			"uid":                     "1234",
		},
	}, js)
}
