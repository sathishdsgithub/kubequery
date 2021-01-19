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
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var clientset kubernetes.Interface

// Init TODO
func Init() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}

// GetClient TODO
func GetClient() kubernetes.Interface {
	return clientset
}

// SetClient TODO
func SetClient(c kubernetes.Interface) {
	clientset = c
}
