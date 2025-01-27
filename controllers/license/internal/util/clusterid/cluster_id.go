// Copyright © 2023 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clusterid

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetClusterID(ctx context.Context, c client.Client) (string, error) {
	ns := &corev1.Namespace{}
	err := c.Get(ctx, client.ObjectKey{Name: "kube-system"}, ns)
	if err != nil {
		return "", err
	}
	res := string(ns.UID)
	if res == "" || len(res) < 8 {
		return "", fmt.Errorf("failed to get cluster id")
	}
	return res[0:8], nil
}
