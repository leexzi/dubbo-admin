/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package crdclient

import (
	api "github.com/apache/dubbo-admin/api/resource/v1alpha1"
	"github.com/apache/dubbo-admin/pkg/core/logger"
	"github.com/apache/dubbo-admin/pkg/core/model"
	"github.com/apache/dubbo-admin/pkg/core/schema/collection"
	gvks "github.com/apache/dubbo-admin/pkg/core/schema/gvk"
	"github.com/apache/dubbo-admin/pkg/dds/storage"
	"k8s.io/utils/strings/slices"
	"sync"
	"sync/atomic"
)

type PushContext struct {
	rootNamespace string
	mutex         *sync.Mutex
	revision      map[string]int64
	storage       *storage.Storage
	cache         ConfigStoreCache
}

type Handler interface {
	NotifyWithIndex(schema collection.Schema) error
}

func NewHandler(storage *storage.Storage, rootNamespace string) *PushContext {
	return &PushContext{
		mutex:         &sync.Mutex{},
		revision:      map[string]int64{},
		storage:       storage,
		rootNamespace: rootNamespace,
	}
}

func (p *PushContext) NotifyWithIndex(schema collection.Schema) error {
	gvk := schema.Resource().GroupVersionKind()
	configs, err := p.cache.List(gvk, NamespaceAll)
	data := make([]model.Config, len(configs))
	if err != nil {
		logger.Sugar().Error("fail to get the cache from client-go Index")
		return err
	}
	// ignore the "To"
	if gvk.String() == gvks.Authorization {
		// WARNING: the client-go cache is read-only, if we must change the resource, we need to deep copy first
		for _, config := range configs {
			deepCopy := config.DeepCopy()
			policy := deepCopy.Spec.(*api.AuthorizationPolicy)
			if p.rootNamespace != config.Namespace {
				if policy.GetRules() == nil {
					policy.Rules = append(policy.Rules, &api.AuthorizationPolicyRule{
						To: &api.AuthorizationPolicyTarget{
							Namespaces: []string{config.Namespace},
						},
					})
				} else {
					for _, rule := range policy.Rules {
						if rule.To != nil {
							rule.To = &api.AuthorizationPolicyTarget{}
						}
						if !slices.Contains(rule.To.Namespaces, config.Namespace) {
							rule.To.Namespaces = append(rule.To.Namespaces, config.Namespace)
						}
					}
				}
			}
			data = append(data, deepCopy)
		}
	} else if gvk.String() == gvks.Authentication {
		for _, config := range configs {
			deepCopy := config.DeepCopy()
			policy := deepCopy.Spec.(*api.AuthenticationPolicy)
			if p.rootNamespace != config.Namespace {
				if policy.GetSelector() == nil {
					policy.Selector = append(policy.Selector, &api.AuthenticationPolicySelector{
						Namespaces: []string{config.Namespace},
					})
				}
			} else {
				for _, selector := range policy.Selector {
					if !slices.Contains(selector.Namespaces, config.Namespace) {
						selector.Namespaces = append(selector.Namespaces, config.Namespace)
					}
				}
			}
			data = append(data, deepCopy)
		}
	} else {
		data = configs
	}
	rev := p.revision[gvk.String()]
	atomic.AddInt64(&rev, 1)
	p.revision[gvk.String()] = rev

	origin := &storage.Origin{
		Gvk:      gvk.String(),
		Revision: p.revision[gvk.String()],
		Data:     data,
	}
	p.storage.LatestRules[gvk.String()] = origin

	p.storage.Mutex.RLock()
	defer p.storage.Mutex.RUnlock()
	for _, c := range p.storage.Connection {
		c.RawRuleQueue.Add(origin)
	}
	return nil
}
