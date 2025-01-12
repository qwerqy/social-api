package main

import (
	"net/http"

	"github.com/qwerqy/social-api-go/internal/store"
)

// GetFeed godoc
//	@Summary		Gets user feed
//	@Description	Gets user feed by ID
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			since	query		string	false	"Since"
//	@Param			until	query		string	false	"Until"
//	@Param			limit	query		int		false	"Limit"
//	@Param			offset	query		int		false	"Offset"
//	@Param			sort	query		string	false	"Sort"
//	@Param			tags	query		string	false	"Tags"
//	@Param			search	query		string	false	"Search"
//	@Success		200		{object}	[]store.PostWithMetadata
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {

	fq := store.PaginatedFeedQuery{
		Limit: 20,
		Offset: 0, 
		Sort: "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestError(w,r,err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestError(w,r,err)
		return
	}

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, int64(97), fq)

	if err != nil {
		app.internalServerError(w,r,err)
		return
	}

	if err := app.jsonResponse(w,http.StatusOK,feed); err != nil {
		app.internalServerError(w,r,err)
		return
	}
}