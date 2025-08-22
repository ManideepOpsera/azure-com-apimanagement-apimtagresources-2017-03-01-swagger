package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// TagResourceCollection represents the TagResourceCollection schema from the OpenAPI specification
type TagResourceCollection struct {
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []TagResourceContract `json:"value,omitempty"` // Page values.
}

// TagResourceContract represents the TagResourceContract schema from the OpenAPI specification
type TagResourceContract struct {
	Id string `json:"id,omitempty"` // Resource ID.
	Name string `json:"name,omitempty"` // Resource name.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource.
}

// TagResourceContractProperties represents the TagResourceContractProperties schema from the OpenAPI specification
type TagResourceContractProperties struct {
	Operation OperationEntityContract `json:"operation,omitempty"` // Operation Entity Contract Properties.
	Product interface{} `json:"product,omitempty"` // Product details.
	Tag interface{} `json:"tag,omitempty"` // Tag Contract details.
	Api interface{} `json:"api,omitempty"` // API details.
}

// OperationEntityContract represents the OperationEntityContract schema from the OpenAPI specification
type OperationEntityContract struct {
	Id string `json:"id,omitempty"` // Resource ID.
	Name string `json:"name,omitempty"` // Resource name.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource.
}

// OperationEntityContractProperties represents the OperationEntityContractProperties schema from the OpenAPI specification
type OperationEntityContractProperties struct {
	Displayname string `json:"displayName,omitempty"` // Operation name.
	Method string `json:"method,omitempty"` // A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
	Urltemplate string `json:"urlTemplate,omitempty"` // Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
	Apiname string `json:"apiName,omitempty"` // Api Name.
	Apirevision string `json:"apiRevision,omitempty"` // Api Revision.
	Apiversion string `json:"apiVersion,omitempty"` // Api Version.
	Description string `json:"description,omitempty"` // Operation Description.
}
