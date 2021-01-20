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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/fake"
)

func TestInitClientset(t *testing.T) {
	err := initClientset()
	assert.Error(t, err, "Init should fail")

	os.Setenv("KUBERNETES_SERVICE_HOST", "localhost")
	os.Setenv("KUBERNETES_SERVICE_PORT", "65000")
	err = initClientset()
	assert.Nil(t, err, "Init should succeed")
}

func TestGetClient(t *testing.T) {
	os.Setenv("KUBERNETES_SERVICE_HOST", "localhost")
	os.Setenv("KUBERNETES_SERVICE_PORT", "65000")
	err := initClientset()
	assert.Nil(t, err, "Init should succeed")

	clientset := GetClient()
	assert.NotNil(t, clientset, "Clientset should be valid")
}

func TestGetClusterUID(t *testing.T) {
	uid := types.UID("")
	SetClient(fake.NewSimpleClientset(), uid)
	assert.Equal(t, uid, GetClusterUID())
}
