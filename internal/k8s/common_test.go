/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package k8s

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetCommonFields(t *testing.T) {
	meta := metav1.ObjectMeta{}
	assert.Equal(t, GetCommonFields(meta), CommonFields{
		UID:               meta.UID,
		Name:              meta.Name,
		ClusterName:       meta.ClusterName,
		CreationTimestamp: meta.CreationTimestamp,
		Labels:            meta.Labels,
		Annotations:       meta.Annotations,
	}, "Common fields should match")
}

func TestGetNamespaceCommonFields(t *testing.T) {
	meta := metav1.ObjectMeta{}
	assert.Equal(t, GetCommonNamespacedFields(meta), CommonNamespacedFields{
		UID:               meta.UID,
		Name:              meta.Name,
		Namespace:         meta.Namespace,
		ClusterName:       meta.ClusterName,
		CreationTimestamp: meta.CreationTimestamp,
		Labels:            meta.Labels,
		Annotations:       meta.Annotations,
	}, "Common namespace fields should match")
}
