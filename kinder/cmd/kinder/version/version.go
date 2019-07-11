/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"k8s.io/kubeadm/kinder/pkg/constants"
	kindversion "sigs.k8s.io/kind/cmd/kind/version"
)

// NewCommand returns a new cobra.Command for version
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "version",
		Short: "prints the kind CLI version",
		Long:  "prints the kind CLI version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("kinder version: %s\nkind   version: %s\n", constants.KinderVersion, kindversion.Version)
			return nil
		},
	}
	return cmd
}
