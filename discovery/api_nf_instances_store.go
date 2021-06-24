// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery  Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package discovery

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/http_wrapper"
	"github.com/free5gc/nrf/logger"
	"github.com/free5gc/nrf/producer"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

// SearchNFInstances - Search a collection of NF Instances
func HTTPSearchNFInstances(c *gin.Context) {
	// var searchNFInstance context.SearchNFInstances
	// c.BindQuery(&searchNFInstance)
	// logger.DiscoveryLog.Infoln("searchNFInstance: ", searchNFInstance)
	// logger.DiscoveryLog.Infoln("targetNFType: ", searchNFInstance.TargetNFType)
	// logger.DiscoveryLog.Infoln("requesterNFType: ", searchNFInstance.RequesterNFType)
	// logger.DiscoveryLog.Infoln("ChfSupportedPlmn: ", searchNFInstance.ChfSupportedPlmn)

	req := http_wrapper.NewRequest(c.Request, nil)
	req.Query = c.Request.URL.Query()
	httpResponse := producer.HandleNFDiscoveryRequest(req)

	responseBody, err := openapi.Serialize(httpResponse.Body, "application/json")
	if err != nil {
		logger.DiscoveryLog.Warnln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(httpResponse.Status, "application/json", responseBody)
	}
}
