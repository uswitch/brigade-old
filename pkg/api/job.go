package api

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/uswitch/brigade/pkg/storage"
)

// Job represents the job api handlers.
type Job struct {
	store storage.Store
}

// Get creates a new gin handler for the GET /job/:id endpoint
func (api Job) Get(c *gin.Context) {
	id := c.Params.ByName("id")
	job, err := api.store.GetJob(id)
	if err != nil {
		c.JSON(http.StatusNotFound, struct{}{})
		return
	}
	c.JSON(http.StatusOK, job)
}

// Logs creates a new gin handler for the GET /job/:id/logs endpoint
func (api Job) Logs(c *gin.Context) {
	id := c.Params.ByName("id")
	container := c.Query("container")
	job, err := api.store.GetJob(id)
	if err != nil {
		c.JSON(http.StatusNotFound, struct{}{})
		return
	}
	if c.Query("stream") == "true" {
		logReader, err := api.store.GetJobLogStream(job, container)
		if err != nil {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		defer logReader.Close()
		io.Copy(c.Writer, logReader)
	} else {
		logs, err := api.store.GetJobLog(job, container)
		if err != nil {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		if len(logs) == 0 {
			c.JSON(http.StatusNoContent, nil)
		}
		c.JSON(http.StatusOK, logs)
	}
}
