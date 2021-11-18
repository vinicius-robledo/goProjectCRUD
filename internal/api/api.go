package api

import (
	"fmt"
)

func main() {
	fmt.Println("Hi there")
}
/*import (

	"github.com/gin-gonic/gin"
	"net/http"
)*/


// PingHandler returns a successful pong answer to all HTTP requests.
//func (h HealthChecker) PingHandler(c *gin.Context) {
//	if txn := nrgin.Transaction(c); txn != nil {
//		txn.Ignore()
//	}
//
//	c.String(http.StatusOK, "pong")
//}