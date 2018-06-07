/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package server

import (
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/skydive-project/skydive/api/types"
)

// AlertResourceHandler aims to creates and manage a new Alert.
type AlertResourceHandler struct {
	ResourceHandler
}

// AlertAPIHandler aims to exposes the Alert API.
type AlertAPIHandler struct {
	BasicAPIHandler
}

// New creates a new alert
func (a *AlertResourceHandler) New() types.Resource {
	id, _ := uuid.NewV4()

	return &types.Alert{
		UUID:       id.String(),
		CreateTime: time.Now().UTC(),
	}
}

// Name returns resource name "alert"
func (a *AlertResourceHandler) Name() string {
	return "alert"
}

// RegisterAlertAPI registers an Alert's API to a designated API Server
func RegisterAlertAPI(apiServer *Server) (*AlertAPIHandler, error) {
	alertAPIHandler := &AlertAPIHandler{
		BasicAPIHandler: BasicAPIHandler{
			ResourceHandler: &AlertResourceHandler{},
			EtcdKeyAPI:      apiServer.EtcdKeyAPI,
		},
	}
	if err := apiServer.RegisterAPIHandler(alertAPIHandler); err != nil {
		return nil, err
	}
	return alertAPIHandler, nil
}
