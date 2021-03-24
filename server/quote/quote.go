// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quote

import (
	"context"

	"bitbucket.com/iamjollof/server/models"
	"bitbucket.com/iamjollof/server/plugins"
	"bitbucket.com/iamjollof/server/protos/quote"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
)

//Server - Struct holding Quotes
type Server struct {
	quote.QuoteServer
}

//GetAllQuotes - Lists all the quotes
//Todo: Handle errors
//sudo service postgresql start
//export PATH="$PATH:$(go env GOPATH)/bin"
func (*Server) GetAllQuotes(ctx context.Context, r *quote.GetAllQuotesRequest) (*quote.GetAllQuotesResponse, error) {
	quotes, count, err := models.GetAllQuotes(r.GetPageNumber(), r.GetResultPerPage())
	if err != nil {
		plugins.LogError("API Service", "Error at: GetAllQuotes", err)
		return &quote.GetAllQuotesResponse{}, err
	}
	var arrayOfQuotes []*quote.GetQuoteResponse
	copier.Copy(&arrayOfQuotes, quotes)
	for i := 0; i < len(arrayOfQuotes); i++ {
		createdAt, err := ptypes.TimestampProto(quotes[i].CreatedAt)
		if err != nil {
			plugins.LogError("API Service", "Error at: GetAllQuotes - time.Time to Proto TimeStamp conversion err", err)
			return &quote.GetAllQuotesResponse{}, err
		}
		arrayOfQuotes[i].CreatedAt = createdAt
	}
	return &quote.GetAllQuotesResponse{
		PageNumber:    r.GetPageNumber(),
		ResultPerPage: r.GetResultPerPage(),
		TotalCount:    count,
		Quotes:        arrayOfQuotes,
	}, nil
}

//CreateQuote - Create new Quote
func (*Server) CreateQuote(ctx context.Context, r *quote.CreateQuoteRequest) (*quote.CreateQuoteResponse, error) {
	var newQuote models.Quote
	copier.Copy(&newQuote, r)
	createdQuote, err := models.CreateQuote(&newQuote)
	if err != nil {
		plugins.LogError("API Service", "Error at: CreateQuote", err)
		return &quote.CreateQuoteResponse{}, err
	}
	var test quote.CreateQuoteResponse
	copier.Copy(&test, createdQuote)
	return &test, nil
}

//GetQuote - Get Quote by ID
func (*Server) GetQuote(ctx context.Context, r *quote.GetQuoteRequest) (*quote.GetQuoteResponse, error) {
	singleQuote, err := models.GetQuote(r.ID)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetQuote", err)
		return &quote.GetQuoteResponse{}, err
	}
	var test quote.GetQuoteResponse
	copier.Copy(&test, singleQuote)
	createdAt, err := ptypes.TimestampProto(singleQuote.CreatedAt)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetQuote - time.Time to Proto TimeStamp conversion err", err)
		return &quote.GetQuoteResponse{}, err
	}
	test.CreatedAt = createdAt
	return &test, nil
}

//UpdateQuote - Update Quote by ID
func (*Server) UpdateQuote(ctx context.Context, r *quote.UpdateQuoteRequest) (*quote.UpdateQuoteResponse, error) {
	var newQuote models.Quote
	copier.Copy(&newQuote, r)
	if err := models.UpdateQuote(newQuote, r.ID); err != nil {
		plugins.LogError("API Service", "Error at: UpdateQuote", err)
		return &quote.UpdateQuoteResponse{}, err
	}
	return &quote.UpdateQuoteResponse{
		Message: "Quote Updated successfully",
	}, nil
}

//DeleteQuote - Delete Quote by ID
func (*Server) DeleteQuote(ctx context.Context, r *quote.DeleteQuoteRequest) (*quote.DeleteQuoteResponse, error) {
	if err := models.DeleteQuote(r.ID); err != nil {
		plugins.LogError("API Service", "Error at: DeleteQuote", err)
		return &quote.DeleteQuoteResponse{}, err
	}
	return &quote.DeleteQuoteResponse{}, nil
}
