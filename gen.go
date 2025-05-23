// Copyright (c) 2020-2025 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package suites contains go:generate commands.
package suites

//go:generate bash -c "rm -rf ./suites"
//go:generate gotestmd ../deployments-k8s/examples ./suites github.com/bszirtes/integration-tests-1/extensions/base
//go:generate goimports -w -local github.com/networkservicemesh -d "./suites"
//go:generate goimports -w -local github.com/bszirtes/integration-tests-1 -d "./suites"
