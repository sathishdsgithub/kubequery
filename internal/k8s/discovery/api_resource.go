/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package discovery

import (
	"context"

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type apiResource struct {
	metav1.APIResource
	GroupVersion string
}

// APIResourceColumns TODO
func APIResourceColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&apiResource{})
}

// APIResourcesGenerate TODO
func APIResourcesGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	results := make([]map[string]string, 0)

	sr, err := k8s.GetClient().Discovery().ServerResources()
	if err != nil {
		return nil, err
	}

	for _, rl := range sr {
		for _, r := range rl.APIResources {
			item := &apiResource{
				GroupVersion: rl.GroupVersion,
				APIResource:  r,
			}
			results = append(results, k8s.ToMap(item))
		}
	}

	return results, nil
}
