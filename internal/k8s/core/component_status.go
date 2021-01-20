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

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type componentStatus struct {
	Name string
	v1.ComponentCondition
}

// ComponentStatusColumns returns kubernetes component status fields as Osquery table columns.
func ComponentStatusColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&componentStatus{})
}

// ComponentStatusesGenerate generates the kubernetes component statuses as Osquery table data.
func ComponentStatusesGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		css, err := k8s.GetClient().CoreV1().ComponentStatuses().List(context.TODO(), options)
		if err != nil {
			return nil, err
		}

		for _, cs := range css.Items {
			for _, c := range cs.Conditions {
				item := &componentStatus{
					Name:               cs.Name,
					ComponentCondition: c,
				}
				results = append(results, k8s.ToMap(item))
			}
		}

		if css.Continue == "" {
			break
		}
		options.Continue = css.Continue
	}

	return results, nil
}
