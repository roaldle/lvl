package utils

import (
	"bitbucket.org/level27/lvl/types"
)

// Gets an app component from the API
/*
func (c *Client) Component(method string, category string, id interface{}, data interface{}) types.Component {
	var app types.Component

	switch method {
	case "GET":
		endpoint := fmt.Sprintf("appcomponents/%s/%s", category, id)
		c.invokeAPI("GET", endpoint, nil, &app)
	case "CREATE":
		endpoint := fmt.Sprintf("appcomponents/%s", category)
		c.invokeAPI("POST", endpoint, data, &app)
	case "UPDATE":
		endpoint := fmt.Sprintf("appcomponents/%s/%s", category, id)
		c.invokeAPI("PUT", endpoint, data, &app)
	case "DELETE":
		endpoint := fmt.Sprintf("appcomponents/%s/%s", category, id)
		c.invokeAPI("DELETE", endpoint, nil, nil)
	}

	return app
}
*/

func (c *Client) Components(filter string, number string, category string, cType string) types.Components {
	var components types.Components

	endpoint := "appcomponents/" + category + "?limit=" + number + "&filter=" + filter + "&type=" + cType
	err := c.invokeAPI("GET", endpoint, nil, &components)
	AssertApiError(err)

	return components
}