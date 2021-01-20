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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type podTemplate struct {
	k8s.CommonNamespacedFields
	k8s.CommonPodFields
}

// PodTemplateColumns returns kubernetes pod template fields as Osquery table columns.
func PodTemplateColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&podTemplate{})
}

// PodTemplatesGenerate generates the kubernetes pod templates as Osquery table data.
func PodTemplatesGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		pts, err := k8s.GetClient().CoreV1().PodTemplates(metav1.NamespaceAll).List(context.TODO(), options)
		if err != nil {
			return nil, err
		}

		for _, pt := range pts.Items {
			item := &podTemplate{
				CommonNamespacedFields: k8s.GetCommonNamespacedFields(pt.ObjectMeta),
				CommonPodFields:        k8s.GetCommonPodFields(pt.Template.Spec),
			}
			results = append(results, k8s.ToMap(item))
		}

		if pts.Continue == "" {
			break
		}
		options.Continue = pts.Continue
	}

	return results, nil
}
