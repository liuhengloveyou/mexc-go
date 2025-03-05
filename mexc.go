package mexc

import (
	"context"

	mexchttp "github.com/liuhengloveyou/mexc-go/http"
	mexchttpmarket "github.com/liuhengloveyou/mexc-go/http/market"
)

type Rest struct {
	MarketService *mexchttpmarket.Service
}

func NewRest(ctx context.Context, mexcHttp *mexchttp.Client) (*Rest, error) {
	marketService, err := mexchttpmarket.New(ctx, mexcHttp)
	if err != nil {
		return nil, err
	}

	return &Rest{
		MarketService: marketService,
	}, nil
}
