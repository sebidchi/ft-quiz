package application

import (
	"context"

	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
)

type GetUserResultsQuery struct {
	userId string
}

func NewGetUserResultsQuery(userId string) *GetUserResultsQuery {
	return &GetUserResultsQuery{userId: userId}
}

func (g GetUserResultsQuery) Id() string {
	return "GetUserResultsQuery"
}

type GetUserResultsQueryHandler struct {
	userResultsFetcher *domain.UserResultsFetcher
}

func NewGetUserResultsQueryHandler(userResultsFetcher *domain.UserResultsFetcher) *GetUserResultsQueryHandler {
	return &GetUserResultsQueryHandler{userResultsFetcher: userResultsFetcher}
}

func (g GetUserResultsQueryHandler) Handle(ctx context.Context, query bus.Query) (interface{}, error) {
	userRQuery, ok := query.(*GetUserResultsQuery)
	if !ok {
		return nil, bus.NewInvalidQuery(query)
	}

	res, err := g.userResultsFetcher.Fetch(ctx, userRQuery.userId)

	if err != nil {
		return nil, err
	}

	return NewUserResultsResponse(res), nil
}
