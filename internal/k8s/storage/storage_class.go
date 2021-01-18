/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package storage

import (
	"context"

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type storageClass struct {
	v1.StorageClass
}

// SGClassColumns TODO
func SGClassColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&storageClass{})
}

// SGClassesGenerate TODO
func SGClassesGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		classes, err := k8s.GetClient().StorageV1().StorageClasses().List(context.TODO(), options)
		if err != nil {
			return nil, err
		}

		for _, c := range classes.Items {
			item := &storageClass{
				StorageClass: c,
			}
			results = append(results, k8s.ToMap(item))
		}

		if classes.Continue == "" {
			break
		}
		options.Continue = classes.Continue
	}

	return results, nil
}
