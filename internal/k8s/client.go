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

// Init creates in-cluster kubernetes configuration and a client set using the configuration.
// This returns error if KUBERNETES_SERVICE_HOST or KUBERNETES_SERVICE_PORT environment variables are not set.
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

// GetClient returns kubernetes interface that can be used to communicate with API server.
// Init function should be called before this function is called
func GetClient() kubernetes.Interface {
	return clientset
}

// SetClient is helper function to override the kubernetes interface with fake one for testing.
func SetClient(c kubernetes.Interface) {
	clientset = c
}
