package service

import (
	"context"

	"github.com/yehormironenko/reseller/pkg/model"

	"merchant/internal/controllers/requests"
)

type ResellerSearchService interface {
	SearchBook(ctx context.Context, req requests.SearchBook) (model.Books, error)
}
