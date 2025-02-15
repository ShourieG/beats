// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package processors

import (
	"io/fs"

	"github.com/elastic/elastic-agent-system-metrics/metric/system/cgroup"
	"github.com/elastic/elastic-agent-system-metrics/metric/system/resolve"
)

// InitCgroupHandler is a type for creating stubs for the cgroup resolver. Used primarily for testing.
type InitCgroupHandler = func(rootfsMountpoint resolve.Resolver, ignoreRootCgroups bool) (CGReader, error)

// CGReader wraps the group Reader.ProcessCgroupPaths() call, this allows us to
// set different cgroups readers for testing.
type CGReader interface {
	ProcessCgroupPaths(pid int) (cgroup.PathList, error)
}

// NilCGReader does nothing
type NilCGReader struct {
}

// ProcessCgroupPaths returns a blank pathLists and fs.ErrNotExist
func (*NilCGReader) ProcessCgroupPaths(_ int) (cgroup.PathList, error) {
	return cgroup.PathList{}, fs.ErrNotExist
}
