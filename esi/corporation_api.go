/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.2.dev3
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"encoding/json"
	"fmt"
)

type CorporationApiService service

/**
 * Get corporation information
 * Public information about a corporation  ---  Alternate route: &#x60;/v2/corporations/{corporation_id}/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId An Eve corporation ID
 * @param datasource(string) The server name you would like data from
 * @return *GetCorporationsCorporationIdOk
 */
func (a CorporationApiService) GetCorporationsCorporationId(corporationId int32, datasource interface{}) (*GetCorporationsCorporationIdOk, time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/corporations/{corporation_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetCorporationsCorporationIdOk)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return successPayload, time.Now(), err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, time.Now(), err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return successPayload, expires, err
}

/**
 * Get alliance history
 * Get a list of all the alliances a corporation has been a member of  ---  Alternate route: &#x60;/v1/corporations/{corporation_id}/alliancehistory/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/alliancehistory/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/alliancehistory/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId An EVE corporation ID
 * @param datasource(string) The server name you would like data from
 * @return []GetCorporationsCorporationIdAlliancehistory200Ok
 */
func (a CorporationApiService) GetCorporationsCorporationIdAlliancehistory(corporationId int32, datasource interface{}) ([]GetCorporationsCorporationIdAlliancehistory200Ok, time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/corporations/{corporation_id}/alliancehistory/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]GetCorporationsCorporationIdAlliancehistory200Ok)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return *successPayload, time.Now(), err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return *successPayload, time.Now(), err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return *successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return *successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return *successPayload, expires, err
}

/**
 * Get corporation icon
 * Get the icon urls for a corporation  ---  Alternate route: &#x60;/v1/corporations/{corporation_id}/icons/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/icons/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/icons/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId An EVE corporation ID
 * @param datasource(string) The server name you would like data from
 * @return *GetCorporationsCorporationIdIconsOk
 */
func (a CorporationApiService) GetCorporationsCorporationIdIcons(corporationId int32, datasource interface{}) (*GetCorporationsCorporationIdIconsOk, time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/corporations/{corporation_id}/icons/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(GetCorporationsCorporationIdIconsOk)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return successPayload, time.Now(), err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, time.Now(), err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return successPayload, expires, err
}

/**
 * Get corporation members
 * Read the current list of members if the calling character is a member.  ---  Alternate route: &#x60;/v2/corporations/{corporation_id}/members/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/members/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/members/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId A corporation ID
 * @param datasource(string) The server name you would like data from
 * @return []GetCorporationsCorporationIdMembers200Ok
 */
func (a CorporationApiService) GetCorporationsCorporationIdMembers(ts TokenSource, corporationId int32, datasource interface{}) ([]GetCorporationsCorporationIdMembers200Ok, time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/corporations/{corporation_id}/members/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]GetCorporationsCorporationIdMembers200Ok)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return *successPayload, time.Now(), err
	}

	if ts != nil {
		if t, err := ts.Token(); err != nil {
			return *successPayload, time.Now(), err
		} else if t != nil {
			t.SetAuthHeader(r)
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return *successPayload, time.Now(), err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return *successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return *successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return *successPayload, expires, err
}

/**
 * Get corporation member roles
 * Return the roles of all members if the character has the personnel manager role or any grantable role.  ---  Alternate route: &#x60;/v1/corporations/{corporation_id}/roles/&#x60;  Alternate route: &#x60;/legacy/corporations/{corporation_id}/roles/&#x60;  Alternate route: &#x60;/dev/corporations/{corporation_id}/roles/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationId A corporation ID
 * @param datasource(string) The server name you would like data from
 * @return []GetCorporationsCorporationIdRoles200Ok
 */
func (a CorporationApiService) GetCorporationsCorporationIdRoles(ts TokenSource, corporationId int32, datasource interface{}) ([]GetCorporationsCorporationIdRoles200Ok, time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/corporations/{corporation_id}/roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"corporation_id"+"}", fmt.Sprintf("%v", corporationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]GetCorporationsCorporationIdRoles200Ok)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return *successPayload, time.Now(), err
	}

	if ts != nil {
		if t, err := ts.Token(); err != nil {
			return *successPayload, time.Now(), err
		} else if t != nil {
			t.SetAuthHeader(r)
		}
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return *successPayload, time.Now(), err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return *successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return *successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return *successPayload, expires, err
}

/**
 * Get corporation names
 * Resolve a set of corporation IDs to corporation names  ---  Alternate route: &#x60;/v1/corporations/names/&#x60;  Alternate route: &#x60;/legacy/corporations/names/&#x60;  Alternate route: &#x60;/dev/corporations/names/&#x60;   ---  This route is cached for up to 3600 seconds
 *
 * @param corporationIds A comma separated list of corporation IDs
 * @param datasource(string) The server name you would like data from
 * @return []GetCorporationsNames200Ok
 */
func (a CorporationApiService) GetCorporationsNames(corporationIds []int64, datasource interface{}) ([]GetCorporationsNames200Ok, time.Time, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := "https://esi.tech.ccp.is/latest/corporations/names/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := a.client.typeCheckParameter(datasource, "string", "datasource"); err != nil {
		return nil, time.Now(), err
	}
	localVarQueryParams.Add("corporation_ids", a.client.parameterToString(corporationIds, "csv"))
	if datasource != nil {
		localVarQueryParams.Add("datasource", a.client.parameterToString(datasource, ""))
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.client.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new([]GetCorporationsNames200Ok)

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, "application/json")
	if err != nil {
		return *successPayload, time.Now(), err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return *successPayload, time.Now(), err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return *successPayload, time.Now(), errors.New(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return *successPayload, time.Now(), err
	}

	expires := cacheExpires(localVarHttpResponse)
	return *successPayload, expires, err
}
