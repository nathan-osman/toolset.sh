package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nathan-osman/toolset.sh/manager"
)

func convertContextToInput(c *gin.Context, t manager.Tool) *manager.Input {
	params := map[string]string{}
	for _, p := range t.Meta().Params {
		v := c.Query(p.Name)
		if v != "" {
			params[p.Name] = v
		}

		// TODO: verify against Options
	}
	return &manager.Input{
		Params: params,
	}
}
