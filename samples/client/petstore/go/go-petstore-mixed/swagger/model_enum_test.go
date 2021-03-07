/*
 * Swagger Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EnumTest struct {
	EnumString string `json:"enum_string,omitempty"`
	EnumInteger int32 `json:"enum_integer,omitempty"`
	EnumNumber float64 `json:"enum_number,omitempty"`
	OuterEnum *OuterEnum `json:"outerEnum,omitempty"`
}