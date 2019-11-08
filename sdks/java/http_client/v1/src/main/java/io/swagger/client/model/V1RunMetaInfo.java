// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon sdk
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.14.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * V1RunMetaInfo
 */

public class V1RunMetaInfo {
  @SerializedName("service")
  private Boolean service = null;

  @SerializedName("workflow_strategy")
  private String workflowStrategy = null;

  @SerializedName("workflow_concurrency")
  private Integer workflowConcurrency = null;

  public V1RunMetaInfo service(Boolean service) {
    this.service = service;
    return this;
  }

   /**
   * Get service
   * @return service
  **/
  @ApiModelProperty(value = "")
  public Boolean isService() {
    return service;
  }

  public void setService(Boolean service) {
    this.service = service;
  }

  public V1RunMetaInfo workflowStrategy(String workflowStrategy) {
    this.workflowStrategy = workflowStrategy;
    return this;
  }

   /**
   * Get workflowStrategy
   * @return workflowStrategy
  **/
  @ApiModelProperty(value = "")
  public String getWorkflowStrategy() {
    return workflowStrategy;
  }

  public void setWorkflowStrategy(String workflowStrategy) {
    this.workflowStrategy = workflowStrategy;
  }

  public V1RunMetaInfo workflowConcurrency(Integer workflowConcurrency) {
    this.workflowConcurrency = workflowConcurrency;
    return this;
  }

   /**
   * Get workflowConcurrency
   * @return workflowConcurrency
  **/
  @ApiModelProperty(value = "")
  public Integer getWorkflowConcurrency() {
    return workflowConcurrency;
  }

  public void setWorkflowConcurrency(Integer workflowConcurrency) {
    this.workflowConcurrency = workflowConcurrency;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1RunMetaInfo v1RunMetaInfo = (V1RunMetaInfo) o;
    return Objects.equals(this.service, v1RunMetaInfo.service) &&
        Objects.equals(this.workflowStrategy, v1RunMetaInfo.workflowStrategy) &&
        Objects.equals(this.workflowConcurrency, v1RunMetaInfo.workflowConcurrency);
  }

  @Override
  public int hashCode() {
    return Objects.hash(service, workflowStrategy, workflowConcurrency);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1RunMetaInfo {\n");
    
    sb.append("    service: ").append(toIndentedString(service)).append("\n");
    sb.append("    workflowStrategy: ").append(toIndentedString(workflowStrategy)).append("\n");
    sb.append("    workflowConcurrency: ").append(toIndentedString(workflowConcurrency)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
