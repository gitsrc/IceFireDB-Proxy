/*
 *
 *  * Licensed to the Apache Software Foundation (ASF) under one or more
 *  * contributor license agreements.  See the NOTICE file distributed with
 *  * this work for additional information regarding copyright ownership.
 *  * The ASF licenses this file to You under the Apache License, Version 2.0
 *  * (the "License"); you may not use this file except in compliance with
 *  * the License.  You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package proxy

import (
	"net"

	"github.com/sirupsen/logrus"

	"github.com/IceFireDB/IceFireDB-Proxy/pkg/config"
	"github.com/IceFireDB/IceFireDB-Proxy/utils"
	"github.com/IceFireDB/components-go/bareneter"
)

func (p *Proxy) accept(conn bareneter.Conn) bool {
	if config.Get().IPWhiteList.Enable {
		host, _, _ := net.SplitHostPort(conn.RemoteAddr())
		if !utils.InArray(host, config.Get().IPWhiteList.List) {
			return false
		}
	}

	logrus.Infof("client connect %s", conn.RemoteAddr())

	return true
}

func (p *Proxy) closed(conn bareneter.Conn, err error) {
	logrus.Infof("client closed %s", conn.RemoteAddr())
	return
}