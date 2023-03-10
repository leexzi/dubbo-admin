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

package constant

const (
	DubboPropertyKey         = "dubbo.properties"
	RegistryAddressKey       = "dubbo.registry.address"
	MetadataReportAddressKey = "dubbo.metadata-report.address"
)

const (
	AnyValue              = "*"
	InterfaceKey          = "interface"
	GroupKey              = "group"
	VersionKey            = "version"
	ClassifierKey         = "classifier"
	CategoryKey           = "category"
	ProvidersCategory     = "providers"
	ConsumersCategory     = "consumers"
	RoutersCategory       = "routers"
	ConfiguratorsCategory = "configurators"
	EnabledKey            = "enabled"
	CheckKey              = "check"
	AdminProtocol         = "admin"
	Side                  = "side"
	ConsumerSide          = "consumer"
	ProviderSide          = "provider"
	ConsumerProtocol      = "consumer"
	EmptyProtocol         = "empty"
	DefaultGroup          = "dubbo"
	ApplicationKey        = "application"
	DynamicKey            = "dynamic"
	SerializationKey      = "serialization"
	TimeoutKey            = "timeout"
	DefaultTimeout        = 1000
	WeightKey             = "weight"
	DefaultWeight         = 100
	OwnerKey              = "owner"
	Service               = "service"
	Colon                 = ":"
	InterrogationPoint    = "?"
	IP                    = "ip"
	PlusSigns             = "+"
	PunctuationPoint      = "."
)

const (
	MetricsQps                        = "" // QPS
	MetricsHttpRequestTotalCount      = "" // Total number of http requests
	MetricsHttpRequestSuccessCount    = "" // Total number of http successful requests
	MetricsHttpRequestOutOfTimeCount  = "" // Total number of http out of time requests
	MetricsHttpRequestAddressNotFount = "" // Total number of HTTP requests where the address cannot be found
	MetricsHttpRequestOtherException  = "" // Total number of other errors for http requests

	ApiPrefix = "/api/v1"
	EpQuery   = ApiPrefix + "/query?query="
)
