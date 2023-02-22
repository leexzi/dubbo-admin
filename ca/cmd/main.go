// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/tls"
	"github.com/apache/dubbo-admin/ca/pkg/cert"
	"github.com/apache/dubbo-admin/ca/pkg/config"
	"github.com/apache/dubbo-admin/ca/pkg/security"
	"os"
)

// TODO read namespace from env
const namespace = "dubbo-system"

func main() {
	// TODO read options from env
	options := &config.Options{
		EnableKubernetes: true,
		Namespace:        namespace,
		PlainServerPort:  30060,
		SecureServerPort: 30062,
		DebugPort:        30070,
		CaValidity:       30 * 24 * 60 * 60 * 1000, // 30 day
		CertValidity:     1 * 60 * 60 * 1000,       // 1 hour
	}

	s := &security.Server{
		Options: options,
		CertStorage: &cert.Storage{
			AuthorityCert: &cert.Cert{},
			ServerCerts:   map[string]*tls.Certificate{},
		},
	}

	s.Init()
	s.Start()

	c := make(chan os.Signal)
	<-c
}