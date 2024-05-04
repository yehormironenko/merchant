package reseller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"merchant/internal/controllers/requests"
	"merchant/internal/service"
)

func SearchBook(resellerService service.ResellerSearchService, logger *zerolog.Logger) gin.HandlerFunc {

	return func(context *gin.Context) {
		logger.Info().Msg("handlers:SearchBookExecutor")

		var searchBookRequest requests.SearchBook

		if err := context.ShouldBindJSON(&searchBookRequest); err != nil {
			logger.Error().AnErr("invalid json request", err)
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := resellerService.SearchBook(context, searchBookRequest)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			logger.Err(err).Msg("Bad request to register endpoint")
			return
		}

		logger.Info().Object("data", searchBookRequest).Msg("search book request with the following data")

		context.JSON(http.StatusOK, response)
	}
}
