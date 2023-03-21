// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import "github.com/apache/dubbo-admin/pkg/admin/constant"

func ServiceName2Map(serviceName string) map[string]string {
	group := GetGroup(serviceName)
	version := GetVersion(serviceName)
	interfaze := GetInterface(serviceName)

	var ret map[string]string
	if !IsEmpty(group) {
		ret[constant.InterfaceKey] = interfaze
	}
	if !IsEmpty(version) {
		ret[constant.VersionKey] = version
	}
	if !IsEmpty(group) {
		ret[constant.GroupKey] = group
	}

	return ret
}

//func getIdFromDTO(baseDTO BaseDTO) {
//
//}
