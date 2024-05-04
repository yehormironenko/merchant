package reseller

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/yehormironenko/reseller/pkg/api"
	"github.com/yehormironenko/reseller/pkg/model"

	"merchant/internal/controllers/requests"
)

type SearchBookService struct {
	resellerApiClient api.ResellerApiClient
	logger            *zerolog.Logger
}

func NewSearchBookService(resellerApiClient api.ResellerApiClient, logger *zerolog.Logger) *SearchBookService {
	return &SearchBookService{
		resellerApiClient: resellerApiClient,
		logger:            logger,
	}
}

func (sbs SearchBookService) SearchBook(ctx context.Context, req requests.SearchBook) (model.Books, error) {
	sbs.logger.Info().Msg("service:SearchBookService")

	resp, _, err := sbs.resellerApiClient.GetBookByParams(ctx, req.Title, req.Author, req.Genre)
	if err != nil {
		return model.Books{}, err
	}

	return resp, nil
}
