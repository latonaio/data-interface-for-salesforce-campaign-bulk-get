package resources

import (
	"errors"
	"fmt"
)

const suffixRelatedList = "RelatedList"

// Campaign struct
type Campaign struct {
	method   string
	metadata map[string]interface{}
}

func (c *Campaign) objectName() string {
	const obName = "Campaign"
	return obName
}

// NewCampaign writes that new Campaign instance
func NewCampaign(metadata map[string]interface{}) (*Campaign, error) {
	rawMethod, ok := metadata["method"]
	if !ok {
		return nil, errors.New("missing requested parameters: method")
	}
	method, ok := rawMethod.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	return &Campaign{
		method:   method,
		metadata: metadata,
	}, nil
}

// BuildMetadata mold campaign get metadata
func (c *Campaign) BuildMetadata() (map[string]interface{}, error) {
	switch c.method {
	case "get":
		return c.buildMetadata("campaign_bulk_get", c.method, c.objectName()+suffixRelatedList, "", nil, ""), nil
	}
	return nil, fmt.Errorf("invalid method: %s", c.method)
}

func (c *Campaign) buildMetadata(connectionKey, method, object, pathParam string, queryParams map[string]string, body string) map[string]interface{} {
	metadata := map[string]interface{}{
		"method":         method,
		"object":         object,
		"connection_key": connectionKey,
	}
	if len(pathParam) > 0 {
		metadata["path_param"] = pathParam
	}
	if queryParams != nil {
		metadata["query_params"] = queryParams
	}
	if body != "" {
		metadata["body"] = body
	}
	return metadata
}
